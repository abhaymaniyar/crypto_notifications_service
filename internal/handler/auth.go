package handler

import (
	"awesomeProject/internal/web"
	"awesomeProject/utils"
	"encoding/json"
)

type AuthHandler interface {
	Login(r *web.Request) (*web.JSONResponse, web.ErrorInterface)
}

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func MakeAuthHandler() AuthHandler {
	return &authHandler{}
}

// Login creates a new user
func (h *authHandler) Login(r *web.Request) (*web.JSONResponse, web.ErrorInterface) {
	var req interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, web.ErrBadRequest("Failed to decode request body")
	}

	// call to service

	jsonResponse, err := utils.StructToMap("")
	if err != nil {
		return nil, web.ErrInternalServerError(err.Error())
	}

	return (*web.JSONResponse)(&jsonResponse), nil
}
