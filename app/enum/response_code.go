package enum

type ResponseCode string

const (
	SUCCESS                      ResponseCode = "SUCCESS"
	CREATED                      ResponseCode = "CREATED"
	ERR_ENTITY_STILL_USED        ResponseCode = "ERR_ENTITY_STILL_USED"
	ERR_USER_HAS_NO_ROLE         ResponseCode = "ERR_USER_HAS_NO_ROLE"
	ERR_UNIQUENESS               ResponseCode = "ERR_UNIQUENESS"
	ERR_NOT_FOUND                ResponseCode = "ERR_NOT_FOUND"
	ERR_ACTION_UNAUTHORIZED      ResponseCode = "ERR_ACTION_UNAUTHORIZED"
	ERR_INVALID_ACTION           ResponseCode = "ERR_INVALID_ACTION"
	ERR_FORBIDDEN                ResponseCode = "ERR_FORBIDDEN"
	ERR_FORBIDDEN_IP             ResponseCode = "ERR_FORBIDDEN_IP"
	ERR_AUTHENTICATION           ResponseCode = "ERR_AUTHENTICATION"
	ERR_PROFILE_NOT_COMPLETE     ResponseCode = "ERR_PROFILE_NOT_COMPLETE"
	ERR_INSUFFICIENT_QUOTA       ResponseCode = "ERR_INSUFFICIENT_QUOTA"
	ERR_INCOMPLETE_CONFIGURATION ResponseCode = "ERR_INCOMPLETE_CONFIGURATION"
	ERR_UNAUTHENTICATED          ResponseCode = "ERR_UNAUTHENTICATED"
	ERR_UNAUTHORIZED             ResponseCode = "ERR_UNAUTHORIZED"
	ERR_MISSING_REQUIRED_HEADER  ResponseCode = "ERR_MISSING_REQUIRED_HEADER"
	ERR_VALIDATION               ResponseCode = "ERR_VALIDATION"
	ERR_ENTITY_NOT_FOUND         ResponseCode = "ERR_ENTITY_NOT_FOUND"
	ERR_INTERNAL_SERVER_ERROR    ResponseCode = "ERR_INTERNAL_SERVER_ERROR"
	ERR_BAD_REQUEST              ResponseCode = "ERR_BAD_REQUEST"
	ERR_UNKNOWN                  ResponseCode = "ERR_UNKNOWN"
)
