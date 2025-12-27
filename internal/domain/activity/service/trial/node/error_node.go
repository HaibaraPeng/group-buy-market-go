package node

import (
	"context"
	"fmt"
	"group-buy-market-go/internal/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"

	"github.com/go-kratos/kratos/v2/log"
)

// ErrorNode 异常节点处理
// 无营销、流程降级、超时调用等，都可以路由到 ErrorNode 节点统一处理
type ErrorNode struct {
	core.AbstractGroupBuyMarketSupport
	defaultStrategyHandler tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity]
	log                    *log.Helper
}

// NewErrorNode 创建异常节点
func NewErrorNode(logger log.Logger) *ErrorNode {
	errorNode := &ErrorNode{
		defaultStrategyHandler: tree.NewDefaultStrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity](),
		log:                    log.NewHelper(logger),
	}

	// 设置自定义方法实现
	errorNode.SetDoApplyFunc(errorNode.doApply)
	errorNode.SetMultiThreadFunc(errorNode.multiThread)
	errorNode.SetDoGet(errorNode.Get)

	return errorNode
}

// multiThread 异步加载数据 - 异常节点不需要异步加载
func (e *ErrorNode) multiThread(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 异常节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (e *ErrorNode) doApply(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	e.log.Infow("拼团商品查询试算服务-ErrorNode", "userId", requestParameter.UserId, "requestParameter", requestParameter)

	// 无营销配置
	if dynamicContext.GetGroupBuyActivityDiscountVO() == nil || dynamicContext.GetSkuVO() == nil {
		e.log.Infow("商品无拼团营销配置", "goodsId", requestParameter.GoodsId)
		return nil, fmt.Errorf("商品无拼团营销配置")
	}

	return &model.TrialBalanceEntity{}, nil
}

// Get 获取下一个策略处理器
func (e *ErrorNode) Get(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	return e.defaultStrategyHandler, nil
}

// 确保 ErrorNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*ErrorNode)(nil)
