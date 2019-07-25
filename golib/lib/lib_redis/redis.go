package lib_redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-micro/golib/lib/lib_config"
	"log"
)


func InitRedis(config lib_config.ConfRedis) *redis.Client{
	hostAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     hostAddr,
		Password: config.Auth,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatalln("init redis fail: ", err)
	}

	return client
}

//func (p *RedisInstance) Delete(ctx context.Context, key string) error {
//	err := opentracing.GetClientFromContext(ctx, p.redisClient).Del(key).Err()
//	return err
//}
//
//func (p *RedisInstance) HSet(ctx context.Context, key, field, content string) (bool, error) {
//	if "" == key || "" == field || "" == content {
//		return false, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HSet(key, field, content).Result()
//}
//
//func (p *RedisInstance) HMSet(ctx context.Context, key string, content map[string]interface{}) error {
//	if "" == key || nil == content {
//		return unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HMSet(key, content).Err()
//}
//
//func (p *RedisInstance) HGet(ctx context.Context, key, field string) (string, error) {
//	if "" == key || "" == field {
//		return "", unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HGet(key, field).Result()
//}
//
//func (p *RedisInstance) HDel(ctx context.Context, key, field string) (int64, error) {
//	if "" == key || "" == field {
//		return 0, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HDel(key, field).Result()
//}
//
//func (p *RedisInstance) HGetAll(ctx context.Context, key string) (map[string]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HGetAll(key).Result()
//}
//
//func (p *RedisInstance) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
//	if "" == key || "" == field {
//		return 0, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).HIncrBy(key, field, incr).Result()
//}
//
//func (p *RedisInstance) LPush(ctx context.Context, key string, value string) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//	err := opentracing.GetClientFromContext(ctx, p.redisClient).LPush(key, value).Err()
//	return err
//}
//
//func (p *RedisInstance) LGetAll(ctx context.Context, key string) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).LRange(key, 0, -1).Result()
//}
//
//func (p *RedisInstance) Set(ctx context.Context, key string, value interface{}, expr time.Duration) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).Set(key, value, expr).Err()
//}
//
//func (p *RedisInstance) SetNX(ctx context.Context, key string, value interface{}, expr time.Duration) (bool, error) {
//	if "" == key {
//		return false, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).SetNX(key, value, expr).Result()
//}
//
//func (p *RedisInstance) Get(ctx context.Context, key string) (string, error) {
//	if "" == key {
//		return "", unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).Get(key).Result()
//}
//
//func (p *RedisInstance) Del(ctx context.Context, key string) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).Del(key).Err()
//}
//
//func (p *RedisInstance) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
//	if "" == key || 0 == expiration {
//		return false, unicornstd.ErrParamsIsEmpty
//	}
//	return opentracing.GetClientFromContext(ctx, p.redisClient).Expire(key, expiration).Result()
//}
//
//func (p *RedisInstance) ZAdd(ctx context.Context, key string, score float64, content interface{}) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZAdd(key, redis.Z{
//		Score:  score,
//		Member: content,
//	}).Err()
//}
//
//func (p *RedisInstance) ZRange(ctx context.Context, key string, start, end int64) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZRange(key, start, end).Result()
//}
//
//func (p *RedisInstance) ZRangeAll(ctx context.Context, key string) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZRange(key, 0, -1).Result()
//}
//
//func (p *RedisInstance) ZScore(ctx context.Context, key, member string) (float64, error) {
//	if "" == key {
//		return 0, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZScore(key, member).Result()
//}
//
//func (p *RedisInstance) ZRevRangeAll(ctx context.Context, key string) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZRevRange(key, 0, -1).Result()
//}
//
//func (p *RedisInstance) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZRevRange(key, start, stop).Result()
//}
//
//func (p *RedisInstance) ZREM(ctx context.Context, key string, start, stop int64) (int64, error) {
//	if "" == key {
//		return 0, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).ZRemRangeByRank(key, start, stop).Result()
//}
//
//func (p *RedisInstance) SAdd(ctx context.Context, key string, content interface{}) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).SAdd(key, content).Err()
//}
//
//func (p *RedisInstance) Incr(ctx context.Context, key string) (int64, error) {
//	return opentracing.GetClientFromContext(ctx, p.redisClient).Incr(key).Result()
//}
//
//func (p *RedisInstance) SMembers(ctx context.Context, key string) ([]string, error) {
//	if "" == key {
//		return nil, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).SMembers(key).Result()
//}
//
//func (p *RedisInstance) SIsMember(ctx context.Context, key string, content interface{}) (bool, error) {
//	if "" == key {
//		return false, unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).SIsMember(key, content).Result()
//}
//
//func (p *RedisInstance) SRem(ctx context.Context, key string, content interface{}) error {
//	if "" == key {
//		return unicornstd.ErrParamsIsEmpty
//	}
//
//	return opentracing.GetClientFromContext(ctx, p.redisClient).SRem(key, content).Err()
//}
