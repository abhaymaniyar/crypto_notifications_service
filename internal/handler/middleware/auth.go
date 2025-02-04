package middleware

import (
	"awesomeProject/internal/web"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func AuthMiddleware(next Controller) Controller {
	return func(r *web.Request) (*web.JSONResponse, web.ErrorInterface) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return nil, web.ErrUnauthorizedRequest("Missing authorization header")
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte("SecretKey"), nil
		})

		if err != nil {
			return nil, web.ErrUnauthorizedRequest("Invalid token")
		}

		if !token.Valid {
			return nil, web.ErrUnauthorizedRequest("Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "userID", claims["user_id"])
			r.Request = r.Request.WithContext(ctx)
		} else {
			return nil, web.ErrUnauthorizedRequest("Invalid token")
		}

		return next(r)
	}
}
