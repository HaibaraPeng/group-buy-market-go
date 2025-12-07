package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
)

// SwitchRoot 开关节点
// 用于判断营销活动是否开启
type SwitchRoot struct {
	core.AbstractGroupBuyMarketSupport
}

// NewSwitchRoot 创建开关节点
func NewSwitchRoot() *SwitchRoot {
	return &SwitchRoot{}
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *SwitchRoot) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	return r.Router(requestParameter, dynamicContext)
}

// 确保 SwitchRoot 实现了 StrategyHandler 接口
var _ core.StrategyHandler = (*SwitchRoot)(nil)
