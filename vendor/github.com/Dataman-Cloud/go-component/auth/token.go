package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	DEFAULT_SALT = "thequickbrownfoxjumpdog"
)

func GenToken(a *Account) string {
	hash := md5.New()
	hash.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", a.Email, a.Password, a.LoginAt.Format(time.RFC3339), DEFAULT_SALT)))

	return hex.EncodeToString(hash.Sum(nil))
}
