package model2

// DynamicContext 对应Java版本中的Model2TradeRuleFactory.DynamicContext类
type DynamicContext struct {
	Age string
}

// GetAge 获取年龄
func (dc *DynamicContext) GetAge() string {
	return dc.Age
}

// SetAge 设置年龄
func (dc *DynamicContext) SetAge(age string) {
	dc.Age = age
}

// Model2TradeRuleFactory 是对应Java版本Model2TradeRuleFactory类的Go实现
type Model2TradeRuleFactory struct{}

// NewModel2TradeRuleFactory 创建一个新的工厂实例
func NewModel2TradeRuleFactory() *Model2TradeRuleFactory {
	return &Model2TradeRuleFactory{}
}

// Demo01 创建第一个责任链实例，包含RuleLogic201和RuleLogic202
func (rf *Model2TradeRuleFactory) Demo01() *BusinessLinkedList[string, *DynamicContext, *XxxResponse] {
	ruleLogic201 := NewRuleLogic201()
	ruleLogic202 := NewRuleLogic202()

	linkArmory := NewLinkArmory("demo01",
		ILogicHandler[string, *DynamicContext, *XxxResponse](ruleLogic201),
		ILogicHandler[string, *DynamicContext, *XxxResponse](ruleLogic202),
	)

	return linkArmory.GetLogicLink()
}

// Demo02 创建第二个责任链实例，只包含RuleLogic202
func (rf *Model2TradeRuleFactory) Demo02() *BusinessLinkedList[string, *DynamicContext, *XxxResponse] {
	ruleLogic202 := NewRuleLogic202()

	linkArmory := NewLinkArmory("demo02",
		ILogicHandler[string, *DynamicContext, *XxxResponse](ruleLogic202),
	)

	return linkArmory.GetLogicLink()
}
