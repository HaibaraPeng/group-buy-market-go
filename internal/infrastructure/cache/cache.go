package cache

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"group-buy-market-go/internal/conf"
)

// ProviderSet is cache providers.
var ProviderSet = wire.NewSet(
	NewRedisClient,
)

// NewRedisClient 创建redis客户端
func NewRedisClient(conf *conf.Data, logger log.Logger) (*redis.Client, func(), error) {
	// 创建redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})

	rdb.AddHook(redisotel.TracingHook{})
	return rdb, func() {
		log.Info("message", "closing the redis resources")
		if err := rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
