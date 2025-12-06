package core

import (
	"group-buy-market-go/internal/domain/activity/model"
)

// StrategyHandler 策略处理器接口
// 定义了策略树中每个节点需要实现的方法
type StrategyHandler interface {
	// Apply 应用策略，对输入参数进行处理并返回结果
	Apply(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error)

	// Get 获取下一个策略处理器
	// 根据当前节点的处理结果决定下一个执行的节点
	Get(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (StrategyHandler, error)
}
