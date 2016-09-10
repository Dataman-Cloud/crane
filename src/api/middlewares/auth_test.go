package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	authGroup := router.Group("/auth/", Authorization)
	{
		authGroup.GET("/network/access", AuthorizeNetworkAccess(auth.PermReadOnly), func(c *gin.Context) {
			c.String(200, "OK")
		})
		authGroup.GET("/service/access", AuthorizeServiceAccess(auth.PermReadOnly), func(c *gin.Context) {
			c.String(200, "OK")
		})
	}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/auth/network/access", nil)
	router.ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)

	w1 := httptest.NewRecorder()
	r1, _ := http.NewRequest("GET", "/auth/service/access", nil)
	router.ServeHTTP(w1, r1)
	assert.Equal(t, 200, w.Code)
}
