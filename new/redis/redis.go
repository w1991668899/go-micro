package redis

import (
	"github.com/go-redis/redis"
	"go-micro/golib/lib/lib_redis"
	"go-micro/server-one/common"
)

var redisCli *redis.Client

func InitRedis()  {
	redisCli = lib_redis.InitRedis(common.GloConf.Redis)
}

