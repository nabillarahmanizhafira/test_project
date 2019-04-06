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
	return
}

// // ProductBid will increase value by 1
// func ProductBid(productID string, bidPrice string) error {
// 	redisConn, err := configs.GlobalConfig.GetRedisConn()
// 	if err != nil {
// 		log.Println("Cannot get redis connection from redis. Error : ", err)
// 		return err
// 	}
// 	sortedSetKey := constructRedisKey(productID)

// 	err = redisConn.ZIncrBy(sortedSetKey, 1, bidPrice)
// 	if err != nil {
// 		log.Println("Cannot ZIncrBy. Error : ", err)
// 		return err
// 	}

// 	return nil
// }

// // GetProduct will return product based on ID
// func GetProduct(productID string) (int, error) {
// 	redisConn, err := configs.GlobalConfig.GetRedisConn()
// 	if err != nil {
// 		log.Println("Cannot get redis connection from redis. Error : ", err)
// 		return -1, err
// 	}
// 	sortedSetKey := constructRedisKey(productID)

// 	val, err := redisConn.Zrange(sortedSetKey, 0, 1, false)
// 	if err != nil {
// 		log.Println("Cannot ZIncrBy. Error : ", err)
// 		return -1, err
// 	}
// 	var winner int
// 	if len(val) > 0 {
// 		winner, _ = strconv.Atoi(val[0])
// 	}

// 	return winner, nil
// }

// func constructRedisKey(productID string) string {
// 	return fmt.Sprintf("HCK:BID:%s", productID)
// }
