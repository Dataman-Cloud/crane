package encrypt

import (
	"testing"
)

func TestencodeBase64(T *testing.T) {
	if result := encodeBase64([]byte("test")); result == "dGVzdA==" {
		T.Log("encodeBase64 pass")
	} else {
		T.Error("encodeBase64 error")
	}
}

func TestdecodeBase64(T *testing.T) {
	if result := decodeBase64("dGVzdA=="); string(result) == "test" {
		T.Log("decodeBase64 pass")
	} else {
		T.Error("decodeBase64 error")
	}
}

func TestEncrypt(T *testing.T) {
	key := "abcdefghijklmnopqrstuvwx"
	if _, err := Encrypt(key, "test"); err == nil {
		T.Log("encrypt pass")
	} else {
		T.Log("encrypt error")
	}
}

func TestDecrypt(T *testing.T) {
	key := "abcdefghijklmnopqrstuvwx"
	if _, err := Decrypt(key, "Dy5QmA=="); err == nil {
		T.Log("decrypt pass")
	} else {
		T.Error("decrypt error")
	}
}
