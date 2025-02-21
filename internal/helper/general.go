package helper

import (
	"isolusi/internal/model"
	"math/rand"
)

func JsonResponse(code int, message, error string, data interface{}) model.Response {
	meta := model.Meta{
		Code:    code,
		Message: message,
		Error:   error,
	}

	response := model.Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func GenerateRand(n int64) int {
	result := rand.New(rand.NewSource(n))
	return result.Int()
}
