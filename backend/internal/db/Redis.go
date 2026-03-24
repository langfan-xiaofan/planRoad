package db

import (
	"backend/internal/config"
	"context"

	"github.com/redis/go-redis/v9"
)

var Redis = config.Conf.Redis
var RedisDatabase *redis.Client
var ctx = context.Background()

func RedisInit() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Redis.Url,
		Password: Redis.Password,
		DB:       Redis.Db,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	RedisDatabase = rdb
	return nil
}
