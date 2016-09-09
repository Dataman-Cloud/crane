package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvFile(t *testing.T) {
	tempFileName := filepath.Join(os.TempDir(), "env_file")
	err := ioutil.WriteFile(tempFileName, []byte("FOO=bar"), 0666)
	assert.Nil(t, err)
	defer func() {
		os.Remove(tempFileName)
	}()

	env := os.Getenv("FOO")
	assert.Equal(t, env, "")
	LoadEnvFile(tempFileName)
	assert.NotNil(t, os.Getenv("FOO"))
	assert.Equal(t, os.Getenv("FOO"), "bar")
	defer func() {
		os.Unsetenv("FOO")
	}()
}

func TestParseLn(t *testing.T) {
	keyValueLine := "FOO=bar"
	k, v, err := Parseln(keyValueLine)
	assert.Equal(t, k, "FOO")
	assert.Equal(t, v, "bar")
	assert.Nil(t, err)

	keyValueLineIncorrect := "FOObar"
	k, v, err = Parseln(keyValueLineIncorrect)
	assert.NotNil(t, err)
}
