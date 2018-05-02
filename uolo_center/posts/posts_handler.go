package posts

import (
	"github.com/labstack/echo"
	"net/http"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/sirupsen/logrus"
)

type (
	PostHandler struct {

	}

	PostRes struct {
		Success bool `json:"success"`
		Message string `json:"message"`
	}
)

func (handler *PostHandler) PostNew(ec echo.Context) error {
	return ec.JSON(http.StatusOK, PostRes{
		Success: true,
		Message: "Success",
	})
}

func (handler *PostHandler) GetPost(ec echo.Context) error {
	input := `# this is a title

**strong**
	
- hello nice
- hello nice

<img align="right" src="//api.echoface.cn/utils/qrcode2?text=gonghuan"/>

这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈

this is content detail

this is content detail

this is content detail

this is content detail

this is content detail

this is content detail

这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈

这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈

这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈；这是测试图片右边排列的文字， 可以吗这样， 如果可以的话， 那就太哈了哈哈哈哈哈哈哈

![qrcode2](//api.echoface.cn/utils/qrcode2?text="gonghuan")

this is content detail`
	unsafe := blackfriday.Run([]byte(input))
	logrus.Debugln(string(unsafe))
	return ec.HTMLBlob(200, unsafe)
}
