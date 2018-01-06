package database

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"auction/helper"
	"auction/config"
)

var pool *redis.Pool

// GetRedis 获取redis实例
func GetRedis() redis.Conn {
	return pool.Get()
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				helper.Log("redis start error", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	server := config.Get("Redis")
	password := ""

	pool = newPool(server, password)
}
