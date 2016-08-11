package search

import (
	"sort"

	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
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

	groups, ok := ctx.Get("groups")

	results := []Document{}
	indexs := fuzzy.RankFind(query, searchApi.Index)
	sort.Sort(indexs)
	if len(indexs) > 0 {
		if !ok {
			if len(indexs) > 10 {
				indexs = indexs[:10]
			}
			for _, index := range indexs {
				if result, ok := searchApi.Store[index.Target]; ok {
					results = append(results, result)
				}
			}
		} else {
			for _, index := range indexs {
				if result, ok := searchApi.Store[index.Target]; ok {
					switch result.Type {
					case DOCUMENT_STACK, DOCUMENT_SERVICE, DOCUMENT_TASK:
						for _, group := range groups.([]auth.Group) {
							if group.ID == result.GroupId {
								results = append(results, result)
								break
							}
						}
					default:
						results = append(results, result)
					}
				}
				if len(results) >= 10 {
					break
				}
			}
		}
	}
	rolexgin.HttpOkResponse(ctx, results)
}
