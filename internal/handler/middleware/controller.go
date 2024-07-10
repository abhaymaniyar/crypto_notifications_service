package middleware

import (
	"awesomeProject/internal/web"
)

type Controller func(request *web.Request) (*web.JSONResponse, web.ErrorInterface)
