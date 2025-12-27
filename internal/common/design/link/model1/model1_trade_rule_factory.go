package model1

import "github.com/go-kratos/kratos/v2/log"

// DynamicContext 动态上下文
type DynamicContext struct {
	Age string
}

// Model1TradeRuleFactory 交易规则工厂
type Model1TradeRuleFactory struct {
	ruleLogic101 *Model1RuleLogic1
	ruleLogic102 *Model1RuleLogic2
}

// NewModel1TradeRuleFactory 创建新的Model1TradeRuleFactory实例
func NewModel1TradeRuleFactory(logger log.Logger) *Model1TradeRuleFactory {
	return &Model1TradeRuleFactory{
		ruleLogic101: NewModel1RuleLogic1(logger),
		ruleLogic102: NewModel1RuleLogic2(logger),
	}
}

// OpenLogicLink 打开逻辑链
func (f *Model1TradeRuleFactory) OpenLogicLink() ILogicLink[string, *DynamicContext, string] {
	f.ruleLogic101.AppendNext(f.ruleLogic102)
	return f.ruleLogic101
}
