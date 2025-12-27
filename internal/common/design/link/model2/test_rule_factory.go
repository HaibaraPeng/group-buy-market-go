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
