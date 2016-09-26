package registry

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"

	"github.com/docker/distribution/registry/auth/token"
	"github.com/stretchr/testify/assert"
)

func TestParseResourceActionsNone(t *testing.T) {
	scope := ""
	regi := &Registry{}
	res := regi.ParseResourceActions(scope)
	assert.Nil(t, res)
}

func TestParseResourceActions(t *testing.T) {
	regi := &Registry{}
	scope := "type:name:action1,action2"
	var expectedReturn []*token.ResourceActions
	expectedReturn = append(expectedReturn, &token.ResourceActions{
		Type:    "type",
		Name:    "name",
		Actions: strings.Split("action1,action2", ","),
	})

	res := regi.ParseResourceActions(scope)
	assert.Equal(t, expectedReturn, res)
}

func TestFilterAccessFirstIFReturn(t *testing.T) {
	regi := &Registry{}
	a := &token.ResourceActions{
		Type: "registry",
		Name: "catalog",
	}
	regi.FilterAccess("", true, a)
}

func TestFilterAccess(t *testing.T) {
	regi := &Registry{}
	a := &token.ResourceActions{
		Type: "repository",
		Name: "name/image",
	}
	regi.FilterAccess("name@test.com", true, a)
}

func TestMakeTokenNoneKey(t *testing.T) {
	regi := &Registry{}
	var a []*token.ResourceActions
	a = append(a, &token.ResourceActions{
		Type: "repository",
		Name: "name/image",
	})
	rs, err := regi.MakeToken("", "", "", a)
	assert.NotNil(t, err)
	assert.Equal(t, rs, "")
}

func TestMakeTokenCore(t *testing.T) {
	fakePrivateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
	fakefakefake
	-----END RSA PRIVATE KEY-----`)
	tmpfile, err := ioutil.TempFile("", "privatekey")
	assert.Nil(t, err)
	defer os.Remove(tmpfile.Name()) // clean up
	_, err = tmpfile.Write(fakePrivateKey)
	assert.Nil(t, err)
	err = tmpfile.Close()
	assert.Nil(t, err)

	regi := &Registry{}
	var a []*token.ResourceActions
	a = append(a, &token.ResourceActions{
		Type: "repository",
		Name: "name/image",
	})
	_, err = regi.MakeToken(tmpfile.Name(), "name@test.com", "", a)
	assert.NotNil(t, err)
}

func TestRandString(t *testing.T) {
	rb, err := randString(10)
	assert.Nil(t, err)
	assert.Equal(t, len(rb), 10)
}

func TestBase64UrlEncode(t *testing.T) {
	input := []byte("left=right")
	s := base64UrlEncode(input)
	assert.Equal(t, "bGVmdD1yaWdodA", s)
}

func TestRegistryNamespaceForAccount(t *testing.T) {
	a := auth.Account{
		Email: "test@test.com",
	}
	pre := RegistryNamespaceForAccount(a)
	assert.Equal(t, "test", pre)
}

func TestRegistryNamespaceForEmail(t *testing.T) {
	pre := RegistryNamespaceForEmail("test@test.com")
	assert.Equal(t, "test", pre)
}
