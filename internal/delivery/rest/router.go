package rest

import "github.com/labstack/echo/v4"

func LoadRouters(e *echo.Echo, handler *handler) {

	authMiddleware := GetAuthMiddleware(handler.restoUsecase)

	api := e.Group("/api")

	menuGroup := api.Group("/menu")
	menuGroup.GET("", handler.GetMenuList)
	menuGroup.GET("/:OrderCode", handler.GetMenuDetail)

	orderGroup := api.Group("/order", authMiddleware.CheckAuth)
	orderGroup.POST("", handler.Order)
	orderGroup.GET("/:orderID", handler.GetOrderData)

	userGroup := api.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}
