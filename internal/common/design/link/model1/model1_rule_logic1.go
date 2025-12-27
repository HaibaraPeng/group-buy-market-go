package model1

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Model1RuleLogic1 实现了具体的业务逻辑处理器
type Model1RuleLogic1 struct {
	*AbstractLogicLink[string, *DynamicContext, string]
	log *log.Helper
}

// NewModel1RuleLogic1 创建新的Model1RuleLogic1实例
func NewModel1RuleLogic1(logger log.Logger) *Model1RuleLogic1 {
	model1RuleLogic1 := &Model1RuleLogic1{
		log: log.NewHelper(logger),
	}

	// 设置自定义方法实现
	model1RuleLogic1.AbstractLogicLink = &AbstractLogicLink[string, *DynamicContext, string]{}
	model1RuleLogic1.SetDoApplyFunc(model1RuleLogic1.doApply)

	return model1RuleLogic1
}

// doApply 实现具体的业务逻辑
func (r *Model1RuleLogic1) doApply(ctx context.Context, requestParameter string, dynamicContext *DynamicContext) (string, error) {
	r.log.Info("link model01 RuleLogic1")

	// 调用下一个节点
	return r.NextLink(requestParameter, dynamicContext)
}
