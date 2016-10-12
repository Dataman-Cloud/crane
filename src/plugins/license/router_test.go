package license

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestApiRouter(t *testing.T) {
	licenseApi := &LicenseApi{}

	router := gin.New()

	licenseApi.ApiRegister(router)

	pathMix := []string{
		"/license/v1/license",
	}

	for _, info := range router.Routes() {
		assert.Contains(t, pathMix, info.Path)
	}
}
