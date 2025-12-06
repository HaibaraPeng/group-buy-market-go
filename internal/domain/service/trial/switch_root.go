package trial

import (
	"group-buy-market-go/internal/domain/service/trial/types"
	"log"
)

// SwitchRoot 开关节点
// 用于判断营销活动是否开启
type SwitchRoot struct {
	AbstractGroupBuyMarketSupport
}

// NewSwitchRoot 创建开关节点
func NewSwitchRoot() *SwitchRoot {
	return &SwitchRoot{}
}

// Apply 应用开关节点策略
// 判断活动是否开启，如果未开启则直接返回错误结果
func (s *SwitchRoot) Apply(requestParameter *types.MarketProductEntity, dynamicContext *types.DynamicContext) (*types.TrialBalanceEntity, error) {
	log.Printf("检查营销活动开关状态，活动ID: %d", dynamicContext.ActivityID)

	// 这里应该查询数据库或者配置中心判断活动是否开启
	// 暂时模拟为活动总是开启

	isActivityEnabled := true // 模拟活动开启状态

	if !isActivityEnabled {
		return &types.TrialBalanceEntity{
			Success: false,
			Message: "营销活动未开启",
			Code:    "ACTIVITY_DISABLED",
		}, nil
	}

	return &types.TrialBalanceEntity{
		Success: true,
		Message: "营销活动已开启",
	}, nil
}

// Get 获取下一个策略处理器
// 如果活动开启，则进入营销节点；否则进入结束节点
func (s *SwitchRoot) Get(requestParameter *types.MarketProductEntity, dynamicContext *types.DynamicContext) (types.StrategyHandler, error) {
	log.Printf("开关节点处理完成，进入营销节点")

	// 返回营销节点作为下一个处理器
	marketNode := NewMarketNode()
	return marketNode, nil
}

// 确保 SwitchRoot 实现了 StrategyHandler 接口
var _ types.StrategyHandler = (*SwitchRoot)(nil)
