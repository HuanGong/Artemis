package uolo_center

import (
	"artemis/uolo_center/utils"
	"bytes"
	"github.com/labstack/echo"
	"github.com/qpliu/qrencode-go/qrencode"
	"github.com/satori/go.uuid"
	"image/png"
	"net/http"
	"time"
)

type (
	UtilsHandler struct {
	}
)

func NewUtilsHandler() *UtilsHandler {
	return &UtilsHandler{}
}

func (handler *UtilsHandler) GenQrcode(ec echo.Context) error {
	//rawImage, _ := qrcode2.Encode("http://www.baidu.comadfads", qrcode2.High, 256)
	content := ec.QueryParam("text")
	if len(content) == 0 {
		content = "http://www.echoface.cn"
	}

	grid, _ := qrencode.Encode(content, qrencode.ECLevelH)
	img := grid.ImageWithMargin(8, 2)
	buf := new(bytes.Buffer)

	if err := png.Encode(buf, img); err != nil {
		return ec.String(http.StatusNotFound, "")
	}
	return ec.Blob(http.StatusOK, "image/png", buf.Bytes())
}

func (handler *UtilsHandler) GenCaptcha(ec echo.Context) error {
	type (
		Res struct {
			Code         int32  `json:"code"`
			Message      string `json:"message,omitempty"`
			CaptchaToken string `json:"token,omitempty"`
			CaptchaImage string `json:"image,omitempty"`
		}
	)
	r := &Res{
		Code: -1,
	}
	uuidGenerator, _ := uuid.NewV4()
	r.CaptchaToken = uuidGenerator.String()

	client, has := redisStorage["verify"]
	if !has || client == nil {
		r.Message = "内部错误"
		return ec.JSON(http.StatusOK, r)
	}

	token, img := utils.GenB64CaptchaImage()
	if _, err := client.Set(r.CaptchaToken, token, time.Minute*2).Result(); err != nil {
		r.Message = "储存存根失败"
		return ec.JSON(http.StatusOK, r)
	}
	r.Code = 0
	r.CaptchaImage = img
	return ec.JSON(http.StatusOK, r)
}

func (handler *UtilsHandler) GenCaptchaImage(ec echo.Context) error {
	_, img := utils.GenPngCaptchaImage()
	return ec.Blob(http.StatusOK, "image/png", img)
}
