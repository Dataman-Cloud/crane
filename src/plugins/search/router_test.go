package search

import (
	"reflect"
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"

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

func TestInit(t *testing.T) {
	Init()
	defer delete(apiplugin.ApiPlugins, apiplugin.Search)

	assert.NotNil(t, apiplugin.ApiPlugins[apiplugin.Search])
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

func TestApiRegister(t *testing.T) {
	router := gin.New()
	searchApi := &SearchApi{}
	searchApi.ApiRegister(router)
	routes := router.Routes()
	assert.Equal(t, len(routes), 1, "there should be one router")
	if len(routes) == 1 {
		route := routes[0]
		assert.Equal(t, route.Method, "GET", "method should be get")
		assert.Equal(t, route.Path, "/search/v1/luckysearch", "route path should be equal")
		assert.Equal(t, route.Handler, "github.com/Dataman-Cloud/crane/src/plugins/search.(*SearchApi).Search-fm", "handler should be equal")
	}

}

type MockCraneIndexer struct {
	Indexer
}

var nodeDocument = Document{
	ID:   "24ifsmvkjbyhk",
	Name: "bf3067039e47",
	Type: DOCUMENT_NODE,
	Param: map[string]string{
		"NodeId": "24ifsmvkjbyhk",
	},
}
var networkDocument = Document{
	Name: "blah",
	ID:   "8dfafdbc3a40",
	Type: DOCUMENT_NETWORK,
	Param: map[string]string{
		"NodeId":    "24ifsmvkjbyhk",
		"NetworkID": "8dfafdbc3a40",
	},
}

func (mockIndexer *MockCraneIndexer) Index(store *DocumentStorage) {
	store.Set("24ifsmvkjbyhkbf3067039e47", nodeDocument)
	store.Set("8dfafdbc3a40blah24ifsmvkjbyhk", networkDocument)

}

func TestIndexData(t *testing.T) {
	searchApi := &SearchApi{}
	searchApi.Indexer = &MockCraneIndexer{}
	searchApi.IndexData()
	time.Sleep(time.Minute * time.Duration(SEARCH_LOAD_DATA_INTERVAL) * 1)
	value1 := searchApi.PrefetchStore.Get("24ifsmvkjbyhkbf3067039e47")
	value2 := searchApi.Store.Get("8dfafdbc3a40blah24ifsmvkjbyhk")
	assert.Equal(t, len(searchApi.Index), 2, "should be equal")
	assert.True(t, reflect.DeepEqual(value1, &nodeDocument), "should be equal")
	assert.True(t, reflect.DeepEqual(value2, &networkDocument), "should be equal")
}
