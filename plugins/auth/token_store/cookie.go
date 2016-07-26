package token_store

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Dataman-Cloud/rolex/plugins/auth"

	"github.com/gin-gonic/gin"
)

const (
	ROLEX_SESSION_KEY     = "ROLEX_SESSION_KEY"
	COOKIE_EXPIRE_TIMEOUT = 10 * time.Minute
)

var (
	ErrCookieNotExist = errors.New("cookie does not exists")
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
var keyText = "astaxie12798akljzmknm.ahkjkljl;k"
var cipherBlock cipher.Block
var err error

func init() {
	cipherBlock, err = aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}
}

type Cookie struct {
	auth.TokenStore
}

func NewCookieStore() *Cookie {
	return &Cookie{}
}

func (d *Cookie) Set(ctx *gin.Context, token, accountId string, expiredAt time.Time) error {
	cookieValue := []byte(fmt.Sprintf("%s:%s", token, accountId))
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    ROLEX_SESSION_KEY,
		Value:   string(Encrypt(cookieValue, len(cookieValue))),
		Expires: time.Now().Add(COOKIE_EXPIRE_TIMEOUT),
	})

	return nil
}

func (d *Cookie) Get(ctx *gin.Context, token string) (string, error) {
	var cookie *http.Cookie
	var err error

	if cookie, err = ctx.Request.Cookie(ROLEX_SESSION_KEY); err != nil {
		return "", ErrCookieNotExist
	}
	decryptedValue := string(Decrypt([]byte(cookie.Value), len(cookie.Value)))
	return strings.SplitN(decryptedValue, ":", 2)[1], nil
}

func (d *Cookie) Del(ctx *gin.Context, token string) error {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    ROLEX_SESSION_KEY,
		Value:   "",
		Expires: time.Now().Add(COOKIE_EXPIRE_TIMEOUT),
	})
	return nil
}

func Encrypt(plain []byte, len int) []byte {
	cfb := cipher.NewCFBEncrypter(cipherBlock, commonIV)
	ciphertext := make([]byte, len)
	cfb.XORKeyStream(ciphertext, plain)
	return ciphertext
}

func Decrypt(cipertext []byte, len int) []byte {
	cfbdec := cipher.NewCFBDecrypter(cipherBlock, commonIV)
	plaintext := make([]byte, len)
	cfbdec.XORKeyStream(plaintext, cipertext)
	return plaintext
}
