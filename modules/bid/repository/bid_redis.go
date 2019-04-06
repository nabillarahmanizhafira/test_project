package repository

import (
	"github.com/nabillarahmanizhafira/test_project/common/configs"
	"github.com/nabillarahmanizhafira/test_project/connections/cache"
	"github.com/nabillarahmanizhafira/test_project/models"
	"github.com/nabillarahmanizhafira/test_project/modules/bid"
)

type bidRedis struct {
	redisConn cache.Cache
}

// NewBidRedis returns an instance of bidRedis
func NewBidRedis() bid.Repository {
	return &bidRedis{
		redisConn: configs.GlobalConfig.GetRedisConn(),
	}
}

func (bc *bidRedis) GetByID(ID string) (res models.Product, err error) {
	val, err := bc.redisConn.Get(ID)
	res.ID = ID
	res.Value = val
	return
}

func (bc *bidRedis) SetProduct(ID, value string) (err error) {
	err = bc.redisConn.Set(ID, value, 3600)
	return
}
