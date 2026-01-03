package filter

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/domain/trade/model"
)

// EndRuleFilter 结束节点
type EndRuleFilter struct {
	log *log.Helper
}

// Ensure EndRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] = (*EndRuleFilter)(nil)

// NewEndRuleFilter 创建结束节点过滤器
func NewEndRuleFilter(logger log.Logger) *EndRuleFilter {
	return &EndRuleFilter{
		log: log.NewHelper(logger),
	}
}

// Next 实现责任链处理器接口
func (f *EndRuleFilter) Next(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	return nil, nil
}

// Apply 实现责任链处理器接口
func (f *EndRuleFilter) Apply(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	f.log.Infow("结算规则过滤-结束节点", "userId", command.UserId, "outTradeNo", command.OutTradeNo)

	// 获取上下文对象
	groupBuyTeamEntity := dynamicContext.GroupBuyTeamEntity

	// 返回封装数据
	return &model.TradeSettlementRuleFilterBackEntity{
		TeamId:         groupBuyTeamEntity.TeamId,
		ActivityId:     groupBuyTeamEntity.ActivityId,
		TargetCount:    groupBuyTeamEntity.TargetCount,
		CompleteCount:  groupBuyTeamEntity.CompleteCount,
		LockCount:      groupBuyTeamEntity.LockCount,
		Status:         groupBuyTeamEntity.Status,
		ValidStartTime: groupBuyTeamEntity.ValidStartTime,
		ValidEndTime:   groupBuyTeamEntity.ValidEndTime,
	}, nil
}
