package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resto-app/internal/model"
	"resto-app/internal/model/constant"

	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
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
	orderData, err := h.restoUsecase.Order(request)
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
	orderID := c.Param("orderID")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUsecase.GetOrderData(model.GetOrderDataRequest{
		OrderID: orderID,
		UserID:  userID,
	})
	if err != nil {
		fmt.Printf("Error %s\n", err.Error())

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
