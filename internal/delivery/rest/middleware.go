package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"resto-app/internal/model/constant"
	"resto-app/internal/tracing"
	"resto-app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func LoadMiddelwares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://restocilik.com"},
	}))
}

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{restoUsercase: restoUsecase}
}

type authMiddleware struct {
	restoUsercase resto.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "Check Authentication")
		defer span.End()
		fmt.Printf("request :\n%v", json.NewDecoder(c.Request().Body))
		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("[delivery][rest][middleware][CheckAuth] unable to get session data\n")
			return &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: map[string]interface{}{
					"status":  403,
					"message": err.Error(),
				},
				Internal: err,
			}
		}

		userID, err := am.restoUsercase.CheckSession(ctx, sessionData)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("[delivery][rest][middleware][CheckAuth] unable to get session data\n")
			return &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: map[string]interface{}{
					"status":  403,
					"message": err.Error(),
				},
				Internal: err,
			}
		}

		if userID == "" {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("[delivery][rest][middleware][CheckAuth] Invalid token\n")
			return &echo.HTTPError{
				Code: http.StatusUnauthorized,
				Message: map[string]interface{}{
					"status":  403,
					"message": "Invalid token",
				},
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}
