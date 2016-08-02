package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestConfigFeatureEnabled(t *testing.T) {
	config := &Config{
		FeatureFlags: []string{"foo", "bar"},
	}

	if config.FeatureEnabled("foo") {
		t.Log("feature foo enabled")
	} else {
		t.Error("feature foo should enabled")
	}
}

func TestReadNodeAddrFromFile(t *testing.T) {
	content := []byte("test=127.0.0.1")
	fileName := "node_addr_tmp"
	tmpfile, err := ioutil.TempFile("./", fileName)
	if err != nil {
		t.Fatal("create temp file got err: ", err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	ReadNodeAddrFromFile(tmpfile.Name())

	value, ok := NodeAddrMap["test"]
	if !ok {
		t.Fatal("can not get test ip")
	}

	if value != "127.0.0.1" {
		t.Fatal("get ip wrong")
	}
}
