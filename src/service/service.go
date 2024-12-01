package service

import "github.com/uchupx/dating-api/pkg/errors"

func serviceError(httpCode int, message error) *errors.ErrorMeta {
	return &errors.ErrorMeta{
		HTTPCode: httpCode,
		Message:  message.Error(),
	}
}
