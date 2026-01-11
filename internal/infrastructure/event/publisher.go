package event

import (
	"context"
	"encoding/json"

	"github.com/google/wire"
	"group-buy-market-go/internal/infrastructure/data"
)

// EventPublisher 定义事件发布器接口
type EventPublisher interface {
	PublishTeamSuccessEvent(ctx context.Context, payload interface{}) error
	Publish(ctx context.Context, routingKey string, payload interface{}) error
}

// RabbitMQEventPublisher 基于RabbitMQ的事件发布器实现
type RabbitMQEventPublisher struct {
	rabbitmqClient *data.RabbitMQClient
}

// NewRabbitMQEventPublisher 创建RabbitMQ事件发布器实例
func NewRabbitMQEventPublisher(rabbitmqClient *data.RabbitMQClient) *RabbitMQEventPublisher {
	return &RabbitMQEventPublisher{
		rabbitmqClient: rabbitmqClient,
	}
}

// Publish 发布消息到指定路由键
func (p *RabbitMQEventPublisher) Publish(ctx context.Context, routingKey string, payload interface{}) error {
	// 序列化消息体
	messageBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// 发布消息
	err = p.rabbitmqClient.Publish(ctx, routingKey, "", messageBody)
	if err != nil {
		return err
	}

	return nil
}

// PublishTeamSuccessEvent 发布团队成功事件
func (p *RabbitMQEventPublisher) PublishTeamSuccessEvent(ctx context.Context, payload interface{}) error {
	// 使用配置中的团队成功事件路由键
	return p.Publish(ctx, "topic.team_success", payload)
}

// ProviderSet 提供依赖注入集合
var ProviderSet = wire.NewSet(NewRabbitMQEventPublisher)
