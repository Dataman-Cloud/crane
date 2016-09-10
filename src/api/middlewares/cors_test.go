package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOptionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(OptionHandler())
	router.OPTIONS("/cors/test", func(c *gin.Context) {
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("OPTIONS", "/cors/test", nil)
	router.ServeHTTP(w, r)

	assert.Equal(t, "*", w.HeaderMap.Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "true", w.HeaderMap.Get("Access-Control-Allow-Credentials"))
	assert.Equal(t, "GET, POST, PUT, DELETE, OPTIONS, PATCH", w.HeaderMap.Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Depth, User-Agent, X-File-Size, X-Requested-With, X-Requested-By, If-Modified-Since, X-File-Name, Cache-Control, X-XSRFToken, Authorization", w.HeaderMap.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "application/json", w.HeaderMap.Get("Content-Type"))
	assert.Equal(t, 204, w.Code)
}
