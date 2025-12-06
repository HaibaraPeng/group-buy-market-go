package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/service/trial/types"
	"log"
)

// MarketNode 营销优惠节点
// 负责计算各种营销优惠
type MarketNode struct {
	core.AbstractGroupBuyMarketSupport
}

// NewMarketNode 创建营销节点
func NewMarketNode() *MarketNode {
	return &MarketNode{}
}

// Apply 应用营销节点策略
// 计算商品的营销优惠，包括折扣、满减等
func (m *MarketNode) Apply(requestParameter *model.MarketProductEntity, dynamicContext *types.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("计算商品营销优惠，商品ID: %d, 原价: %.2f", requestParameter.ID, requestParameter.Price)

	// 模拟营销优惠计算过程
	totalAmount := requestParameter.Price
	discountAmount := 0.0

	// 根据用户等级计算折扣
	switch dynamicContext.UserLevel {
	case 1:
		discountAmount = totalAmount * 0.1 // 普通用户9折
	case 2:
		discountAmount = totalAmount * 0.2 // 黄金用户8折
	case 3:
		discountAmount = totalAmount * 0.3 // 钻石用户7折
	default:
		discountAmount = 0 // 默认无折扣
	}

	finalAmount := totalAmount - discountAmount

	result := &model.TrialBalanceEntity{
		TotalAmount:    totalAmount,
		DiscountAmount: discountAmount,
		FinalAmount:    finalAmount,
		Success:        true,
		Message:        "营销优惠计算完成",
	}

	return result, nil
}

// Get 获取下一个策略处理器
// 营销节点处理完成后进入结束节点
func (m *MarketNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *types.DynamicContext) (types.StrategyHandler, error) {
	log.Printf("营销节点处理完成，进入结束节点")

	// 返回结束节点作为下一个处理器
	endNode := NewEndNode()
	return endNode, nil
}

// 确保 MarketNode 实现了 StrategyHandler 接口
var _ types.StrategyHandler = (*MarketNode)(nil)
