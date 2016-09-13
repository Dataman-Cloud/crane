package httpresponse

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/utils/cranerror"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Success struct {
	Code int
	Data string
}

func TestOk(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ok", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/ok", func(c *gin.Context) {
		Ok(c, "success")
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)

	var success Success
	err := json.Unmarshal([]byte(w.Body.String()), &success)
	assert.Nil(t, err)
	assert.Equal(t, success.Code, CodeOk)
	assert.Equal(t, success.Data, "success")
}

func TestCreate(t *testing.T) {
	req, _ := http.NewRequest("POST", "/post", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/post", func(c *gin.Context) {
		Create(c, "success")
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)

	var success Success
	err := json.Unmarshal([]byte(w.Body.String()), &success)
	assert.Nil(t, err)
	assert.Equal(t, success.Code, CodeOk)
	assert.Equal(t, success.Data, "success")
}

func TestDelte(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/delete", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.DELETE("/delete", func(c *gin.Context) {
		Delete(c, "success")
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNoContent)

	var success Success
	err := json.Unmarshal([]byte(w.Body.String()), &success)
	assert.Nil(t, err)
	assert.Equal(t, success.Code, CodeOk)
	assert.Equal(t, success.Data, "success")
}

func TestUpdate(t *testing.T) {
	req, _ := http.NewRequest("PATCH", "/update", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.PATCH("/update", func(c *gin.Context) {
		Update(c, "success")
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusAccepted)

	var success Success
	err := json.Unmarshal([]byte(w.Body.String()), &success)
	assert.Nil(t, err)
	assert.Equal(t, success.Code, CodeOk)
	assert.Equal(t, success.Data, "success")
}

func TestNormalError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/error", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/error", func(c *gin.Context) {
		err := errors.New("nomal test error")
		Error(c, err)
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusServiceUnavailable)

	type ResponseError struct {
		Code    int
		Data    interface{}
		Message string
		Source  string
	}
	var respError ResponseError
	err := json.Unmarshal([]byte(w.Body.String()), &respError)
	assert.Nil(t, err)
	assert.Equal(t, respError.Code, CodeUndefined)
	assert.Equal(t, respError.Message, "nomal test error")
	assert.Equal(t, respError.Source, "docker")
}

func TestCustomError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/error", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/error", func(c *gin.Context) {
		err := cranerror.NewError("401-11111", "custom error")
		Error(c, err)
	})

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 401)

	type ResponseError struct {
		Code    int
		Data    interface{}
		Message string
		Source  string
	}
	var respError ResponseError
	err := json.Unmarshal([]byte(w.Body.String()), &respError)
	assert.Nil(t, err)
	assert.Equal(t, respError.Code, 11111)
	assert.Equal(t, respError.Message, "custom error")
	assert.Equal(t, respError.Source, "crane")
}

func TestSSEventOk(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ok", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/ok", func(c *gin.Context) {
		SSEventOk(c, "test", "success-1")
	})

	r.ServeHTTP(w, req)

	type SseSuccess struct {
		Event string
		Data  Success
	}
	assert.Equal(t, w.Code, http.StatusOK)

	assert.NotNil(t, w.Body.String())
	assert.Equal(t, strings.Replace(w.Body.String(), " ", "", -1), strings.Replace("event:test\ndata:{\"code\":0,\"data\":\"success-1\"}\n\n", " ", "", -1))
}

func TestSSEventNormalError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ok", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/ok", func(c *gin.Context) {
		err := errors.New("sse error")
		SSEventError(c, "test", err)
	})

	r.ServeHTTP(w, req)

	type SseSuccess struct {
		Event string
		Data  Success
	}
	assert.Equal(t, w.Code, http.StatusOK)

	assert.NotNil(t, w.Body.String())
	assert.Equal(t, strings.Replace(w.Body.String(), " ", "", -1), strings.Replace("event:test\ndata:{\"code\":10001,\"data\":{},\"message\":\"sse error\",\"source\":\"docker\"}\n\n", " ", "", -1))
}

func TestSSEventCustomError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ok", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/ok", func(c *gin.Context) {
		err := cranerror.NewError("401-11111", "custom error")
		SSEventError(c, "test", err)
	})

	r.ServeHTTP(w, req)

	type SseSuccess struct {
		Event string
		Data  Success
	}
	assert.Equal(t, w.Code, http.StatusOK)

	assert.NotNil(t, w.Body.String())
	assert.Equal(t, strings.Replace(w.Body.String(), " ", "", -1), strings.Replace("event:test\ndata:{\"code\":11111,\"data\":{},\"message\":\"custom error\",\"source\":\"crane\"}\n\n", " ", "", -1))
}

func TestExtractCraneError(t *testing.T) {
	var err error
	c, ok := extractCraneError(err)
	assert.False(t, ok)
	assert.Nil(t, c)

	var cranerror cranerror.CraneError
	cranerror.Code = "code"
	c, ok = extractCraneError(&cranerror)
	assert.True(t, ok)
	assert.NotNil(t, c)
}

func TestParseHttpCodeAndErrCode(t *testing.T) {
	hCode, errCode := parseHttpCodeAndErrCode("400-10000")
	assert.Equal(t, hCode, 400)
	assert.Equal(t, errCode, 10000)
}
