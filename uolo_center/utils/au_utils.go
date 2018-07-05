package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func GenHttpOnlyCookie(name, value, domain string, maxage time.Duration) *http.Cookie {
	jwt := &http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		Domain:   domain,
		Path:     "/",
		MaxAge:   int(maxage),
	}
	return jwt
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
