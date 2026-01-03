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

// OutTradeNoRuleFilter 外部交易单号过滤；外部交易单号是否为退单
type OutTradeNoRuleFilter struct {
	log             *log.Helper
	tradeRepository *repository.TradeRepository
}

// Ensure OutTradeNoRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] = (*OutTradeNoRuleFilter)(nil)

// NewOutTradeNoRuleFilter 创建外部交易单号过滤器
func NewOutTradeNoRuleFilter(logger log.Logger, tradeRepository *repository.TradeRepository) *OutTradeNoRuleFilter {
	return &OutTradeNoRuleFilter{
		log:             log.NewHelper(logger),
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *OutTradeNoRuleFilter) Next(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	return nil, nil
}

// Apply 实现责任链处理器接口
func (f *OutTradeNoRuleFilter) Apply(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	f.log.Infow("结算规则过滤-外部单号校验", "userId", command.UserId, "outTradeNo", command.OutTradeNo)

	// 查询拼团信息
	marketPayOrderEntity, err := f.tradeRepository.QueryMarketPayOrderEntityByOutTradeNo(ctx, command.UserId, command.OutTradeNo)
	if err != nil {
		return nil, err
	}

	if marketPayOrderEntity == nil {
		f.log.Errorw("不存在的外部交易单号或用户已退单，不需要做支付订单结算", "userId", command.UserId, "outTradeNo", command.OutTradeNo)
		return nil, exception.NewAppException(consts.E0104)
	}

	dynamicContext.MarketPayOrderEntity = marketPayOrderEntity

	return f.Next(ctx, command, dynamicContext)
}
