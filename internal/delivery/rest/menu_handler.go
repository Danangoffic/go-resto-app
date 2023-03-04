package rest

import (
	"fmt"
	"net/http"
	"resto-app/internal/tracing"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Menu List")
	defer span.End()

	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(ctx, menuType)
	if err != nil {
		fmt.Printf("error with %s\n", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    menuData,
		"status":  200,
		"message": "success",
	})
}

func (h *handler) GetMenuDetail(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Menu Detail")
	defer span.End()

	OrderCode := c.Param("OrderCode")
	menuData, err := h.restoUsecase.GetMenu(ctx, OrderCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    menuData,
		"status":  200,
		"message": "success",
	})
}
