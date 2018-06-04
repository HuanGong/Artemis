package recommendation

import (
	"artemis/uolo_lens/model"
	"artemis/uolo_lens/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	RecommendRedisPrefix        = "re_"
	CommonRecommendationsUser   = "commom_recommendations_user"
	CommonRecommendationsKey    = "re_commom_recommendations_user"
	RecommendationRedisLifeTime = time.Hour * 1
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
		return errors.New("OrmEngine Broken")
	}
	err := DbEngine.Select("*").OrderBy("id").Desc("id").Limit(5).Find(&articles)
	if err != nil {
		return errors.Wrapf(err, "mysql db query error")
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

	raw, err := utils.PBEncode(newRecommendations)
	if err != nil {
		return errors.Wrapf(err, "Encode Recommendations to url encoded string failed")
	}

	_, err = client.Set(CommonRecommendationsKey, raw, time.Hour*24).Result()
	if err != nil {
		logrus.Errorln("Store CommonRecommendations To redis Failed", err.Error())
		return errors.New("Store Commendations to redis failed, error:" + err.Error())
	}
	impl.CommonRecommendation = newRecommendations
	return nil
}

func (impl *RecommendImpl) BuildRecommendationsForUser(usr string) error {
	newRecommendations := &model.Recommendations{
		Articles: &model.RecommendArticles{
			Details: make([]*model.RecommendArticles_Details, 0),
		},
	}
	articles := make([]*model.Article, 0)
	DbEngine := impl.DataProvider.DBOrmEngine()
	if DbEngine == nil {
		logrus.Errorln("OrmEngine Broken, CommonRecommendations Load Failed")
		return errors.New("OrmEngine Broken")
	}
	err := DbEngine.Alias("tb").Select("*").Desc("id").Limit(5).Find(&articles)
	if err != nil {
		logrus.Errorln("mysql db query error:", err.Error())
		return errors.Wrapf(err, "mysql db query error")
	}
	for _, a := range articles {
		logrus.Debugln("Got Recommend article, Id:", a.Id)
		newRecommendations.Articles.Details = append(newRecommendations.Articles.Details, &model.RecommendArticles_Details{
			Uid: a.Uuid,
		})
	}

	// store to redis
	client := impl.DataProvider.RedisClient("recommend")
	if client == nil {
		logrus.Errorln("Null Redis client for recommend client")
		return errors.New("Null Redis client for recommend client")
	}

	raw, err := utils.PBEncode(newRecommendations)
	if err != nil {
		logrus.Errorln("Encode Recommendation to url encoded string failed")
		return errors.Wrapf(err, "Encode Recommendations to url encoded string failed")
	}

	key := RecommendRedisPrefix + usr
	_, err = client.Set(key, raw, RecommendationRedisLifeTime).Result()
	if err != nil {
		logrus.Errorln("Store CommonRecommendations To redis Failed", err.Error())
		return errors.New("Store Commendations to redis failed, error:" + err.Error())
	}
	logrus.Debugln("Generate Recommendation for user:", key)
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
			go impl.BuildRecommendationsForUser(context.UserId)
		}
		recommends = impl.CommonRecommendation
	} else {
		logrus.Debugln("Got Content From User:", key)
		err = utils.PBDecode(content, recommends)
	}

	return recommends, err
}
