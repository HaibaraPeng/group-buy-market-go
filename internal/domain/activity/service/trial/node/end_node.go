package node

import (
	"context"
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"

	"github.com/go-kratos/kratos/v2/log"
)

// EndNode 结束节点
// 策略树的终止节点，负责收尾工作和返回最终结果
type EndNode struct {
	core.AbstractGroupBuyMarketSupport
	log *log.Helper
}

// NewEndNode 创建结束节点
func NewEndNode(logger log.Logger) *EndNode {
	endNode := &EndNode{
		log: log.NewHelper(logger),
	}

	// 设置自定义方法实现
	endNode.SetDoApplyFunc(endNode.doApply)
	endNode.SetMultiThreadFunc(endNode.multiThread)
	endNode.SetDoGet(endNode.Get)

	return endNode
}

// multiThread 异步加载数据 - 结束节点不需要异步加载
func (e *EndNode) multiThread(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 结束节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (e *EndNode) doApply(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	e.log.Infow("拼团商品查询试算服务-EndNode", "userId", requestParameter.UserId, "requestParameter", requestParameter)

	groupBuyActivityDiscountVO := dynamicContext.GetGroupBuyActivityDiscountVO()
	skuVO := dynamicContext.GetSkuVO()
	deductionPrice := dynamicContext.GetDeductionPrice()

	// 如果没有计算出折扣价格，默认为0
	deductionPriceValue := 0.0
	if deductionPrice != nil {
		deductionPriceValue, _ = deductionPrice.Float64()
	}

	// 返回结果
	result := &model.TrialBalanceEntity{
		GoodsId:                    skuVO.GoodsId,
		GoodsName:                  skuVO.GoodsName,
		OriginalPrice:              skuVO.OriginalPrice,
		DeductionPrice:             deductionPriceValue,
		TargetCount:                groupBuyActivityDiscountVO.Target,
		StartTime:                  groupBuyActivityDiscountVO.StartTime.Unix(),
		EndTime:                    groupBuyActivityDiscountVO.EndTime.Unix(),
		IsVisible:                  dynamicContext.IsVisible(),
		IsEnable:                   dynamicContext.IsEnable(),
		GroupBuyActivityDiscountVO: *dynamicContext.GetGroupBuyActivityDiscountVO(),
	}

	return result, nil
}

// Get 获取下一个策略处理器（结束节点通常返回nil）
// 结束节点之后没有其他处理器
func (e *EndNode) Get(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	e.log.Info("处理流程完全结束")

	// 结束节点没有下一个处理器
	return nil, nil
}

// 确保 EndNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*EndNode)(nil)
