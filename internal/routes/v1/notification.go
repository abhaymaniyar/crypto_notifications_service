package v1

import (
	"awesomeProject/internal/handler"
	"awesomeProject/internal/handler/middleware"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router) {
	notificationHandler := handler.NewNotificationHandler()
	router.POST("/api/v1/notification", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, notificationHandler.Create))
	router.GET("/api/v1/notifications", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, notificationHandler.List))
	router.DELETE("/api/v1/notifications/:id", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, notificationHandler.Delete))
}
