package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// RootNode 根节点
// 策略树的起始节点，负责初始化处理流程
type RootNode struct {
	core.AbstractGroupBuyMarketSupport
}

// NewRootNode 创建根节点
func NewRootNode() *RootNode {
	return &RootNode{}
}

// Apply 应用根节点策略
// 根节点主要负责初始化处理流程
func (r *RootNode) Apply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("开始处理营销活动请求，商品ID: %d, 用户ID: %d", requestParameter.ID, dynamicContext.UserID)

	// 根节点不进行具体业务处理，直接返回空结果
	return &model.TrialBalanceEntity{
		Success: true,
		Message: "根节点处理完成",
	}, nil
}

// Get 获取下一个策略处理器
// 根节点之后通常是开关节点，用于判断是否启用营销活动
func (r *RootNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (core.StrategyHandler, error) {
	log.Printf("根节点处理完成，进入开关节点")

	// 返回开关节点作为下一个处理器
	switchNode := NewSwitchRoot()
	return switchNode, nil
}

// 确保 RootNode 实现了 StrategyHandler 接口
var _ core.StrategyHandler = (*RootNode)(nil)
