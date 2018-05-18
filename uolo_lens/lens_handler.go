package uolo_lens

import (
	"artemis/uolo_lens/model"
	"artemis/uolo_lens/recommendation"
	"artemis/uolo_lens/utils"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	LensHandler struct {
	}
)

func (handler *LensHandler) LensList(ec echo.Context) error {
	type Response struct {
		Code     int32           `json:"code"`
		Message  string          `json:"message"`
		LensList *model.LensList `json:"lensList"`
	}

	res := &Response{
		Code: 0,
	}

	failedResponse := func(code int32, message string) error {
		res.Code = code
		res.Message = message
		return ec.JSON(http.StatusOK, res)
	}

	userId, login := utils.IsUserLogin(ec)
	if !login {
		userId = recommendation.CommonRecommendationsUser
	}
	logrus.Debugf("User:[%s] ", userId)
	recommends, err := Recommender.GetLensRecommendation(&recommendation.RecommendContext{
		UserId: userId,
	})
	if err != nil {
		logrus.Errorln("get recommend failed with err:", err.Error())
		return failedResponse(-1, "Not Found Content")
	}

	//extract articles
	DbqueryIds := make([]string, 0)
	for _, detail := range recommends.Articles.Details {
		DbqueryIds = append(DbqueryIds, detail.Uid)
	}
	articles := make([]*model.Article, 0)
	err = Orm.In("uuid", DbqueryIds).Find(&articles)
	if err != nil {
		return failedResponse(-2, "Query In DB Failed")
	}

	res.Code = 0
	res.Message = "ok"
	res.LensList = &model.LensList{
		Articles: articles,
	}
	return ec.JSON(http.StatusOK, res)
}
