package node

import (
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// SwitchRoot 开关节点
// 用于判断营销活动是否开启
type SwitchRoot struct {
	core.AbstractGroupBuyMarketSupport
	marketNode *MarketNode
}

// NewSwitchRoot 创建开关节点
func NewSwitchRoot() *SwitchRoot {
	return &SwitchRoot{}
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *SwitchRoot) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-SwitchRoot requestParameter:%+v", requestParameter)

	// todo xfg 判断营销活动开关是否打开

	return r.Router(requestParameter, dynamicContext)
}

// Get 获取下一个策略处理器
func (r *SwitchRoot) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	log.Printf("开关节点处理完成，进入营销节点")

	// 返回营销节点作为下一个处理器
	return r.marketNode, nil
}

// 确保 SwitchRoot 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*SwitchRoot)(nil)
