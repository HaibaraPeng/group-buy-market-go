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

// SCRuleFilter SC 渠道来源过滤 - 当某个签约渠道下架后，则不会记账
type SCRuleFilter struct {
	log             *log.Helper
	tradeRepository *repository.TradeRepository
}

// Ensure SCRuleFilter implements model2.ILogicHandler
var _ model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] = (*SCRuleFilter)(nil)

// NewSCRuleFilter 创建SC渠道过滤器
func NewSCRuleFilter(logger log.Logger, tradeRepository *repository.TradeRepository) *SCRuleFilter {
	return &SCRuleFilter{
		log:             log.NewHelper(logger),
		tradeRepository: tradeRepository,
	}
}

// Next 实现责任链处理器接口
func (f *SCRuleFilter) Next(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	return nil, nil
}

// Apply 实现责任链处理器接口
func (f *SCRuleFilter) Apply(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	f.log.Infow("结算规则过滤-渠道黑名单校验", "userId", command.UserId, "outTradeNo", command.OutTradeNo)

	// sc 渠道黑名单拦截
	intercept := f.tradeRepository.IsSCBlackIntercept(command.Source, command.Channel)
	if intercept {
		f.log.Errorw("渠道黑名单拦截", "source", command.Source, "channel", command.Channel)
		return nil, exception.NewAppException(consts.E0105)
	}

	return f.Next(ctx, command, dynamicContext)
}
