package host

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisCfg redis.Options

var RDB *redis.Client

func initRedis() (err error) {
	RDB = redis.NewClient((*redis.Options)(Cfg.Redis))
	err = RDB.Ping(context.Background()).Err()
	return
}
