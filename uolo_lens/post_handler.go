package uolo_lens

import (
	"artemis/uolo_lens/model"
	"artemis/uolo_lens/utils"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"time"
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
			Code    int32  `json:"code"`
			Message string `json:"message"`
			Content string `json:"content"`
		}
	)

	form := &PostContentForm{}

	if err := ec.Bind(form); err != nil {
		return ec.JSON(http.StatusOK, Response{
			Code:    1,
			Message: "bad argument",
		})
	}

	var fullUri string
	if len(form.Path) == 0 {
		fullUri = conf.PostDataDir + "wellcome.md"
	}
	fullUri = conf.PostDataDir + form.Path
	logrus.Debugln("full uri path:", fullUri)
	//TODO: do a lru cache search
	content, err := ioutil.ReadFile(fullUri)
	if err != nil {
		return ec.JSON(http.StatusOK, Response{
			Code:    1,
			Message: "No Such Post Data",
		})
	}

	res := Response{
		Code:    0,
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

func (handler *PostHandler) ArticleNew(ec echo.Context) error {
	type (
		ArticleNewForm struct {
			Tag     string `json:"tag" query:"tag" form:"tag"`
			Author  string `json:"author" query:"author" form:"author"`
			Title   string `json:"title" query:"title" form:"title"`
			Mime    string `json:"mime" query:"mime" form:"mime"`
			Origin  string `json:"origin" query:"origin" form:"origin"`
			Content string `json:"content" query:"content" form:"content"`
		}
		Response struct {
			Code    int32  `json:"code"`
			Message string `json:"message"`
		}
	)

	form := &ArticleNewForm{}
	err := ec.Bind(form)
	if err != nil {
		return ec.JSON(http.StatusOK, Response{
			Code:    -1,
			Message: "表单错误",
		})
	}

	if len(form.Title) > 128 || len(form.Title) == 0 {
		return ec.JSON(http.StatusOK, Response{
			Code:    -2,
			Message: "标题不符合要求",
		})
	}

	//data dir
	dateFolder := time.Now().Format("20060102")
	fileName := strings.Replace(form.Title, " ", "-", -1)
	fullPath := filepath.Join(conf.PostDataDir, dateFolder, fileName, ".", form.Mime)

	article := &model.Article{
		Tag:          "tag",
		Mime:         form.Mime,
		Title:        form.Title,
		Origin:       form.Origin,
		Author:       form.Author,
		Summary:      "NoSummary",
		ResourcePath: fullPath,
		Status:       1,
	}

	switch form.Mime {
	case "md":
		go utils.SaveArticleAsFileSync(fullPath, form.Content)
		//schedule task do storage content to file
	default:
		logrus.Errorln("current not supported")
		ec.JSON(http.StatusOK, Response{
			Code:    -3,
			Message: "内容类型不支持",
		})
	}

	orm.Insert(article)

	return ec.JSON(http.StatusOK, nil)
}
