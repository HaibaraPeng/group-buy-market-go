package settlement

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/dao"
)

// setupTestDB 设置测试数据库
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/group_buy_market?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	return db
}

// createTestRepository 创建测试仓库
func createTestRepository(db *gorm.DB) *repository.TradeRepository {
	if db == nil {
		return nil
	}

	// 创建必要的DAO实例
	groupBuyOrderDAO := dao.NewMySQLGroupBuyOrderDAO(db)
	groupBuyOrderListDAO := dao.NewMySQLGroupBuyOrderListDAO(db)
	groupBuyActivityDAO := dao.NewMySQLGroupBuyActivityDAO(db)
	notifyTaskDAO := dao.NewMySQLNotifyTaskDAO(db)

	return repository.NewTradeRepository(
		groupBuyOrderDAO,
		groupBuyOrderListDAO,
		groupBuyActivityDAO,
		notifyTaskDAO,
	)
}

func TestTradeSettlementOrderService_SettlementMarketPayOrder_Integration(t *testing.T) {
	// 设置测试数据库
	db := setupTestDB()
	if db == nil {
		t.Skip("Could not connect to test database, skipping integration test")
	}

	// 创建测试仓库
	testRepo := createTestRepository(db)
	if testRepo == nil {
		t.Skip("Could not create test repository, skipping integration test")
	}

	// 创建服务
	service := NewTradeSettlementOrderService(testRepo)

	tradePaySuccessEntity := &model.TradePaySuccessEntity{
		Source:     "s01",
		Channel:    "c01",
		UserId:     "xfg04",
		OutTradeNo: "451517755304",
	}

	_, err := service.SettlementMarketPayOrder(context.Background(), tradePaySuccessEntity)

	assert.NoError(t, err) // 没有错误，但返回nil

}
