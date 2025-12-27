package model1

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestModel1TradeRuleFactory(t *testing.T) {
	logger := log.DefaultLogger
	factory := NewModel1TradeRuleFactory(logger)

	// 创建逻辑链
	logicChain := factory.OpenLogicLink()

	// 准备测试数据
	requestParam := "test request"
	dynamicContext := &DynamicContext{Age: "25"}

	// 执行链式调用
	result, err := logicChain.Apply(requestParam, dynamicContext)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 验证结果
	expected := "link model01 单实例链"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// 验证链式结构
	if logicChain.Next() == nil {
		t.Error("Expected next handler to exist in chain")
	}
}
