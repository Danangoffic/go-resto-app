package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resto-app/internal/model"
	"resto-app/internal/model/constant"
	"resto-app/internal/tracing"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  200,
		})
	}

	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	fmt.Printf("request api : %v\n", request)
	orderData, err := h.restoUsecase.Order(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  500,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"status":  200,
		"data":    orderData,
	})
}

func (h *handler) GetOrderData(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderData(ctx, model.GetOrderDataRequest{
		OrderID: orderID,
		UserID:  userID,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][GetOrderData] unable to get order data\n")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  500,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"status":  200,
		"data":    orderData,
	})
}
