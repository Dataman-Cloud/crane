package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDb(t *testing.T) {
	assert.Error(t, initDb("foobar", "ff"))
}

func TestInitDbWithDsnError(t *testing.T) {
	assert.Error(t, initDb("mysql", "fake-dsn"))
}
