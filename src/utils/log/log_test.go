package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestInit(t *testing.T) {
	assert.NotNil(t, L)
}

func TestWithLogger(t *testing.T) {
	c := context.Background()
	ctx := WithLogger(c, L)
	assert.NotNil(t, ctx)
	loggerWithinCtx := GetLogger(ctx)
	assert.NotNil(t, loggerWithinCtx)
}
