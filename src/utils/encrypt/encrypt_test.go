package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	var key = "abcdefghijklmnopqrstuvwx"
	encrypted, _ := Encrypt(key, "foobar")
	decrypted, _ := Decrypt(key, encrypted)
	assert.Equal(t, decrypted, "foobar")
}
