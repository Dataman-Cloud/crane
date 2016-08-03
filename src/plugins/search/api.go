package search

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

const (
	RESULT_LEN = 10
)

func (searchApi *SearchApi) Search(ctx *gin.Context) {
	query := ctx.Query("keyword")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "invalid search keyword"})
		return
	}

	results := []Document{}
	indexs := fuzzy.RankFind(query, searchApi.Index)
	sort.Sort(indexs)
	if len(indexs) > 0 {
		if len(indexs) > 10 {
			indexs = indexs[:10]
		}
		for _, index := range indexs {
			if result, ok := searchApi.Store[index.Target]; ok {
				results = append(results, result)
			}
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": results})
}
