package cranerror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError(CodeErrorUpdateNodeMethod, "foobar")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "foobar")
}

func TestContainerStatsStopError(t *testing.T) {
	err := errors.New("foobar")

	var containerStatsStopError ContainerStatsStopError
	containerStatsStopError.ID = "container-id"
	containerStatsStopError.Err = err

	assert.Equal(t, "foobar", containerStatsStopError.Error())
}

func TestNodeConnError(t *testing.T) {
	err := errors.New("foobar")
	var nodeConnError NodeConnError
	nodeConnError.ID = "node-id"
	nodeConnError.Endpoint = "node-endpoint"
	nodeConnError.Err = err

	assert.Equal(t, "foobar", nodeConnError.Error())
}

func TestServicePortConflictError(t *testing.T) {
	err := errors.New("foobar")
	var servicePortConflictError ServicePortConflictError
	servicePortConflictError.Name = "name"
	servicePortConflictError.Namespace = "namespace"
	servicePortConflictError.Err = err

	assert.Equal(t, "foobar", servicePortConflictError.Error())
}
