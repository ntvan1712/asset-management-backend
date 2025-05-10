package error_app

type ErrorResponse struct {
	ErrorCode      string                 `json:"error_code"`
	Message        string                 `json:"message"`
}

const (
	notFoundErrCode       = "not-found-error"
	unauthorizedErrCode   = "unauthorized-error"
	internalServerErrCode = "internal-server-error"
	badRequestErrCode     = "bad-request-error"
)

func NewErrorResponse(errorCode string, message string) *ErrorResponse {
	return &ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func BadRequestErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(badRequestErrCode, message)
}

func NotFoundErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(notFoundErrCode, message)
}

func UnauthorizedErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(unauthorizedErrCode, message)
}

func InternalServerErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(internalServerErrCode, message)
}

func ConflictErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrConflict.Error(), message)
}

func PermissionDeniedErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrPermissionDenied.Error(), message)
}

func InvalidAccessTokenErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrInvalidAccessToken.Error(), message)
}

func ExpiredJwtTokenErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrExpiredAccessToken.Error(), message)
}

func PremiumRequiredErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrPremiumRequired.Error(), message)
}

func UserNotFoundErrorResponse(message string) *ErrorResponse {
	return NewErrorResponse(ErrUserNotFound.Error(), message)
}
