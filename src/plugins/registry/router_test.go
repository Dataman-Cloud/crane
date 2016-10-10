package registry

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRouter(t *testing.T) {
	registry := &Registry{}

	router := gin.New()

	registry.ApiRegister(router)

	pathMix := []string{
		"/registry/v1/token",
		"/registry/v1/notifications",
		"/registry/v1/repositories/mine",
		"/registry/v1/repositories/public",
		"/registry/v1/tag/list/:namespace/:image",
		"/registry/v1/manifests/:reference/:namespace/:image",
		"/registry/v1/:namespace/:image/publicity",
	}

	for _, info := range router.Routes() {
		assert.Contains(t, pathMix, info.Path)
	}
}
