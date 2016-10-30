package registry

import (
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/db"

	_ "github.com/erikstmartin/go-testdb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	dbClient, err := db.NewDB("testdb", "")
	assert.Nil(t, err)
	Init("db", "", "", "testdb", "", dbClient)
}

func TestApiRegister(t *testing.T) {
	registry := &Registry{}

	router := gin.New()

	registry.ApiRegister(router)

	pathMix := []string{
		"/registry/v1/namespace",
		"/registry/v1/token",
		"/registry/v1/notifications",
		"/registry/v1/repositories/mine",
		"/registry/v1/repositories/public",
		"/registry/v1/tag/list/:namespace/:image",
		"/registry/v1/manifests/:reference/:namespace/:image",
		"/registry/v1/:namespace/:image/publicity",
		"/registry/v1/manifests/:namespace/:image",
	}

	for _, info := range router.Routes() {
		assert.Contains(t, pathMix, info.Path)
	}
}
