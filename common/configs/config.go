package configs

import (
	"os"
	"strconv"

	"github.com/nabillarahmanizhafira/test_project/common/log"
	"github.com/nabillarahmanizhafira/test_project/connections/cache"
	"github.com/subosito/gotenv"
)

// InitConfigVar will load var from .env and parse it to struct Configurations
func InitConfigVar() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Load env error:", err)
	}

	// process to load the env
	ReadConfigFile()
}

// ReadConfigFile will parse env to struct from env vars
func ReadConfigFile() {
	VarsConfig = Configuration{
		Server: ServerConfiguration{
			Port: os.Getenv("SERVER_PORT"),
		},
		Redis: RedisConfiguration{
			Host:     os.Getenv("REDIS_HOST"),
			PoolSize: os.Getenv("REDIS_POOL_SIZE"),
		},
	}
	return
}

// InitRedisConn will init redis and pass it into GlobalConfig
func InitRedisConn() {
	if VarsConfig.Redis.Host == "" {
		log.Fatal("Failed init redis, empty host from env")
	}

	// init object redis
	poolSize, err := strconv.Atoi(VarsConfig.Redis.PoolSize)
	if err != nil {
		log.Error(err, "failed parse pool size")
		poolSize = 5
	}
	GlobalConfig.cache = cache.NewRedis(VarsConfig.Redis.Host, poolSize)
}

// GetRedisConn will return redisConn
func (pkg *AppConfig) GetRedisConn() (redisConn cache.Cache) {
	if pkg.cache == nil {
		log.Fatal("redis conn hasn't been initialized")
	}

	redisConn = pkg.cache
	return
}
