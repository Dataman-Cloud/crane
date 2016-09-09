package cranerror

import "errors"

const (
	// Default ok code
	CodeOk = 0

	//Default error code
	CodeUndefined = "503-10001"

	//Network error code

	//Node error code
	CodeErrorUpdateNodeMethod = "503-11302"
	CodeErrorNodeRole         = "503-11303"
	CodeErrorNodeAvailability = "503-11304"
	CodeGetNodeInfoError      = "503-11305"
)

type CraneError struct {
	Code string `json:"Code"`
	Err  error  `json:"Err"`
}

func NewError(code string, message string) error {
	return &CraneError{
		Code: code,
		Err:  errors.New(message),
	}
}

func (e *CraneError) Error() string {
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

type ServicePortConflictError struct {
	Name          string
	Namespace     string
	PublishedPort string
	Err           error
}

func (e *ServicePortConflictError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Namespace + ":" + e.Name + " has been published port: " + e.PublishedPort
}
