package token_store

import (
	"github.com/Dataman-Cloud/rolex/src/plugins/auth/token_store"
	"testing"
)

func TestDecryptEncrypt(T *testing.T) {
	plain := "foobar"
	if string(token_store.Decrypt(token_store.Encrypt([]byte(plain), 6), 6)) != plain {
		T.Error("decrypted error")
	}

	T.Log(token_store.Encrypt([]byte(plain), 6))
	T.Log(token_store.Decrypt(token_store.Encrypt([]byte(plain), 6), 6))
	T.Log(string(token_store.Decrypt(token_store.Encrypt([]byte(plain), 6), 6)))
}
