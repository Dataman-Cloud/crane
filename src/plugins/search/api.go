package search

import (
	"errors"
	"sort"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	"github.com/gin-gonic/gin"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

const (
	RESULT_LEN = 10
)

type SearchClient struct {
	Index []string
	Store map[string]Document
}

type Document struct {
	ID      string
	Name    string
	Type    string
	GroupId uint64 `json:"-"`
	Param   map[string]string
}

func (searchApi *SearchApi) Search(ctx *gin.Context) {
	query := ctx.Query("keyword")
	results, err := searchClient.SearchResult(query)
	if err != nil {
		rerror := rolexerror.NewRolexError(rolexerror.CodeInvalidSearchKeywords, err.Error())
		rolexgin.HttpErrorResponse(ctx, rerror)
		return
	}

	rolexgin.HttpOkResponse(ctx, results)
}

func (searchClient *SearchClient) SearchResult(query string) ([]Document, error) {
	if query == "" {
		return nil, errors.New("invalid search query")
	}
	results := []Document{}
	indexs := fuzzy.RankFind(query, searchClient.Index)
	sort.Sort(indexs)
	if len(indexs) > 0 {
		if len(indexs) > 10 {
			indexs = indexs[:10]
		}
		for _, index := range indexs {
			if result, ok := searchClient.Store[index.Target]; ok {
				results = append(results, result)
			}
		}
	}
	return results, nil
}

func (searchClient *SearchClient) StoreData(index string, document Document) {
	searchClient.Index = append(searchClient.Index, index)
	searchClient.Store[index] = document
}
