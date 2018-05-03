package uolo_lens

import (
	"github.com/labstack/echo"
	"gopkg.in/russross/blackfriday.v2"
	"net/http"
	"io/ioutil"
	"github.com/sirupsen/logrus"
)

type (
	PostHandler struct {
	}
)

func (handler *PostHandler) NewPost(ec echo.Context) error {
	return ec.String(http.StatusOK, "")
}

func (handler *PostHandler) PostContentDetail(ec echo.Context) error {
	type (
		PostContentForm struct {
			Path string `json:"path" query:"path" form:"path"`
			Type string `json:"type" query:"type" form:"type"`
		}
		Response struct {
			Code int32		`json:"code"`
			Message string	`json:"message"`
			Content string	`json:"content"`
		}
	)

	form := &PostContentForm{}

	if err := ec.Bind(form); err != nil {
		return ec.JSON(http.StatusOK, Response{
			Code: 1,
			Message: "bad argument",
		})
	}

	var fullUri string
	if len(form.Path) == 0 {
		fullUri = conf.PostDataDir + "wellcome.md"
	}
	fullUri = conf.PostDataDir +form.Path
	logrus.Debugln("full uri path:", fullUri)
	//TODO: do a lru cache search
	content, err := ioutil.ReadFile(fullUri)
	if err != nil {
		return ec.JSON(http.StatusOK, Response{
			Code: 1,
			Message: "No Such Post Data",
		})
	}

	res := Response{
		Code: 0,
		Message: "success",
		Content: "",
	}
	if form.Type == "md" {
		unsafe := blackfriday.Run(content)
		res.Content = string(unsafe)
	} else {
		res.Content = ""
		res.Code = 2
		res.Message = "Current Only Support Markdown Type"
	}

	return ec.JSON(http.StatusOK, res)
}

func (handler *PostHandler) PostNew(ec echo.Context) error {
	type (
		PostContentForm struct {
			Path string `json:"path" query:"path" form:"path"`
			Type string `json:"type" query:"type" form:"type"`
		}
		Response struct {
			Code int32		`json:"code"`
			Message string	`json:"message"`
			Content string	`json:"content"`
		}
	)

	return ec.JSON(http.StatusOK, nil)
}
