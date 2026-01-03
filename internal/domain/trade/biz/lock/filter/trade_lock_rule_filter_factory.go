package filter

import (
	"context"
	"group-buy-market-go/internal/common/design/link/model2"
	"group-buy-market-go/internal/domain/trade/model"
)

// DynamicContext 交易规则过滤工厂动态上下文
type DynamicContext struct {
	GroupBuyActivity *model.GroupBuyActivityEntity `json:"groupBuyActivity"`
}

// TradeLockRuleFilterFactory 交易规则过滤工厂
type TradeLockRuleFilterFactory struct {
	activityUsabilityRuleFilter *ActivityUsabilityRuleFilter
	UserTakeLimitRuleFilter     *UserTakeLimitRuleFilter
}

// NewTradeLockRuleFilterFactory 创建交易规则过滤工厂
func NewTradeLockRuleFilterFactory(
	activityUsabilityRuleFilter *ActivityUsabilityRuleFilter,
	userTakeLimitRuleFilter *UserTakeLimitRuleFilter,
) *TradeLockRuleFilterFactory {
	return &TradeLockRuleFilterFactory{
		activityUsabilityRuleFilter: activityUsabilityRuleFilter,
		UserTakeLimitRuleFilter:     userTakeLimitRuleFilter,
	}
}

// TradeRuleFilter 创建交易规则过滤链
func (f *TradeLockRuleFilterFactory) TradeRuleFilter() *model2.BusinessLinkedList[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity] {
	// 组装链
	linkArmory := model2.NewLinkArmory(
		"交易规则过滤链",
		model2.ILogicHandler[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity](f.activityUsabilityRuleFilter),
		model2.ILogicHandler[*model.TradeRuleCommandEntity, *DynamicContext, *model.TradeRuleFilterBackEntity](f.UserTakeLimitRuleFilter),
	)

	// 链对象
	return linkArmory.GetLogicLink()
}

// Execute 执行过滤链
func (f *TradeLockRuleFilterFactory) Execute(ctx context.Context, command *model.TradeRuleCommandEntity, dynamicContext *DynamicContext) (*model.TradeRuleFilterBackEntity, error) {
	businessLinkedList := f.TradeRuleFilter()
	return businessLinkedList.Apply(ctx, command, dynamicContext)
}
