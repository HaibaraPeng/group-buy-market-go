package link

import (
	"testing"
)

// 测试用的简单结构
type TestRequest struct {
	Value string
}

type TestContext struct {
	Data map[string]interface{}
}

type TestResponse struct {
	Result string
}

func TestChainOfResponsibility(t *testing.T) {
	// 创建处理器链
	handler1 := NewExampleHandler[TestRequest, TestContext, TestResponse]("handler1")
	handler2 := NewExampleHandler[TestRequest, TestContext, TestResponse]("handler2")
	handler3 := NewExampleHandler[TestRequest, TestContext, TestResponse]("handler3")

	// 构建链式结构
	handler1.AppendNext(handler2)
	handler2.AppendNext(handler3)

	// 准备测试数据
	req := TestRequest{Value: "test"}
	ctx := TestContext{Data: map[string]interface{}{"key": "value"}}

	// 执行链式调用
	_, err := handler1.Apply(req, ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 验证链式结构
	if handler1.Next() != handler2 {
		t.Errorf("Expected handler1.Next() to be handler2")
	}

	if handler2.Next() != handler3 {
		t.Errorf("Expected handler2.Next() to be handler3")
	}

	if handler3.Next() != nil {
		t.Errorf("Expected handler3.Next() to be nil")
	}
}

func TestChainWithSingleHandler(t *testing.T) {
	// 创建单个处理器
	handler := NewExampleHandler[TestRequest, TestContext, TestResponse]("single")

	// 准备测试数据
	req := TestRequest{Value: "test"}
	ctx := TestContext{Data: map[string]interface{}{"key": "value"}}

	// 执行调用
	_, err := handler.Apply(req, ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
