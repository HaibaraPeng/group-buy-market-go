package dao

import (
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewMySQLGroupBuyActivityDAO,
	NewMySQLGroupBuyDiscountDAO,
	NewMySQLSkuDAO,
	NewMySQLCrowdTagsDAO,
	NewMySQLCrowdTagsDetailDAO,
	NewMySQLCrowdTagsJobDAO,
	NewMySQLSCSkuActivityDAO,
	NewMySQLGroupBuyOrderDAO,
	NewMySQLGroupBuyOrderListDAO,
	NewMySQLNotifyTaskDAO,
)
