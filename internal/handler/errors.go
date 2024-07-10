package handler

import (
	"awesomeProject/internal/web"
	"errors"
	"fmt"
)

var (
	ErrInvalidCardType    = errors.New("invalid card type")
	ErrInvalidCoinType    = errors.New("invalid coin type")
	ErrInvalidChannelType = errors.New("invalid channel type")
)

func errorResponse(err error) web.ErrorInterface {
	if errors.Is(err, ErrInvalidCardType) ||
		errors.Is(err, ErrInvalidCoinType) ||
		errors.Is(err, ErrInvalidChannelType) {
		return web.ErrBadRequest(err.Error())
	}

	return web.ErrInternalServerError(fmt.Sprintf("error while processing request : %s", err))
}
