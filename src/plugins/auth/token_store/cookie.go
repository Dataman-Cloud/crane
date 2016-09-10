package token_store

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/utils/encrypt"

	"github.com/gin-gonic/gin"
)

const (
	CRANE_SESSION_KEY = "CRANE_SESSION_KEY"
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
	value, _ := encrypt.Encrypt(key, cookieValue)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    CRANE_SESSION_KEY,
		Value:   value,
		Expires: time.Now().Add(auth.SESSION_DURATION),
	})

	return nil
}

func (d *Cookie) Get(ctx *gin.Context, token string) (string, error) {
	var cookie *http.Cookie
	var err error

	if cookie, err = ctx.Request.Cookie(CRANE_SESSION_KEY); err != nil {
		return "", ErrCookieNotExist
	}

	decryptedValue, _ := encrypt.Decrypt(key, cookie.Value)
	return strings.SplitN(decryptedValue, ":", 2)[1], nil
}

func (d *Cookie) Del(ctx *gin.Context, token string) error {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    CRANE_SESSION_KEY,
		Value:   "",
		Expires: time.Now().Add(auth.SESSION_DURATION),
	})
	return nil
}
