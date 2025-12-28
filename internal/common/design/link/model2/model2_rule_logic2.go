package model2

import (
	"context"
	"log"
)

// Model2RuleLogic2 对应Java版本的Model2RuleLogic2类
type Model2RuleLogic2 struct{}

// NewModel2RuleLogic2 创建一个新的Model2RuleLogic2实例
func NewModel2RuleLogic2() *Model2RuleLogic2 {
	return &Model2RuleLogic2{}
}

// Apply 是Model2RuleLogic2的业务逻辑实现
func (rl *Model2RuleLogic2) Apply(ctx context.Context, requestParameter string, dynamicContext *DynamicContext) (*XxxResponse, error) {
	log.Println("link model02 Model2RuleLogic2")

	return NewXxxResponse("hi 小傅哥！"), nil
}

// Next 是Model2RuleLogic2的下一个处理逻辑
func (rl *Model2RuleLogic2) Next(ctx context.Context, requestParameter string, dynamicContext *DynamicContext) (*XxxResponse, error) {
	// Model2RuleLogic2直接返回结果，不调用下一个处理器
	return NewXxxResponse("hi 小傅哥！"), nil
}
