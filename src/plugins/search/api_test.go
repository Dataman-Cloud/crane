package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	searchServer *httptest.Server
)

func SetupServer() {
	router := gin.New()

	searchApi := &SearchApi{}
	searchApi.Indexer = &MockCraneIndexer{}
	searchApi.RegisterApiForSearch(router)
	searchServer = httptest.NewServer(router)
}

func CloseServer() {
	searchServer.Close()
}

func TestSearch(t *testing.T) {
	SetupServer()
	defer CloseServer()
	time.Sleep(time.Minute * time.Duration(SEARCH_LOAD_DATA_INTERVAL))
	req, err := http.NewRequest("GET", searchServer.URL+"/search/v1/luckysearch?keyword=blah", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK, "response status code should be equal")

	type ResponseBody struct {
		Code int
		Data []*Document
	}
	var respBody ResponseBody

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("fail to read response body: ", err)
	}
	json.Unmarshal(body, &respBody)
	assert.Equal(t, respBody.Code, httpresponse.CodeOk, "should be equal")
	assert.Equal(t, len(respBody.Data), 1, "should be equal")
}
