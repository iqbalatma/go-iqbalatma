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

func QuotaExceedException() *HTTPError {
	return NewHttpError(
		enum.ERR_INSUFFICIENT_QUOTA,
		"Quota has been exceeded",
		http.StatusTooManyRequests,
	)
}

func QueryParameterInvalid(messages ...string) *HTTPError {
	message := "Invalid query parameter"
	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}
	return NewHttpError(
		enum.ERR_INVALID_QUERY_PARAMETER,
		message,
		http.StatusBadRequest,
	)
}

func InvalidTokenTypeException(messages ...string) *HTTPError {
	message := "Invalid token type"
	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}
	return NewHttpError(
		enum.ERR_AUTHENTICATION,
		message,
		http.StatusUnauthorized,
	)
}

func InternalServerError(messages ...string) *HTTPError {
	message := "Internal server error"
	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}
	return NewHttpError(
		enum.ERR_INTERNAL_SERVER_ERROR,
		message,
		http.StatusInternalServerError,
	)
}

func UnauthorizedException(messages ...string) *HTTPError {
	message := "Unauthorized"
	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}
	return NewHttpError(
		enum.ERR_ACTION_UNAUTHORIZED,
		message,
		http.StatusUnauthorized,
	)
}
