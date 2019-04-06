package cache

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

//PkgRedis implementation of RedisInterface
type PkgRedis struct {
	rpool *redis.Pool
}

// NewRedis return new object redis which implement cache
func NewRedis(host string, poolSize int) (cache Cache) {
	cache = InitRedis(host, poolSize)
	return
}

// InitRedis will return implementation of redigo which implements Cache Interface
func InitRedis(host string, poolSize int) (pkgRedis *PkgRedis) {
	pkgRedis = &PkgRedis{}
	pkgRedis.rpool = &redis.Pool{
		MaxIdle:         (poolSize * 3 / 4),
		MaxActive:       poolSize,
		IdleTimeout:     3 * time.Second,
		MaxConnLifetime: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host)
			if err != nil {
				log.Println("Error when dialling connection:", err)
				return nil, err
			}
			return conn, err
		},
	}
	return
}

// Get will return value based on key
func (p *PkgRedis) Get(key string) (value string, err error) {
	conn, err := p.rpool.Dial()
	if err != nil {
		log.Println("Failed to dial connection from connection pool. ", err)
		return "", err
	}
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", key))
	if err != nil {
		log.Println("Failed to send GET command. ", err)
		return "", err
	}
	return val, err
}

// Ping will do healthcheck
func (p *PkgRedis) Ping() error {
	conn, err := p.rpool.Dial()
	if err != nil {
		log.Println("Failed to dial connection from connection pool. ", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("PING")
	if err != nil {
		log.Println("Failed to send PING command. ", err)
		return err
	}

	return err
}

// Set will set value to key with ttl
func (p *PkgRedis) Set(key, value string, ttl ...int) (err error) {
	conn, err := p.rpool.Dial()
	if err != nil {
		log.Println("Failed to dial connection from connection pool. ", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value, ttl)
	if err != nil {
		log.Println("Failed to send Set command. ", err)
		return err
	}

	return
}
