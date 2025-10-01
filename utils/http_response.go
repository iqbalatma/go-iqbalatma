package utils

import (
	"iqbalatma/go-iqbalatma/app/enum"
	"time"
)

type HTTPResponse struct {
	Code      enum.ResponseCode `json:"code"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Payload   *Payload          `json:"payload"`
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

type ResponseOption struct {
	HTTPStatusCode int
	Code           enum.ResponseCode
}
