package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateOptionsStruct(t *testing.T) {
	var opts UpdateOptions
	opts.Method = "foobar"
	assert.Equal(t, "foobar", opts.Method)
}

func TestJoiningNode(t *testing.T) {
	var node JoiningNode
	node.Endpoint = "foobar"
	assert.Equal(t, "foobar", node.Endpoint)
}
