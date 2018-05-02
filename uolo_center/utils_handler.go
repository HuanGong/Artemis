package uolo_center

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/qpliu/qrencode-go/qrencode"
	"image/png"
	"bytes"
)

type (
	UtilsHandler struct {
	}
)

func NewUtilsHandler() *UtilsHandler {
	return &UtilsHandler{
	}
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
