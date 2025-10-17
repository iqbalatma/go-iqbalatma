package utils

import (
	"fmt"
	"iqbalatma/go-iqbalatma/app/enum"
	"time"
)

type BaseHttpResponse struct {
	Code       enum.ResponseCode `json:"code"`
	Message    string            `json:"message"`
	StatusCode int               `json:"status_code"`
	Timestamp  time.Time         `json:"timestamp"`
	Payload    *Payload          `json:"payload"`
}
type HTTPResponse struct {
	BaseHttpResponse
}

type HTTPError struct {
	BaseHttpResponse
}

func (H HTTPError) Error() string {
	return fmt.Sprintf(
		"message : %s, code: %s, timetamp: %s",
		H.Message,
		H.Code,
		H.Timestamp,
	)
}

type Payload struct {
	Data interface{}     `json:"data"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	LastPage    int    `json:"last_page"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	Total       int64  `json:"total"`
}

func NewHttpSuccess(message string, payload *Payload) *HTTPResponse {
	return &HTTPResponse{
		BaseHttpResponse: BaseHttpResponse{
			Code:       enum.SUCCESS,
			Timestamp:  time.Now(),
			StatusCode: enum.SUCCESS.HTTPStatus(),
			Message:    message,
			Payload:    payload,
		},
	}
}

func NewHttpError(message string, code enum.ResponseCode) *HTTPError {
	return &HTTPError{
		BaseHttpResponse: BaseHttpResponse{
			Code:       code,
			StatusCode: code.HTTPStatus(),
			Timestamp:  time.Now(),
			Message:    message,
			Payload:    nil,
		},
	}
}
