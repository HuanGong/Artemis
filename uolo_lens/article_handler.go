package uolo_lens

import (
	"artemis/uolo_lens/model"
	"artemis/uolo_lens/utils"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	ArticlePublished      = 0
	ArticleWaitForPublish = 1
)

type (
	PostHandler struct {
	}
)

func (handler *PostHandler) ArticleDetail(ec echo.Context) error {
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
			Code:    -1,
			Message: "bad argument",
		})
	}

	var fullUri string
	if len(form.Path) == 0 { //404.html
		fullUri = conf.PostDataDir + "404-not-found.md"
	}
	fullUri = conf.PostDataDir + form.Path
	logrus.Debugln("full uri path:", fullUri)

	//TODO: do a lru cache search
	content, err := ioutil.ReadFile(fullUri)
	if os.IsNotExist(err) {
		fullUri = conf.PostDataDir + "404-not-found.md"
		content, err = ioutil.ReadFile(fullUri)
		if err != nil {
			return ec.JSON(http.StatusOK, Response{
				Code:    -1,
				Message: "What Happend!!!",
			})
		}
	}

	res := Response{
		Code:    0,
		Message: "success",
	}

	switch form.Type {
	case "md":
		res.Content = string(content)
	case "html":
		unsafe := blackfriday.Run(content)
		res.Content = string(unsafe)
	default:
		res.Code = -2
		res.Message = fmt.Sprintf("Current Only Support Your Requested Type %s", form.Type)
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
			Summary string `json:"summary" query:"summary" form:"summary"`
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
	uuidName, _ := uuid.NewV4()
	fileName := base64.RawURLEncoding.EncodeToString(uuidName[:])
	//fileName := strings.Replace(form.Title, " ", "-", -1) + "." + form.Mime
	fullPath := filepath.Join(conf.PostDataDir, dateFolder, fileName)

	article := &model.Article{
		Tag:     form.Tag,
		Mime:    form.Mime,
		Title:   form.Title,
		Origin:  form.Origin,
		Author:  form.Author,
		Summary: form.Summary,
		Rpath:   fullPath,
		Status:  ArticleWaitForPublish,
		Count:   0,
	}

	switch form.Mime {
	case "md":
		go utils.SaveArticleAsFileSync(fullPath, form.Content)
		//schedule task do storage content to file
	default:
		logrus.Errorln("current not supported")
		return ec.JSON(http.StatusOK, Response{
			Code:    -3,
			Message: "内容类型不支持",
		})
	}

	_, err = orm.Insert(article)
	if err != nil {
		logrus.Errorln("Insert to database err:", err.Error())
		return ec.JSON(http.StatusOK, Response{
			Code:    -4,
			Message: "插入数据库失败",
		})
	}

	return ec.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "OK",
	})
}

func (handler *PostHandler) DialysisConent(ec echo.Context) error {
	type (
		In struct {
			Url string `json:"url" query:"url" form:"url"`
		}
		Response struct {
			Code     int32  `json:"code" query:"code" form:"code"`
			Message  string `json:"message" query:"message" form:"message"`
			Author   string `json:"author" query:"author" form:"author"`
			Title    string `json:"title" query:"title" form:"title"`
			Mime     string `json:"mime" query:"mime" form:"mime"`
			Desc     string `json:"desc" query:"desc" form:"desc"`
			Origin   string `json:"origin" query:"origin" form:"origin"`
			Date     string `json:"date" query:"date" form:"date"`
			Keywords string `json:"keywords" query:"keywords" form:"keywords"`
			Content  string `json:"content" query:"content" form:"content"`
		}
	)

	res := Response{
		Code: -1,
	}

	in := &In{}
	if err := ec.Bind(in); err != nil {
		res.Message = "参数错误"
		return ec.JSON(http.StatusOK, res)
	}

	result, err := utils.ExtractArticleFromUrl(in.Url)
	if err != nil {
		res.Message = "解析错误"
		return ec.JSON(http.StatusOK, res)
	}

	res.Code = 0
	res.Message = "Ok"
	res.Mime = "md"
	res.Desc = result["desc"]
	res.Date = result["date"]
	res.Title = result["title"]
	res.Origin = result["link"]
	res.Author = result["author"]
	res.Keywords = result["keywords"]
	res.Content = result["content"]
	return ec.JSON(http.StatusOK, res)
}
