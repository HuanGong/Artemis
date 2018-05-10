package uolo_lens

import (
	"artemis/uolo_lens/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/gogo/protobuf/proto"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	LensHandler struct {
	}
)

var defaultRecommondation = &model.Recommendations{
	Articles: &model.RecommendArticles{
		Details: make([]*model.RecommendArticles_Details, 0),
	},
}

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
	// 从redis中获取当前用户聚合推荐的内容
	redisClient, hasClient := RedisClientMap["recommend"]
	if !hasClient {
		logrus.Errorln("Recommendation Redis DB Broken")
		return failedResponse(-1, "Query KV Error")
	}
	userId := "TestUserId"
	if "" == ec.QueryParam("test") {
		jwtToken := ec.Get("user")
		if jwtToken == nil {
			logrus.Errorln("jwt token error for nil jwt.Token")
			return failedResponse(-2, "Jwt Token Error")
		}

		oldClaims := jwtToken.(*jwt.Token).Claims.(jwt.MapClaims)
		userId = (oldClaims["id"]).(string)
		if len(userId) == 0 {
			logrus.Errorln("jwt token error for nil jwt.Token")
			return failedResponse(-3, "Jwt Claims Error")
		}
	}

	recommends := &model.Recommendations{}

	content, err := redisClient.Get(userId).Result()
	if err == redis.Nil || err != nil {
		recommends = defaultRecommondation
	} else {
		err := proto.Unmarshal([]byte(content), recommends)
		if err != nil {
			recommends = defaultRecommondation
			logrus.Errorln("Decode Recommendations from string failed, userid:", userId)
		}
	}
	logrus.Debugln("Query From DB")

	var articles []*model.Article = make([]*model.Article, 0)
	err = Orm.OrderBy("RAND()").Limit(5).Find(&articles)
	if err != nil {
		logrus.Debugln("Query DB Err:", err.Error())
		return failedResponse(-4, "Qeury DB Failed")
	}
	for _, a := range articles {
		logrus.Debugln("Got Article:", a)
	}

	res.Code = 0
	res.Message = "ok"
	res.LensList = &model.LensList{
		Articles: articles,
	}
	logrus.Debugln("response: ", res)
	return ec.JSON(http.StatusOK, res)
}
