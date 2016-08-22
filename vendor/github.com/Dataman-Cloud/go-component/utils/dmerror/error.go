package dmerror

import (
	"errors"
)

type DmError struct {
	Code string `json:"Code"`
	Err  error  `json:"Err"`
}

func NewError(code string, message string) error {
	return &DmError{
		Code: code,
		Err:  errors.New(message),
	}
}

func (e *DmError) Error() string {
	return e.Err.Error()
}
