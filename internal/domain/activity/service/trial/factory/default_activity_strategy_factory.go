package factory

import (
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
)

// DefaultActivityStrategyFactory 默认活动策略工厂
// 负责创建和管理活动策略处理器
type DefaultActivityStrategyFactory struct {
	rootNode tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity]
}

// NewDefaultActivityStrategyFactory 创建默认活动策略工厂
func NewDefaultActivityStrategyFactory(rootNode tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity]) *DefaultActivityStrategyFactory {
	return &DefaultActivityStrategyFactory{
		rootNode: rootNode,
	}
}

// StrategyHandler 获取策略处理器
// 返回根节点策略处理器，作为整个策略树的入口点
func (f *DefaultActivityStrategyFactory) StrategyHandler() tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] {
	return f.rootNode
}
