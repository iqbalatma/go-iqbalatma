package exception

import (
	"iqbalatma/go-iqbalatma/app/enum"
	"iqbalatma/go-iqbalatma/utils"
)

func newError(code enum.ResponseCode, defaultMessage string, messages ...string) *utils.HTTPError {
	message := defaultMessage
	if len(messages) > 0 && messages[0] != "" {
		message = messages[0]
	}

	return utils.NewHttpError(message, code)
}
func InvalidAction(messages ...string) *utils.HTTPError {
	return newError(enum.ERR_INVALID_ACTION, "Invalid action")
}

func QueryParameterInvalid(messages ...string) *utils.HTTPError {
	return newError(enum.ERR_INVALID_QUERY_PARAMETER, "Invalid query parameter")
}

func InvalidTokenTypeException(messages ...string) *utils.HTTPError {
	return newError(enum.ERR_AUTHENTICATION, "Invalid token type")
}

func InternalServerError(messages ...string) *utils.HTTPError {
	return newError(enum.ERR_INTERNAL_SERVER_ERROR, "Internal server error")
}

func UnauthorizedException(messages ...string) *utils.HTTPError {
	return newError(enum.ERR_ACTION_UNAUTHORIZED, "Unauthorized")
}
