package recommendation

import (
	"artemis/uolo_lens/model"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type (
	DataProvider interface {
		DBOrmEngine() *xorm.Engine
		RedisClient(name string) *redis.Client
	}

	RecommendContext struct {
		UserId string
	}

	Recommendation interface {
		GetLensRecommendation(context *RecommendContext) (*model.Recommendations, error)
	}
)
