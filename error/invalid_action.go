package exception

import (
	"iqbalatma/go-iqbalatma/app/enum"
	"net/http"
)

func InvalidAction(opts ...HTTPErrorOption) *HTTPError {
	err := NewHttpError(
		enum.ERR_INVALID_ACTION,
		"Your action is invalid",
		http.StatusForbidden,
	)

	for _, opt := range opts {
		opt(err)
	}

	return err
}
