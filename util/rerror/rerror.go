package rerror

const (
	PARAMETER_ERROR        = "10001-400"
	DEFAULT_ERROR_CODE     = "10000-500"
	ENGINE_OPERATION_ERROR = "10002-500"
	ERROR_CODE_FORBIDDEN   = "10002-403"
)

type RolexError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewRolexError(errorCode string, errorMessage string) RolexError {
	return RolexError{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

func (e RolexError) Error() string {
	return e.ErrorMessage
}
