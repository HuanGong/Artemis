package recommendation

import (
	"artemis/uolo_lens/model"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	RecommendRedisPrefix     = "re_"
	CommonRecommendationsKey = "re_commom_recommendations_user"
)

type (
	RecommendImpl struct {
		DataProvider         DataProvider
		CommonRecommendation *model.Recommendations
	}
)

func NewRecommendImpl(provider DataProvider) *RecommendImpl {
	impl := &RecommendImpl{
		DataProvider: provider,
	}
	return impl
}

func (impl *RecommendImpl) ReloadCommonRecommendations() error {
	newRecommendations := &model.Recommendations{
		Articles: &model.RecommendArticles{
			Details: make([]*model.RecommendArticles_Details, 0),
		},
	}
	var articles []*model.Article = make([]*model.Article, 0)
	DbEngine := impl.DataProvider.DBOrmEngine()
	if DbEngine == nil {
		logrus.Errorln("OrmEngine Broken, CommonRecommendations Load Failed")
	}
	err := DbEngine.Select("*").OrderBy("id").Desc("id").Limit(5).Find(&articles)
	if err != nil {
		logrus.Errorln("Query Mysql DB Err:", err.Error())
		return errors.New("mysql db query error")
	}
	for _, a := range articles {
		newRecommendations.Articles.Details = append(newRecommendations.Articles.Details, &model.RecommendArticles_Details{
			Uid: a.Uuid,
		})
	}

	// store to redis
	client := impl.DataProvider.RedisClient("recommend")
	if client == nil {
		return errors.New("Null Redis client for recommend client")
	}
	_, err = client.Set(CommonRecommendationsKey, newRecommendations.String(), time.Hour*24).Result()
	if err != nil {
		logrus.Errorln("Store CommonRecommendations To redis Failed", err.Error())
		return errors.New("Store Commendations to redis failed, error:" + err.Error())
	}
	impl.CommonRecommendation = newRecommendations
	return nil
}

func (impl *RecommendImpl) GetLensRecommendation(context *RecommendContext) (*model.Recommendations, error) {

	recommends := &model.Recommendations{}

	client := impl.DataProvider.RedisClient("recommend")
	if client == nil {
		return nil, errors.New("redis client broken")
	}

	key := RecommendRedisPrefix + context.UserId

	content, err := client.Get(key).Result()
	if /*err == redis.Nil || */ err != nil {
		err = nil
		if impl.CommonRecommendation == nil {
			err = impl.ReloadCommonRecommendations()
		}
		recommends = impl.CommonRecommendation
	} else {
		err = proto.Unmarshal([]byte(content), recommends)
	}

	return recommends, err
}
