package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func IsUserLogin(ec echo.Context) (string, bool) {
	jwtToken := ec.Get("user")
	if jwtToken == nil {
		return "", false
	}
	oldClaims := jwtToken.(*jwt.Token).Claims.(jwt.MapClaims)
	userId := (oldClaims["id"]).(string)
	return userId, len(userId) == 0
}
