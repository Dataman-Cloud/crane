package search

import (
	"sort"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	"github.com/gin-gonic/gin"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

const (
	RESULT_LEN = 10
)

const (
	//Search
	CodeInvalidSearchKeywords = "503-13001"
)

func (searchApi *SearchApi) Search(ctx *gin.Context) {
	query := ctx.Query("keyword")
	if query == "" {
		craneError := cranerror.NewError(CodeInvalidSearchKeywords, "invalid search keywords")
		httpresponse.Error(ctx, craneError)
		return
	}

	//groups, ok := ctx.Get("groups")

	results := []*Document{}
	indexs := fuzzy.RankFind(query, searchApi.Index)
	sort.Sort(indexs)
	if len(indexs) > 0 {
		if len(indexs) > 10 {
			indexs = indexs[:10]
		}
		for _, index := range indexs {
			results = append(results, searchApi.Store.Get(index.Target))
		}
	}
	httpresponse.Ok(ctx, results)
}
