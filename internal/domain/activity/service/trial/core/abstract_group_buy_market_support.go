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

// Get 获取待执行策略 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) Get(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity], error) {
	// 子类需要实现此方法
	var handler tree.StrategyHandler[*model.MarketProductEntity, *DynamicContext, *model.TrialBalanceEntity]
	return handler, nil
}
