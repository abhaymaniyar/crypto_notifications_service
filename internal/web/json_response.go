package web

import (
	"awesomeProject/internal/logger"
	"context"
	"encoding/json"
)

type JSONResponse map[string]interface{}

func (r *JSONResponse) ByteArray(ctx context.Context) []byte {
	encodedData, encodeErr := json.Marshal(r)
	if encodeErr != nil {
		logger.E(ctx, encodeErr, "unable to marshal json response",
			logger.Field("error", encodeErr.Error()))

		return nil
	}

	return encodedData
}
