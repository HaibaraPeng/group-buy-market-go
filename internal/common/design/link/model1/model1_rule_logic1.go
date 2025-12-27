package model1

import "github.com/go-kratos/kratos/v2/log"

// Model1RuleLogic1
type Model1RuleLogic1 struct {
	AbstractLogicLink[string, *DynamicContext, string]
	log *log.Helper
}

// NewModel1RuleLogic1
func NewModel1RuleLogic1(logger log.Logger) *Model1RuleLogic1 {
	model1RuleLogic1 := &Model1RuleLogic1{
		log: log.NewHelper(logger),
	}

	// 设置自定义方法实现
	model1RuleLogic1.SetDoApplyFunc(model1RuleLogic1.doApply)

	return model1RuleLogic1
}
