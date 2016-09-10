package catalog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCatalog(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)

	catalogApi := NewCatalog(fakeDb.Db)
	assert.Equal(t, catalogApi.DbClient, fakeDb.Db)
}
