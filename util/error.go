package util

type StatusForbiddenError struct {
	s string
}

func NewStatusForbiddenError(text string) error {
	return &StatusForbiddenError{text}
}

func (e StatusForbiddenError) Error() string {
	return e.s
}
