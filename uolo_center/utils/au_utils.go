package utils

import (
	"artemis/uolo_center/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"time"
)

var (
	B64CaptchaConfig = base64Captcha.ConfigCharacter{
		Height:             42,
		Width:              200,
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		IsUseSimpleFont:    true,
		CaptchaLen:         4,
	}
)

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

func SetJwtAuthCookie(ec echo.Context, token, domain string) {
	jwt := &http.Cookie{
		Name:     model.JwtCookieAuthName,
		Value:    token,
		HttpOnly: true,
		Domain:   domain,
		Path:     "/",
		MaxAge:   int(time.Hour * 24),
	}
	ec.SetCookie(jwt)
}

func SetUoloUserCookieIfNeeded(ec echo.Context, id, domain string) {
	cookie, err := ec.Cookie(model.UoloCidName)
	if err != nil || cookie.Value != id {
		newCookie := &http.Cookie{
			Name:     model.UoloCidName,
			Value:    id,
			HttpOnly: true,
			Domain:   domain,
			Path:     "/",
			MaxAge:   int(time.Hour * 24 * 30),
		}
		ec.SetCookie(newCookie)
	}
}

func GenB64CaptchaImage() (string, string) {
	idKeyC, capC := base64Captcha.GenerateCaptcha("", B64CaptchaConfig)
	return idKeyC, base64Captcha.CaptchaWriteToBase64Encoding(capC)
}

func GenPngCaptchaImage() (string, []byte) {
	idKeyC, capC := base64Captcha.GenerateCaptcha("", B64CaptchaConfig)
	return idKeyC, capC.BinaryEncodeing()
}
