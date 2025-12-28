package filter

import (
	"context"
	"group-buy-market-go/internal/common/consts"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/common/exception"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// UserTakeLimitRuleFilter 用户参与限制规则过滤器
type UserTakeLimitRuleFilter struct {
	tradeRepository *repository.TradeRepository
}

// Ensure UserTakeLimitRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeRuleCommandEntity, *TradeRuleFilterFactoryDynamicContext, *model.TradeRuleFilterBackEntity] = (*UserTakeLimitRuleFilter)(nil)

// NewUserTakeLimitRuleFilter 创建用户参与限制规则过滤器
func NewUserTakeLimitRuleFilter(tradeRepository *repository.TradeRepository) *UserTakeLimitRuleFilter {
	return &UserTakeLimitRuleFilter{
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *UserTakeLimitRuleFilter) Next(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *TradeRuleFilterFactoryDynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	return f.filter(ctx, command, dynamicContext)
}

// Apply 实现责任链处理器接口
func (f *UserTakeLimitRuleFilter) Apply(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *TradeRuleFilterFactoryDynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	return f.filter(ctx, command, dynamicContext)
}

// filter 是实际的过滤逻辑
func (f *UserTakeLimitRuleFilter) filter(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *TradeRuleFilterFactoryDynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	// 查询用户参与活动的订单量
	userTakeOrderCount, err := f.tradeRepository.QueryOrderCountByActivityId(ctx, command.ActivityId, command.UserId)
	if err != nil {
		return nil, err
	}

	// 检查活动参与限制
	if dynamicContext.GroupBuyActivity != nil && userTakeOrderCount >= dynamicContext.GroupBuyActivity.TakeLimitCount {
		// 用户参与次数达到限制，抛出异常
		return nil, exception.NewAppException(consts.E0103)
	}

	// 用户未达到参与限制，继续执行下个过滤器
	return &model.TradeRuleFilterBackEntity{
		UserTakeOrderCount: userTakeOrderCount,
	}, nil
}
