package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"group-buy-market-go/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDB,
	NewData,
	NewRedisClient,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

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

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.NewHelper(logger).Errorf("failed to connect to redis: %v", err)
		return nil, nil, err
	}

	return rdb, func() {
		log.Info("message", "closing the redis resources")
		if err := rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client) *Data {
	d := &Data{
		db:  db,
		rdb: rdb,
	}
	return d
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func (d *Data) Rdb(ctx context.Context) *redis.Client {
	return d.rdb
}
