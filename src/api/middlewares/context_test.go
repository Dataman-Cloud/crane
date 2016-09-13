package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestCraneApiContext(t *testing.T) {
	var craneContextNodeId string
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CraneApiContext())
	router.GET("/:node_id/test", func(c *gin.Context) {
		var craneContext context.Context
		craneContext = c.Value("craneContext").(context.Context)
		craneContextNodeId = craneContext.Value("node_id").(string)
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/fakeNodeId/test", nil)
	router.ServeHTTP(w, r)

	assert.Equal(t, "fakeNodeId", craneContextNodeId)
}
