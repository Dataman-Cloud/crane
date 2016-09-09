package token_store

import (
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDefaultTokenStore(t *testing.T) {
	d := NewDefaultStore()

	// token expired test test case
	err := d.Set(new(gin.Context), "token", "1", time.Now().Add(-auth.SESSION_DURATION))
	assert.Equal(t, err, nil, "should be equal")

	id, err := d.Get(new(gin.Context), "token")
	assert.Equal(t, err, TokenExpired, "should be equal")

	// normal test case
	err = d.Set(new(gin.Context), "token", "1", time.Now().Add(auth.SESSION_DURATION))
	assert.Equal(t, err, nil, "should be equal")

	id, err = d.Get(new(gin.Context), "token")
	assert.Equal(t, err, nil, "should be equal")
	assert.Equal(t, id, "1", "should be equal")

	err = d.Del(new(gin.Context), "token")
	assert.Equal(t, err, nil, "should be equal")

	// token not found test case
	id, err = d.Get(new(gin.Context), "token")
	assert.Equal(t, id, "", "should be equal")
	assert.Equal(t, err, TokenNotFound, "should be equal")
}
