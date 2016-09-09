package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOptModel(t *testing.T) {
	opt := new(ListOptions)
	assert.NotNil(t, opt)
}
