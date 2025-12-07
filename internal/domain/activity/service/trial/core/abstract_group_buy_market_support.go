package core

import (
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
)

// AbstractGroupBuyMarketSupport 是拼团营销支撑的抽象基类
// 它提供了所有策略节点的公共功能和基础实现
type AbstractGroupBuyMarketSupport struct {
	tree.AbstractMultiThreadStrategyRouter[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity]
}

// SetDefaultStrategyHandler 设置默认策略处理器
func (s *AbstractGroupBuyMarketSupport) SetDefaultStrategyHandler(handler tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity]) {
	s.AbstractMultiThreadStrategyRouter.SetDefaultStrategyHandler(handler)
}

// GetDefaultStrategyHandler 获取默认策略处理器
func (s *AbstractGroupBuyMarketSupport) GetDefaultStrategyHandler() tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity] {
	return s.AbstractMultiThreadStrategyRouter.GetDefaultStrategyHandler()
}

// Router 路由策略
func (s *AbstractGroupBuyMarketSupport) Router(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	return s.AbstractMultiThreadStrategyRouter.Router(requestParameter, dynamicContext)
}

// Apply 应用策略
func (s *AbstractGroupBuyMarketSupport) Apply(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	return s.AbstractMultiThreadStrategyRouter.Apply(requestParameter, dynamicContext)
}

// MultiThread 异步加载数据 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) MultiThread(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) error {
	// 子类需要实现此方法
	return nil
}

// DoApply 业务流程受理 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) DoApply(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	// 子类需要实现此方法
	return &model.TrialBalanceEntity{}, nil
}

// Get 获取待执行策略 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) Get(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity], error) {
	// 子类需要实现此方法
	var handler tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity]
	return handler, nil
}
