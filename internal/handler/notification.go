package handler

import (
	"awesomeProject/enums"
	"awesomeProject/internal/services"
	"awesomeProject/internal/web"
)

type NotificationRequest struct {
	Type     enums.Type      `json:"type"`
	Coin     enums.Coin      `json:"coin"`
	Channels []enums.Channel `json:"channels"`
	Data     interface{}     `json:"data"`
}

type Notification interface {
	Create(request *web.Request) (web.JSONResponse, web.ErrorInterface)
	List(request *web.Request) (web.JSONResponse, web.ErrorInterface)
	Delete(request *web.Request) (web.JSONResponse, web.ErrorInterface)
}

type NotificationHandler struct {
	service services.Notification
}

func NewNotificationHandler() Notification {
	return &NotificationHandler{
		service: NewNotificationHandler(),
	}
}

func (NotificationHandler) Create(request *web.Request) (web.JSONResponse, web.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (NotificationHandler) List(request *web.Request) (web.JSONResponse, web.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}

func (NotificationHandler) Delete(request *web.Request) (web.JSONResponse, web.ErrorInterface) {
	//TODO implement me
	panic("implement me")
}
