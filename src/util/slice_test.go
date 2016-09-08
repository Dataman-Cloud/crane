package util

import (
	"testing"
)

func TestStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b"}); ok {
		t.Log("pass")
	} else {
		t.Error("faild")
	}
}
