package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/service/trial"
	"log"
)

// EndNode 结束节点
// 策略树的终止节点，负责收尾工作和返回最终结果
type EndNode struct {
	trial.AbstractGroupBuyMarketSupport
}

// NewEndNode 创建结束节点
func NewEndNode() *EndNode {
	return &EndNode{}
}

// Apply 应用结束节点策略
// 结束节点汇总前面节点的处理结果，并做最后的封装
func (e *EndNode) Apply(requestParameter *model.MarketProductEntity, dynamicContext *trial.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("营销活动处理流程结束，商品ID: %d", requestParameter.ID)

	// 这里可以做一些收尾工作，比如记录日志、更新统计数据等

	// 返回最终的结果
	result := &model.TrialBalanceEntity{
		Success: true,
		Message: "营销活动处理完成",
		Code:    "SUCCESS",
	}

	return result, nil
}

// Get 获取下一个策略处理器（结束节点通常返回nil）
// 结束节点之后没有其他处理器
func (e *EndNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *trial.DynamicContext) (trial.StrategyHandler, error) {
	log.Printf("处理流程完全结束")

	// 结束节点没有下一个处理器
	return nil, nil
}

// 确保 EndNode 实现了 StrategyHandler 接口
var _ trial.StrategyHandler = (*EndNode)(nil)
