package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dataman-Cloud/crane/src/dockerclient"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	"github.com/stretchr/testify/assert"
)

type Success struct {
	Code int
	Data []*RouteInfo
}

func TestHelp(t *testing.T) {
	api := &Api{
		Client: &dockerclient.CraneDockerClient{},
	}
	router := api.ApiRouter()

	req, _ := http.NewRequest("GET", "/misc/v1/help", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)

	var success Success
	err := json.Unmarshal([]byte(w.Body.String()), &success)
	assert.Nil(t, err)
	assert.Equal(t, success.Code, cranerror.CodeOk)
}
