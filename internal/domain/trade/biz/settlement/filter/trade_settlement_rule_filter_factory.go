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
	scRuleFilter         *SCRuleFilter
	outTradeNoRuleFilter *OutTradeNoRuleFilter
	settableRuleFilter   *SettableRuleFilter
	endRuleFilter        *EndRuleFilter
}

// NewTradeSettlementRuleFilterFactory 创建交易结算规则过滤工厂
func NewTradeSettlementRuleFilterFactory(
	scRuleFilter *SCRuleFilter,
	outTradeNoRuleFilter *OutTradeNoRuleFilter,
	settableRuleFilter *SettableRuleFilter,
	endRuleFilter *EndRuleFilter,
) *TradeSettlementRuleFilterFactory {
	return &TradeSettlementRuleFilterFactory{
		scRuleFilter:         scRuleFilter,
		outTradeNoRuleFilter: outTradeNoRuleFilter,
		settableRuleFilter:   settableRuleFilter,
		endRuleFilter:        endRuleFilter,
	}
}

// TradeRuleFilter 创建交易结算规则过滤链
func (f *TradeSettlementRuleFilterFactory) TradeRuleFilter() *model2.BusinessLinkedList[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity] {
	// 组装链
	linkArmory := model2.NewLinkArmory(
		"交易结算规则过滤链",
		model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity](f.scRuleFilter),
		model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity](f.outTradeNoRuleFilter),
		model2.ILogicHandler[*model.TradeSettlementRuleCommandEntity, *DynamicContext, *model.TradeSettlementRuleFilterBackEntity](f.settableRuleFilter),
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
