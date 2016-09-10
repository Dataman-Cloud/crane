package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListIntercept(t *testing.T) {
	var limit uint64
	var offset uint64
	var filters string

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ListIntercept())
	router.GET("/listIntercept/test", func(c *gin.Context) {
		listOptions := c.Value("listOptions").(model.ListOptions)
		limit = uint64(listOptions.Limit)
		offset = uint64(listOptions.Offset)
		filters = listOptions.Filter["fakeFilter"].(string)
		c.String(200, "OK")
	})

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/listIntercept/test?page=1&per_page=10&fakeFilter=fakeValue", nil)
	router.ServeHTTP(w, r)

	assert.Equal(t, uint64(1), limit)
	assert.Equal(t, uint64(9), offset)
	assert.Equal(t, "fakeValue", filters)
}
