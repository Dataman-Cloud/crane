package utils

import (
	"testing"
)

func TestSuccessStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b"}); ok {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}

func TestFaildStringInSlice(t *testing.T) {
	if ok := StringInSlice("c", []string{"a", "b"}); ok {
		t.Error("faild")
	} else {
		t.Log("pass")
	}
}
