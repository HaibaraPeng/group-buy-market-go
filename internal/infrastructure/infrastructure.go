package infrastructure

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/dcc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	dao.ProviderSet,
	dcc.ProviderSet,
	repository.ProviderSet,
)

type Data struct {
	Rdb *redis.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

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

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
		return nil, nil, err
	}

	d := &Data{
		Rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.Rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
