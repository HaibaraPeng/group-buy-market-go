package node

import (
	"group-buy-market-go/common/design/tree"
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
	endNode := &EndNode{}

	// 设置自定义方法实现
	endNode.SetDoApplyFunc(endNode.doApply)
	endNode.SetMultiThreadFunc(endNode.multiThread)

	return endNode
}

// multiThread 异步加载数据 - 结束节点不需要异步加载
func (e *EndNode) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 结束节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (e *EndNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-EndNode userId:%s requestParameter:%+v", requestParameter.UserId, requestParameter)

	groupBuyActivityDiscountVO := dynamicContext.GetGroupBuyActivityDiscountVO()
	skuVO := dynamicContext.GetSkuVO()

	// 返回空结果
	result := &model.TrialBalanceEntity{
		GoodsId:        skuVO.GoodsId,
		GoodsName:      skuVO.GoodsName,
		OriginalPrice:  skuVO.OriginalPrice,
		DeductionPrice: 0.0,
		TargetCount:    groupBuyActivityDiscountVO.Target,
		StartTime:      groupBuyActivityDiscountVO.StartTime.Unix(),
		EndTime:        groupBuyActivityDiscountVO.EndTime.Unix(),
		IsVisible:      false,
		IsEnable:       false,
	}

	return result, nil
}

// Get 获取下一个策略处理器（结束节点通常返回nil）
// 结束节点之后没有其他处理器
func (e *EndNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	log.Printf("处理流程完全结束")

	// 结束节点没有下一个处理器
	return nil, nil
}

// 确保 EndNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*EndNode)(nil)
