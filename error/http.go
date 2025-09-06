package exception

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"time"
)

type HTTPErrorOption func(*HTTPError)

type HTTPError struct {
	Message    string            `json:"message"`
	Code       enum.ResponseCode `json:"code"`
	StatusCode int               `json:"status_code"`
	Timestamp  time.Time         `json:"timestamp"`
}

func (e HTTPError) Error() string {
	return fmt.Sprintf(
		"Message: %s,  Code: %s, Status Code: %d, Timestamp: %s",
		e.Message,
		e.Code,
		e.StatusCode,
		e.Timestamp,
	)
}

func NewHttpError(code enum.ResponseCode, message string, statusCode int) *HTTPError {
	return &HTTPError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Timestamp:  time.Now(),
	}
}
