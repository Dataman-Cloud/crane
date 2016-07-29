package token_store

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Dataman-Cloud/rolex/plugins/auth"

	"github.com/gin-gonic/gin"
)

const (
	ROLEX_SESSION_KEY = "ROLEX_SESSION_KEY"
)

var (
	ErrCookieNotExist = errors.New("cookie does not exists")
)

var key = "abcdefghijklmnopqrstuvwx"

type Cookie struct {
	auth.TokenStore
}

func NewCookieStore() *Cookie {
	return &Cookie{}
}

func (d *Cookie) Set(ctx *gin.Context, token, accountId string, expiredAt time.Time) error {
	cookieValue := fmt.Sprintf("%s:%s", token[0:10], accountId)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    ROLEX_SESSION_KEY,
		Value:   Encrypt(key, cookieValue),
		Expires: time.Now().Add(auth.SESSION_DURATION),
	})

	return nil
}

func (d *Cookie) Get(ctx *gin.Context, token string) (string, error) {
	var cookie *http.Cookie
	var err error

	if cookie, err = ctx.Request.Cookie(ROLEX_SESSION_KEY); err != nil {
		return "", ErrCookieNotExist
	}

	decryptedValue := Decrypt(key, cookie.Value)
	return strings.SplitN(decryptedValue, ":", 2)[1], nil
}

func (d *Cookie) Del(ctx *gin.Context, token string) error {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    ROLEX_SESSION_KEY,
		Value:   "",
		Expires: time.Now().Add(auth.SESSION_DURATION),
	})
	return nil
}

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}

func Decrypt(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}
