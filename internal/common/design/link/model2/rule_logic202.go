package model2

import "log"

// RuleLogic202 对应Java版本的RuleLogic202类
type RuleLogic202 struct{}

// NewRuleLogic202 创建一个新的RuleLogic202实例
func NewRuleLogic202() *RuleLogic202 {
	return &RuleLogic202{}
}

// Apply 是RuleLogic202的业务逻辑实现
func (rl *RuleLogic202) Apply(requestParameter string, dynamicContext *Rule02TradeRuleFactoryDynamicContext) (*XxxResponse, error) {
	log.Println("link model02 RuleLogic202")

	return NewXxxResponse("hi 小傅哥！"), nil
}

// Next 是RuleLogic202的下一个处理逻辑
func (rl *RuleLogic202) Next(requestParameter string, dynamicContext *Rule02TradeRuleFactoryDynamicContext) (*XxxResponse, error) {
	// RuleLogic202直接返回结果，不调用下一个处理器
	return NewXxxResponse("hi 小傅哥！"), nil
}
