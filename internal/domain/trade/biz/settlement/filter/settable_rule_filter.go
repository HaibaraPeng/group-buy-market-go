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

// SettableRuleFilter 可结算规则过滤；交易时间
type SettableRuleFilter struct {
	log             *log.Helper
	tradeRepository *repository.TradeRepository
}

// Ensure SettableRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] = (*SettableRuleFilter)(nil)

// NewSettableRuleFilter 创建可结算规则过滤器
func NewSettableRuleFilter(logger log.Logger, tradeRepository *repository.TradeRepository) *SettableRuleFilter {
	return &SettableRuleFilter{
		log:             log.NewHelper(logger),
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *SettableRuleFilter) Next(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	return nil, nil
}

// Apply 实现责任链处理器接口
func (f *SettableRuleFilter) Apply(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	f.log.Infow("结算规则过滤-有效时间校验", "userId", command.UserId, "outTradeNo", command.OutTradeNo)

	// 上下文；获取数据
	marketPayOrderEntity := dynamicContext.MarketPayOrderEntity

	// 查询拼团对象
	groupBuyTeamEntity, err := f.tradeRepository.QueryGroupBuyTeamByTeamId(ctx, marketPayOrderEntity.TeamId)
	if err != nil {
		return nil, err
	}

	if groupBuyTeamEntity == nil {
		return nil, exception.NewAppException(consts.UN_ERROR)
	}

	// 外部交易时间 - 也就是用户支付完成的时间，这个时间要在拼团有效时间范围内
	outTradeTime := command.OutTradeTime

	// 判断，外部交易时间，要小于拼团结束时间。否则抛异常。
	if !outTradeTime.Before(groupBuyTeamEntity.ValidEndTime) {
		f.log.Errorw("订单交易时间不在拼团有效时间范围内")
		return nil, exception.NewAppException(consts.E0106)
	}

	// 设置上下文
	dynamicContext.GroupBuyTeamEntity = groupBuyTeamEntity

	return f.Next(ctx, command, dynamicContext)
}
