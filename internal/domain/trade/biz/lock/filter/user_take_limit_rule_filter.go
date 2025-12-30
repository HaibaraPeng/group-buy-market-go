package filter

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/common/consts"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/common/exception"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// UserTakeLimitRuleFilter 用户参与限制规则过滤器
type UserTakeLimitRuleFilter struct {
	log             *log.Helper
	tradeRepository *repository.TradeRepository
}

// Ensure UserTakeLimitRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity] = (*UserTakeLimitRuleFilter)(nil)

// NewUserTakeLimitRuleFilter 创建用户参与限制规则过滤器
func NewUserTakeLimitRuleFilter(logger log.Logger, tradeRepository *repository.TradeRepository) *UserTakeLimitRuleFilter {
	return &UserTakeLimitRuleFilter{
		log:             log.NewHelper(logger),
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *UserTakeLimitRuleFilter) Next(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	return nil, nil
}

// Apply 实现责任链处理器接口
func (f *UserTakeLimitRuleFilter) Apply(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	f.log.Infow("交易规则过滤-用户参与次数校验", "userId", command.UserId, "activityId", command.ActivityId)

	// 从动态上下文获取活动信息
	groupBuyActivity := dynamicContext.GroupBuyActivity

	// 查询用户在一个拼团活动上参与的次数
	count, err := f.tradeRepository.QueryOrderCountByActivityId(ctx, command.ActivityId, command.UserId)
	if err != nil {
		return nil, err
	}

	if groupBuyActivity != nil && groupBuyActivity.TakeLimitCount != 0 && count >= groupBuyActivity.TakeLimitCount {
		f.log.Infow("用户参与次数校验，已达可参与上限", "activityId", command.ActivityId)
		return nil, exception.NewAppException(consts.E0103)
	}

	// 用户未达到参与限制，继续执行下个过滤器
	return &model.TradeRuleFilterBackEntity{
		UserTakeOrderCount: count,
	}, nil
}
