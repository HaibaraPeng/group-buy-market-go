package node

import (
	"errors"
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// RootNode 根节点
// 策略树的起始节点，负责初始化处理流程
type RootNode struct {
	core.AbstractGroupBuyMarketSupport
	switchNode *SwitchNode
}

// NewRootNode 创建根节点
func NewRootNode(switchNode *SwitchNode) *RootNode {
	root := &RootNode{
		switchNode: switchNode,
	}

	// 设置自定义方法实现
	root.SetDoApplyFunc(root.doApply)
	root.SetMultiThreadFunc(root.multiThread)
	root.SetDoGet(root.Get)

	return root
}

// multiThread 异步加载数据 - 根节点不需要异步加载
func (r *RootNode) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 根节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *RootNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-RootNode userId:%s requestParameter:%+v", requestParameter.UserId, requestParameter)

	// 参数判断
	if requestParameter == nil || dynamicContext == nil {
		return nil, errors.New("参数不能为空")
	}

	// 注意：Go版本的实体结构与Java版本略有不同，这里根据Go的实际情况进行校验
	if requestParameter.UserId == "" || requestParameter.GoodsId == "" ||
		requestParameter.Source == "" || requestParameter.Channel == "" {
		return nil, errors.New("非法参数: UserId、GoodsId、Source和Channel不能为空")
	}

	return r.Router(requestParameter, dynamicContext)
}

// Get 获取下一个策略处理器
// 根节点之后通常是开关节点，用于判断是否启用营销活动
func (r *RootNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	log.Printf("根节点处理完成，进入开关节点")

	// 返回开关节点作为下一个处理器
	return r.switchNode, nil
}

// 确保 RootNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*RootNode)(nil)
