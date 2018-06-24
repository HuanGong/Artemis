package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	// 如果是登陆过的用户，种的cookie为Uolo用户ID, 如果从未登陆过，访问过的echoface.cn 任意域名下的服务，
	// 则种下一个common的uuid
	kUoloCid          = "_uolocid"
	kUoloCookieMaxAge = time.Hour * 24 * 30
)

func SetUoloUserCookie(ec echo.Context, userId string) {
	uidCookie := &http.Cookie{
		Name:     kUoloCid,
		Value:    userId,
		HttpOnly: true,
		Domain:   ".echoface.cn",
		MaxAge:   int(kUoloCookieMaxAge),
	}
	ec.SetCookie(uidCookie)
}

func IsUserLogin(ec echo.Context) (string, bool) {
	jwtToken := ec.Get("user")

	if jwtToken == nil {
		return "", false
	}
	oldClaims := jwtToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userId, _ := oldClaims["id"]
	id := userId.(string)
	return id, len(id) != 0
}

func SetUoloCookieIdentifierIfNeeded(ec echo.Context) {

	cookie, err := ec.Cookie(kUoloCid)

	if err != nil { // SetCookie
		newCookie := &http.Cookie{
			Name:     kUoloCid,
			Value:    "",
			HttpOnly: true,
			Domain:   ".echoface.cn",
			MaxAge:   int(kUoloCookieMaxAge),
		}

		if uoloUserId, login := IsUserLogin(ec); login {
			newCookie.Value = uoloUserId
		} else {
			uuidGen, _ := uuid.NewV4()
			newCookie.Value = uuidGen.String()
		}
		ec.SetCookie(newCookie)
		return
	}
	logrus.Debugln("Got UoloCookie Identifier Id", cookie.Value)
	ec.Set(kUoloCid, cookie.Value)
}

func ServeCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		SetUoloCookieIdentifierIfNeeded(c)
		return next(c)
	}
}
