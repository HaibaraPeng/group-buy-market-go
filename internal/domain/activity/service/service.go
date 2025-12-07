package service

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/factory"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
)

// IIndexGroupBuyMarketService 拼团营销服务
// 提供对外的营销试算服务接口
type IIndexGroupBuyMarketService struct {
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewIIndexGroupBuyMarketService 创建拼团营销服务实例
func NewIIndexGroupBuyMarketService() *IIndexGroupBuyMarketService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	rootNode := node.NewRootNode()
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &IIndexGroupBuyMarketService{
		strategyFactory: strategyFactory,
	}
}

// IndexMarketTrial 首页营销试算
// 对应Java中的indexMarketTrial方法
func (s *IIndexGroupBuyMarketService) IndexMarketTrial(marketProductEntity *model.MarketProductEntity) (*model.TrialBalanceEntity, error) {
	// 获取策略处理器
	strategyHandler := s.strategyFactory.StrategyHandler()

	// 创建动态上下文
	dynamicContext := &core.DynamicContext{}

	// 应用策略处理器
	trialBalanceEntity, err := strategyHandler.Apply(marketProductEntity, dynamicContext)
	if err != nil {
		return nil, err
	}

	return trialBalanceEntity, nil
}
