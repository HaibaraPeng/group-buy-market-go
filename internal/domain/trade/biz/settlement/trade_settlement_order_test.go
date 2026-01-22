package settlement

import (
	"context"
	"group-buy-market-go/internal/infrastructure/adapter/port"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/event/publish"
	"group-buy-market-go/internal/infrastructure/gateway"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/domain/trade/biz/settlement/filter"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"group-buy-market-go/internal/infrastructure/dao"
	"group-buy-market-go/internal/infrastructure/dcc"
)

// setupTestDB 设置测试数据库
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3406)/group_buy_market?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	return db
}

// createTestRepository 创建测试仓库
func createTestRepository(data *data.Data) *repository.TradeRepository {
	groupBuyOrderDAO := dao.NewMySQLGroupBuyOrderDAO(data)
	groupBuyOrderListDAO := dao.NewMySQLGroupBuyOrderListDAO(data)
	groupBuyActivityDAO := dao.NewMySQLGroupBuyActivityDAO(data)
	notifyTaskDAO := dao.NewMySQLNotifyTaskDAO(data)
	// 创建DCC服务实例
	dccService := dcc.NewDCC(data)
	// 创建空的配置对象
	var config *conf.Data
	config.Rabbitmq.Producer.TopicTeamSuccess.RoutingKey = "topic.team_success"

	return repository.NewTradeRepository(
		data,
		groupBuyOrderDAO,
		groupBuyOrderListDAO,
		groupBuyActivityDAO,
		notifyTaskDAO,
		dccService,
		config,
	)
}

func TestTradeSettlementOrderService_SettlementMarketPayOrder_Integration(t *testing.T) {
	// 设置测试数据库
	db := setupTestDB()
	if db == nil {
		t.Skip("Could not connect to test database, skipping integration test")
	}

	// 创建一个模拟的Redis客户端用于测试
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // 测试环境地址，实际测试时可能需要根据环境调整
	})

	// 创建RabbitMQ客户端
	rabbitmqClient := &data.RabbitMQClient{}

	// 创建必要的DAO实例
	d := data.NewData(db, rdb, rabbitmqClient)

	// 创建测试仓库
	testRepo := createTestRepository(d)
	if testRepo == nil {
		t.Skip("Could not create test repository, skipping integration test")
	}

	// 创建服务
	logger := log.DefaultLogger

	// 创建所有必要的过滤器实例
	scRuleFilter := filter.NewSCRuleFilter(logger, testRepo)
	outTradeNoRuleFilter := filter.NewOutTradeNoRuleFilter(logger, testRepo)
	settableRuleFilter := filter.NewSettableRuleFilter(logger, testRepo)
	endRuleFilter := filter.NewEndRuleFilter(logger)

	// 创建通知服务实例
	notifyService := gateway.NewGroupBuyNotifyService()

	// 创建事件发布器实例
	publisher := &publish.RabbitMQEventPublisher{}

	// 创建端口实例
	port := port.NewTradePort(notifyService, d, publisher) // 需要传入notifyService、data和publisher

	// 创建真实的过滤器工厂
	filterFactory := filter.NewTradeSettlementRuleFilterFactory(
		scRuleFilter,
		outTradeNoRuleFilter,
		settableRuleFilter,
		endRuleFilter,
	)
	service := NewTradeSettlementOrderService(logger, testRepo, port, filterFactory, notifyService)

	tradePaySuccessEntity := &model.TradePaySuccessEntity{
		Source:       "s01",
		Channel:      "c01",
		UserId:       "xfg03",
		OutTradeNo:   "904941690333",
		OutTradeTime: time.Now(),
	}

	_, err := service.SettlementMarketPayOrder(context.Background(), tradePaySuccessEntity)

	assert.NoError(t, err) // 没有错误，但返回nil

}
