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

	// 从redis中获取当前用户聚合推荐的内容
	redisClient, hasClient := RedisClientMap["recommend"]
	if !hasClient {
		logrus.Errorln("Recommendation Redis DB Broken")
		return ec.String(http.StatusBadRequest, "")
	}

	jwtToken := ec.Get("user").(*jwt.Token)
	if jwtToken == nil {
		logrus.Errorln("jwt token error for nil jwt.Token")
		return ec.String(http.StatusBadRequest, "")
	}

	oldClaims := jwtToken.Claims.(jwt.MapClaims)
	userId := (oldClaims["id"]).(string)
	if len(userId) == 0 {
		logrus.Errorln("jwt token error for nil jwt.Token")
		return ec.String(http.StatusBadRequest, "")
	}

	recommends := &model.Recommendations{}

	content, err := redisClient.Get(userId).Result()
	if err == redis.Nil {
		recommends = defaultRecommondation
	} else if err != nil {
		logrus.Errorln("Redis Get Option Failed with error:", err.Error())
		recommends = defaultRecommondation
	} else {
		err := proto.Unmarshal([]byte(content), recommends)
		if err != nil {
			logrus.Errorln("Decode Recommendations from string failed, userid:", userId)
			return ec.String(http.StatusNotFound, "")
		}
	}
	// 由聚合数据取出具体的内容组装成内容列表

	return ec.JSON(http.StatusOK, nil)
}
