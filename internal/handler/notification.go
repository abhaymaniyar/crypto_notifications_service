package handler

import (
	"awesomeProject/internal/services"
	"awesomeProject/internal/web"
	"awesomeProject/repositories"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Notification interface {
	Create(request *web.Request) (*web.JSONResponse, web.ErrorInterface)
	List(request *web.Request) (*web.JSONResponse, web.ErrorInterface)
	Delete(request *web.Request) (*web.JSONResponse, web.ErrorInterface)
}

type NotificationHandler struct {
	service services.NotificationServiceI
}

func NewNotificationHandler() Notification {
	return &NotificationHandler{
		service: services.NewNotificationService(),
	}
}

func (h *NotificationHandler) Create(r *web.Request) (*web.JSONResponse, web.ErrorInterface) {
	req, err := validateCreateNotificationRequest(r.Request)
	if err != nil {
		return nil, errorResponse(err)
	}

	err = h.service.Create(req)
	if err != nil {
		return nil, errorResponse(err)
	}

	return (*web.JSONResponse)(nil), nil
}

func (h *NotificationHandler) List(r *web.Request) (*web.JSONResponse, web.ErrorInterface) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		return nil, web.ErrBadRequest("Invalid page number")
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		return nil, web.ErrBadRequest("Invalid per_page value")
	}

	offset := (page - 1) * perPage

	notifications, err := h.service.List(repositories.GetNotificationsOpts{
		Limit:  perPage,
		Offset: offset,
	})

	if err != nil {
		return nil, errorResponse(err)
	}

	res := map[string]interface{}{
		"notifications": notifications,
		"page":          page,
		"per_page":      perPage,
	}

	return (*web.JSONResponse)(&res), nil
}

func (h *NotificationHandler) Delete(r *web.Request) (*web.JSONResponse, web.ErrorInterface) {
	cId := r.GetPathParam("id")
	id, err := uuid.Parse(cId)
	if err != nil {
		return nil, web.ErrBadRequest("Invalid notification id")
	}

	err = h.service.Delete(id)
	if err != nil {
		return nil, errorResponse(err)
	}

	return (*web.JSONResponse)(nil), nil
}

func validateCreateNotificationRequest(r *http.Request) (*services.NotificationRequest, error) {
	var request services.NotificationRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&request); err != nil {
		return nil, err
	}

	isValidType := request.Type.IsValid()
	if !isValidType {
		return nil, ErrInvalidCardType
	}

	isValidCoin := request.Coin.IsValid()
	if !isValidCoin {
		return nil, ErrInvalidCoinType
	}

	for i, e := range request.Channels {
		if !e.IsValid() {
			request.Channels = append(request.Channels[:i], request.Channels[i+1:]...)
		}
	}

	return &request, nil
}
