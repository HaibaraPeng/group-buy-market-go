package model2

// Rule02TradeRuleFactoryDynamicContext 对应Java版本中的Rule02TradeRuleFactory.DynamicContext类
type Rule02TradeRuleFactoryDynamicContext struct {
	Age string
}

// NewRule02TradeRuleFactoryDynamicContext 创建一个新的动态上下文实例
func NewRule02TradeRuleFactoryDynamicContext(age string) *Rule02TradeRuleFactoryDynamicContext {
	return &Rule02TradeRuleFactoryDynamicContext{
		Age: age,
	}
}

// GetAge 获取年龄
func (dc *Rule02TradeRuleFactoryDynamicContext) GetAge() string {
	return dc.Age
}

// SetAge 设置年龄
func (dc *Rule02TradeRuleFactoryDynamicContext) SetAge(age string) {
	dc.Age = age
}

// Rule02TradeRuleFactory 是对应Java版本Rule02TradeRuleFactory类的Go实现
type Rule02TradeRuleFactory struct{}

// NewRule02TradeRuleFactory 创建一个新的工厂实例
func NewRule02TradeRuleFactory() *Rule02TradeRuleFactory {
	return &Rule02TradeRuleFactory{}
}

// Demo01 创建第一个责任链实例，包含RuleLogic201和RuleLogic202
func (rf *Rule02TradeRuleFactory) Demo01() *BusinessLinkedList[string, *Rule02TradeRuleFactoryDynamicContext, *XxxResponse] {
	ruleLogic201 := NewRuleLogic201()
	ruleLogic202 := NewRuleLogic202()

	linkArmory := NewLinkArmory("demo01",
		ILogicHandler[string, *Rule02TradeRuleFactoryDynamicContext, *XxxResponse](ruleLogic201),
		ILogicHandler[string, *Rule02TradeRuleFactoryDynamicContext, *XxxResponse](ruleLogic202),
	)

	return linkArmory.GetLogicLink()
}

// Demo02 创建第二个责任链实例，只包含RuleLogic202
func (rf *Rule02TradeRuleFactory) Demo02() *BusinessLinkedList[string, *Rule02TradeRuleFactoryDynamicContext, *XxxResponse] {
	ruleLogic202 := NewRuleLogic202()

	linkArmory := NewLinkArmory("demo02",
		ILogicHandler[string, *Rule02TradeRuleFactoryDynamicContext, *XxxResponse](ruleLogic202),
	)

	return linkArmory.GetLogicLink()
}
