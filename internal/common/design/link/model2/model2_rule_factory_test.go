package model2

import (
	"testing"
)

func TestModel2TradeRuleFactory(t *testing.T) {
	factory := NewRule02TradeRuleFactory()

	// 准备测试数据
	requestParam := "test request"
	dynamicContext := &Rule02TradeRuleFactoryDynamicContext{Age: "25"}

	// 测试Demo01 - 包含RuleLogic201和RuleLogic202的责任链
	businessLinkedList := factory.Demo01()
	result, err := businessLinkedList.Apply(requestParam, dynamicContext)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 验证结果 - RuleLogic202会返回"hi 小傅哥！"响应
	if result == nil {
		t.Error("Expected result to not be nil")
	} else if result.GetAge() != "hi 小傅哥！" {
		t.Errorf("Expected age 'hi 小傅哥！', got '%s'", result.GetAge())
	}

	// 测试Demo02 - 只包含RuleLogic202的责任链
	businessLinkedList2 := factory.Demo02()
	result2, err := businessLinkedList2.Apply(requestParam, dynamicContext)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 验证结果
	if result2 == nil {
		t.Error("Expected result2 to not be nil")
	} else if result2.GetAge() != "hi 小傅哥！" {
		t.Errorf("Expected age 'hi 小傅哥！', got '%s'", result2.GetAge())
	}
}
