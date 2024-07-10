package v1

import (
	"awesomeProject/internal/handler"
	"awesomeProject/internal/handler/middleware"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes(router *httprouter.Router) {
	// create notif > notification_id
	// list notif > pagination_params
	// delete notif >
	handler.NotificationHandler
	router.POST("/api/v1/notification", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, handler.C))
	router.GET("/api/v1/notifications", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, nil))
	router.DELETE("/api/v1/notifications/:id", middleware.ServeV1Endpoint(middleware.EmptyMiddleware, nil))
}
