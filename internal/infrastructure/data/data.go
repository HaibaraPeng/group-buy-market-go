package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"group-buy-market-go/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDB,
	NewData,
)

// Data .
type Data struct {
	db *gorm.DB
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
func NewData(db *gorm.DB) (*Data, error) {
	d := &Data{
		db: db,
	}
	return d, nil
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
