package db

import (
	"github.com/garyburd/redigo/redis"
	"mini-douyin-rebuild/config"
)

var Pool *redis.Pool

func RedisInit(s *config.ServerConfig) {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial(
				s.RedisConfig.Network,
				s.RedisConfig.Address,
				redis.DialDatabase(s.RedisConfig.Num),
				redis.DialPassword(s.RedisConfig.RdbPassword))
			if err != nil {
				panic(err.(any))
			}
			return dial, nil
		}}
}
