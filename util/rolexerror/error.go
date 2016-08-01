package rolexerror

type RolexError struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

func NewRolexError(code int, message string) error {
	return &RolexError{
		Code:    code,
		Message: message,
	}
}

func (e RolexError) Error() string {
	return e.Message
}
