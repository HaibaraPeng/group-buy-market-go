package dao

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"group-buy-market-go/internal/conf"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDB,
	NewMySQLGroupBuyActivityDAO,
	NewMySQLGroupBuyDiscountDAO,
	NewMySQLSkuDAO,
	NewMySQLCrowdTagsDAO,
	NewMySQLCrowdTagsDetailDAO,
	NewMySQLCrowdTagsJobDAO,
	NewMySQLSCSkuActivityDAO,
	NewMySQLGroupBuyOrderDAO,
	NewMySQLGroupBuyOrderListDAO,
)

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}
