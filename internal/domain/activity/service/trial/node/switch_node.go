package node

import (
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// SwitchNode 开关节点
// 用于判断营销活动是否开启
type SwitchNode struct {
	core.AbstractGroupBuyMarketSupport
	marketNode *MarketNode
}

// NewSwitchNode 创建开关节点
func NewSwitchNode(marketNode *MarketNode) *SwitchNode {
	switchNode := &SwitchNode{
		marketNode: marketNode,
	}

	// 设置自定义方法实现
	switchNode.SetDoApplyFunc(switchNode.doApply)
	switchNode.SetMultiThreadFunc(switchNode.multiThread)

	return switchNode
}

// multiThread 异步加载数据 - 开关节点不需要异步加载
func (r *SwitchNode) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 开关节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *SwitchNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-SwitchNode requestParameter:%+v", requestParameter)

	// todo xfg 判断营销活动开关是否打开

	return r.Router(requestParameter, dynamicContext)
}

// Get 获取下一个策略处理器
func (r *SwitchNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	log.Printf("开关节点处理完成，进入营销节点")

	// 返回营销节点作为下一个处理器
	return r.marketNode, nil
}

// 确保 SwitchNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*SwitchNode)(nil)
