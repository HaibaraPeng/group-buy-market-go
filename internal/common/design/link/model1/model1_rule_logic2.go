package model1

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Model1RuleLogic2 实现了具体的业务逻辑处理器
type Model1RuleLogic2 struct {
	*AbstractLogicLink[string, *DynamicContext, string]
	log *log.Helper
}

// NewModel1RuleLogic2 创建新的Model1RuleLogic2实例
func NewModel1RuleLogic2(logger log.Logger) *Model1RuleLogic2 {
	model1RuleLogic2 := &Model1RuleLogic2{
		log: log.NewHelper(logger),
	}

	// 设置自定义方法实现
	model1RuleLogic2.AbstractLogicLink = &AbstractLogicLink[string, *DynamicContext, string]{}
	model1RuleLogic2.SetDoApplyFunc(model1RuleLogic2.doApply)

	return model1RuleLogic2
}

// doApply 实现具体的业务逻辑
func (r *Model1RuleLogic2) doApply(ctx context.Context, requestParameter string, dynamicContext *DynamicContext) (string, error) {
	r.log.Info("link model01 RuleLogic2")

	// 调用下一个节点
	return r.NextLink(requestParameter, dynamicContext)
}
