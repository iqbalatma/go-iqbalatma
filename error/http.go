package exception

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"time"
)

type HTTPErrorOption func(*HTTPError)

func WithMessage(message string) HTTPErrorOption {
	return func(e *HTTPError) {
		e.Message = message
	}
}

func WithStatusCode(statusCode int) HTTPErrorOption {
	return func(e *HTTPError) {
		e.StatusCode = statusCode
	}
}

func WithCode(code enum.ResponseCode) HTTPErrorOption {
	return func(e *HTTPError) {
		e.Code = code
	}
}

type HTTPError struct {
	Message    string            `json:"message"`
	Code       enum.ResponseCode `json:"code"`
	StatusCode int               `json:"status_code"`
	Timestamp  time.Time         `json:"timestamp"`
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("Message: %s,  Code: %s, Status Code: %d", e.Message, e.Code, e.StatusCode)
}

func NewHttpError(code enum.ResponseCode, message string, statusCode int) *HTTPError {
	return &HTTPError{
		Message:    message,
		Code:       code,
		StatusCode: statusCode,
		Timestamp:  time.Now(),
	}
}
