package utils

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func SignCookieForAuth(ec echo.Context, jwt string, hours time.Duration) {
	auCookie := &http.Cookie{
		Name:     "_au",
		Value:    jwt,
		HttpOnly: true,
		Domain:   "echoface.cn",
		Expires:  time.Now().Add(hours * time.Hour),
	}
	ec.SetCookie(auCookie)
}

func SignCookieForUserId(ec echo.Context, uid string, hours time.Duration) {

	uidCookie := &http.Cookie{
		Name:     "_uuid",
		Value:    uid,
		HttpOnly: true,
		Domain:   "echoface.cn",
		Expires:  time.Now().Add(hours * time.Hour),
	}
	ec.SetCookie(uidCookie)
}
