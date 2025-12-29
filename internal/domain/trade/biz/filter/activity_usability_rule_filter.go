package filter

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/common/consts"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/common/exception"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"time"
)

// ActivityUsabilityRuleFilter 活动可用性规则过滤器
type ActivityUsabilityRuleFilter struct {
	log             *log.Helper
	tradeRepository *repository.TradeRepository
}

// Ensure ActivityUsabilityRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity] = (*ActivityUsabilityRuleFilter)(nil)

// NewActivityUsabilityRuleFilter 创建活动可用性规则过滤器
func NewActivityUsabilityRuleFilter(logger log.Logger, tradeRepository *repository.TradeRepository) *ActivityUsabilityRuleFilter {
	return &ActivityUsabilityRuleFilter{
		log:             log.NewHelper(logger),
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
	f.log.Infow("交易规则过滤-活动的可用性校验", "userId", command.UserId, "activityId", command.ActivityId)

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

	// 校验；活动状态 - 可以抛业务异常code，或者把code写入到动态上下文dynamicContext中，最后获取。
	if activity.Status != model.ACTIVITY_ACTIVE {
		f.log.Infow("活动的可用性校验，非生效状态", "activityId", command.ActivityId)
		return nil, exception.NewAppException(consts.E0101)
	}

	// 检验活动时间
	currentTime := time.Now()
	if currentTime.Before(activity.StartTime) || currentTime.After(activity.EndTime) {
		f.log.Infow("活动的可用性校验，非可参与时间范围", "activityId", command.ActivityId)
		return nil, exception.NewAppException(consts.E0102)
	}

	// 写入动态上下文
	dynamicContext.GroupBuyActivity = activity

	// 走到下一个责任链节点
	return f.Next(ctx, command, dynamicContext)
}
