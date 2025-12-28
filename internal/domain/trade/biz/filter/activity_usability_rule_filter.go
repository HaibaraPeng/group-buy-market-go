package filter

import (
	"context"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"time"
)

// ActivityUsabilityRuleFilter 活动可用性规则过滤器
type ActivityUsabilityRuleFilter struct {
	tradeRepository *repository.TradeRepository
}

// Ensure ActivityUsabilityRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity] = (*ActivityUsabilityRuleFilter)(nil)

// NewActivityUsabilityRuleFilter 创建活动可用性规则过滤器
func NewActivityUsabilityRuleFilter(tradeRepository *repository.TradeRepository) *ActivityUsabilityRuleFilter {
	return &ActivityUsabilityRuleFilter{
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *ActivityUsabilityRuleFilter) Next(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	return f.filter(ctx, command, dynamicContext)
}

// Apply 实现责任链处理器接口
func (f *ActivityUsabilityRuleFilter) Apply(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	return f.filter(ctx, command, dynamicContext)
}

// filter 是实际的过滤逻辑
func (f *ActivityUsabilityRuleFilter) filter(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	// 获取活动信息
	activity, err := f.tradeRepository.QueryGroupBuyActivityEntityByActivityId(ctx, command.ActivityId)
	if err != nil {
		return nil, err
	}

	if activity == nil {
		// 活动不存在，返回默认反馈实体
		return &model.TradeRuleFilterBackEntity{
			UserTakeOrderCount: 0,
		}, nil
	}

	// 检验活动状态 - 可以抛业务异常或把code写入到动态上下文中
	if activity.Status != model.ACTIVITY_ACTIVE {
		// 活动非生效状态
		return &model.TradeRuleFilterBackEntity{
			UserTakeOrderCount: 0,
		}, nil
	}

	// 检验活动时间
	currentTime := time.Now()
	if currentTime.Before(activity.StartTime) || currentTime.After(activity.EndTime) {
		// 活动不在可参与时间范围内
		return &model.TradeRuleFilterBackEntity{
			UserTakeOrderCount: 0,
		}, nil
	}

	// 写入动态上下文
	dynamicContext.GroupBuyActivity = activity

	// 走到下一个责任链节点
	return nil, nil
}
