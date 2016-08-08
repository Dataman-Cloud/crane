package search

import (
	"sort"

	"github.com/Dataman-Cloud/rolex/src/util/rolexerror"
	"github.com/Dataman-Cloud/rolex/src/util/rolexgin"

	"github.com/gin-gonic/gin"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

const (
	RESULT_LEN = 10
)

func (searchApi *SearchApi) Search(ctx *gin.Context) {
	query := ctx.Query("keyword")
	if query == "" {
		rerror := rolexerror.NewRolexError(rolexerror.CodeInvalidSearchKeywords, "invalid search keywords")
		rolexgin.HttpErrorResponse(ctx, rerror)
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
	rolexgin.HttpOkResponse(ctx, results)
}
