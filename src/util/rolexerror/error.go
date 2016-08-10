package rolexerror

import (
	"errors"
)

type RolexError struct {
	Code int   `json:"Code"`
	Err  error `json:"Err"`
}

func NewRolexError(code int, message string) error {
	return &RolexError{
		Code: code,
		Err:  errors.New(message),
	}
}

func (e *RolexError) Error() string {
	return e.Err.Error()
}

type ContainerStatsStopError struct {
	ID  string
	Err error
}

func (e *ContainerStatsStopError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return "normal stop" + e.ID
}

type NodeConnError struct {
	ID       string
	Endpoint string
	Err      error
}

func (e *NodeConnError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.ID + " : " + e.Endpoint + " conn error"
}
