package filter

import (
	"context"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/domain/trade/model"
)

// DynamicContext 交易规则过滤工厂动态上下文
type DynamicContext struct {
	MarketPayOrderEntity *model.MarketPayOrderEntity `json:"marketPayOrderEntity"`
	GroupBuyTeamEntity   *model.GroupBuyTeamEntity   `json:"groupBuyTeamEntity"`
}

// TradeSettlementRuleFilterFactory 交易结算规则过滤工厂
type TradeSettlementRuleFilterFactory struct {
	outTradeNoRuleFilter *OutTradeNoRuleFilter
	endRuleFilter        *EndRuleFilter
}

// NewTradeSettlementRuleFilterFactory 创建交易结算规则过滤工厂
func NewTradeSettlementRuleFilterFactory(
	outTradeNoRuleFilter *OutTradeNoRuleFilter,
	endRuleFilter *EndRuleFilter,
) *TradeSettlementRuleFilterFactory {
	return &TradeSettlementRuleFilterFactory{
		outTradeNoRuleFilter: outTradeNoRuleFilter,
		endRuleFilter:        endRuleFilter,
	}
}

// TradeRuleFilter 创建交易结算规则过滤链
func (f *TradeSettlementRuleFilterFactory) TradeRuleFilter() *model2.BusinessLinkedList[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] {
	// 组装链
	linkArmory := model2.NewLinkArmory(
		"交易结算规则过滤链",
		model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity](f.outTradeNoRuleFilter),
		model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity](f.endRuleFilter),
	)

	// 链对象
	return linkArmory.GetLogicLink()
}

// Execute 执行过滤链
func (f *TradeSettlementRuleFilterFactory) Execute(ctx context.Context, command *model.TradeSettlementRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeSettlementRuleFilterBackEntity, error) {
	businessLinkedList := f.TradeRuleFilter()
	return businessLinkedList.Apply(ctx, command, dynamicContext)
}
