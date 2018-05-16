package app

import (
	"github.com/garyburd/redigo/redis"
	"config"
	"lib"
)

var RedisClient *redis.Pool

func init() {
	RedisClient = &redis.Pool{
		MaxIdle:     config.RedisMaxIdle,
		MaxActive:   config.RedisMaxActive,
		IdleTimeout: config.RedisIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisAddr)
			lib.CheckErr(err)
			c.Do("SELECT", config.RedisName)
			return c, nil
		},
	}
}


func RedisSet(key, value, exKey, exValue string) {
	redisClient := RedisClient.Get()
	redisClient.Do("SET", key, value, exKey, exValue)
}
