package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTransport(t *testing.T) {
	assert.NotNil(t, DefaultTransport())
}

func TestDefaultPooledTransport(t *testing.T) {
	assert.NotNil(t, DefaultPooledTransport())
}
