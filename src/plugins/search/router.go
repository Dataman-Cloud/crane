package search

import (
	"time"

	"github.com/Dataman-Cloud/crane/src/plugins/apiplugin"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

const SEARCH_LOAD_DATA_INTERVAL = 1

type Indexer interface {
	Index(prefetchStore *DocumentStorage)
}

type DocumentStorage struct {
	Store map[string]Document
}

func NewDocumentStorage() *DocumentStorage {
	return &DocumentStorage{Store: make(map[string]Document)}
}

func (storage *DocumentStorage) Empty() {
	storage.Store = make(map[string]Document)
}

func (storage *DocumentStorage) Copy() *DocumentStorage {
	copied := NewDocumentStorage()
	for k, v := range storage.Store {
		copied.Store[k] = v
	}
	return copied
}

func (storage *DocumentStorage) Indices() []string {
	indices := make([]string, 0)
	for i, _ := range storage.Store {
		indices = append(indices, i)
	}
	return indices
}

func (storage *DocumentStorage) Set(key string, doc Document) {
	storage.Store[key] = doc
}

func (storage *DocumentStorage) Get(key string) *Document {
	doc := storage.Store[key]
	return &doc
}

type SearchApi struct {
	Index         []string
	Store         *DocumentStorage
	PrefetchStore *DocumentStorage

	Indexer Indexer
}

type Document struct {
	ID      string
	Name    string
	Type    string
	GroupId uint64 `json:"-"`
	Param   map[string]string
}

func Init() {
	log.Infof("begin to init and enable plugin: %s", apiplugin.Search)

	apiPlugin := &apiplugin.ApiPlugin{
		Name:         apiplugin.Search,
		Dependencies: []string{},
		Instance:     &SearchApi{},
	}

	apiplugin.Add(apiPlugin)
	log.Infof("init and enable plugin: %s success", apiplugin.Search)
}

func (searchApi *SearchApi) ApiRegister(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	searchApi.IndexData()

	searchV1 := router.Group("/search/v1", middlewares...)
	{
		searchV1.GET("/luckysearch", searchApi.Search)
	}
}

func (searchApi *SearchApi) IndexData() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				searchApi.IndexData()
			}
		}()

		searchApi.PrefetchStore = NewDocumentStorage()
		searchApi.Store = NewDocumentStorage()

		for {
			searchApi.PrefetchStore = NewDocumentStorage()
			searchApi.Indexer.Index(searchApi.PrefetchStore)

			searchApi.Index = searchApi.PrefetchStore.Indices()
			searchApi.Store = searchApi.PrefetchStore.Copy()

			time.Sleep(time.Minute * time.Duration(SEARCH_LOAD_DATA_INTERVAL))
		}
	}()
}
