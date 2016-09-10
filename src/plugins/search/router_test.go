package search

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testDocument Document = Document{
	ID:      "1",
	Name:    "test",
	Type:    "test",
	GroupId: uint64(123),
	Param:   map[string]string{"id": "test"},
}

func TestNewDocumentStorage(t *testing.T) {
	ds := NewDocumentStorage()
	assert.Empty(t, ds.Store, "NewDocumentStorage should be empty")
}

func TestEmpty(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	ds.Empty()
	assert.Empty(t, ds.Store, "should be empty")
}

func TestCopy(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	copidDS := ds.Copy()
	assert.Equal(t, copidDS.Store["document"], testDocument, "should be equal")
}

func TestIndices(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	indices := ds.Indices()
	keyList := make([]string, 1)
	keyList[0] = "document"
	assert.Equal(t, indices, keyList, "should be equal")
}

func TestSet(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Set("document", testDocument)
	assert.Equal(t, ds.Store["document"], testDocument, "should be equal")
}

func TestGet(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	value := ds.Get("document")
	assert.Equal(t, value, &testDocument, "should be equal")
}

func TestRegisterApiForSearch(t *testing.T) {
	router := gin.New()
	searchApi := &SearchApi{}
	searchApi.RegisterApiForSearch(router)
	routes := router.Routes()
	assert.Equal(t, len(routes), 1, "there should be one router")
	if len(routes) == 1 {
		route := routes[0]
		assert.Equal(t, route.Method, "GET", "method should be get")
		assert.Equal(t, route.Path, "/search/v1/luckysearch", "route path should be equal")
		assert.Equal(t, route.Handler, "github.com/Dataman-Cloud/crane/src/plugins/search.(*SearchApi).Search-fm", "handler should be equal")
	}

}

func TestIndexData(t *testing.T) {
	searchApi := &SearchApi{}
	searchApi.IndexData()
	assert.Equal(t, len(searchApi.Index), 0, "should be equal")
	assert.Empty(t, searchApi.Store, "should be equal")
	assert.Empty(t, searchApi.PrefetchStore, "should be equal")
	assert.Empty(t, searchApi.Indexer, "should be equal")
}
