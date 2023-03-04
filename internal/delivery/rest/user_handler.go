package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resto-app/internal/model"
	"resto-app/internal/tracing"

	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Register new User")
	defer span.End()

	var request model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		fmt.Printf("Got Error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  200,
		})
	}

	userData, err := h.restoUsecase.RegisterUser(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  200,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"status":  200,
		"data":    userData,
	})
}

func (h *handler) Login(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Register new User")
	defer span.End()

	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("Got Error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "failed",
		})
	}

	sessionData, err := h.restoUsecase.Login(ctx, request)
	if err != nil {
		fmt.Printf("Got Error %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  400,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"status":  200,
		"data":    sessionData,
	})
}
