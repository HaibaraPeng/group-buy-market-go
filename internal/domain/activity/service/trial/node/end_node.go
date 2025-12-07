package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// EndNode 结束节点
// 策略树的终止节点，负责收尾工作和返回最终结果
type EndNode struct {
	core.AbstractGroupBuyMarketSupport
}

// NewEndNode 创建结束节点
func NewEndNode() *EndNode {
	return &EndNode{}
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (e *EndNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-EndNode  requestParameter:%+v", requestParameter)

	groupBuyActivityDiscountVO := dynamicContext.GetGroupBuyActivityDiscountVO()
	skuVO := dynamicContext.GetSkuVO()

	// 返回空结果
	result := &model.TrialBalanceEntity{
		Success: true,
		Message: "处理完成",
		Code:    "SUCCESS",
	}

	// 如果有有效的VO数据，则填充更多字段
	if skuVO != nil && groupBuyActivityDiscountVO != nil {
		result.TotalAmount = skuVO.OriginalPrice
		result.FinalAmount = skuVO.OriginalPrice
		result.DiscountAmount = 0.0
	}

	return result, nil
}

// 确保 EndNode 实现了 StrategyHandler 接口
var _ core.StrategyHandler = (*EndNode)(nil)
