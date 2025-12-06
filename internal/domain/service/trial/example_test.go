package trial_test

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
	"group-buy-market-go/internal/domain/service/trial"
	"testing"
)

func TestGroupBuyMarketService_CalculateTrialBalance(t *testing.T) {
	// 创建服务实例
	service := trial.NewGroupBuyMarketService()

	// 准备测试数据
	product := &model.MarketProductEntity{
		ID:          1001,
		Name:        "测试商品",
		Description: "这是一个测试商品",
		Price:       100.0,
		SkuID:       2001,
		Stock:       100,
	}

	context := &core.DynamicContext{
		UserID:     12345,
		ActivityID: 5001,
		UserLevel:  2, // 黄金用户
		Timestamp:  1234567890,
		ClientIP:   "127.0.0.1",
	}

	// 执行试算
	result, err := service.CalculateTrialBalance(product, context)
	if err != nil {
		t.Fatalf("计算试算平衡出错: %v", err)
	}

	// 验证结果
	if !result.Success {
		t.Errorf("试算应该成功，但实际失败了: %s", result.Message)
	}

	// 验证返回了正确的消息
	expectedMessage := "营销活动处理完成"
	if result.Message != expectedMessage {
		t.Errorf("期望的消息应该是 '%s'，但实际是 '%s'", expectedMessage, result.Message)
	}
}

func TestRootNode_Get(t *testing.T) {
	// 测试根节点获取下一个处理器
	rootNode := node.NewRootNode()

	product := &model.MarketProductEntity{
		ID:    1001,
		Price: 100.0,
	}

	context := &core.DynamicContext{
		UserID: 12345,
	}

	nextHandler, err := rootNode.Get(product, context)
	if err != nil {
		t.Fatalf("获取下一处理器出错: %v", err)
	}

	// 验证下一个处理器是否为开关节点
	if _, ok := nextHandler.(*node.SwitchRoot); !ok {
		t.Error("根节点的下一个处理器应该是开关节点")
	}
}
