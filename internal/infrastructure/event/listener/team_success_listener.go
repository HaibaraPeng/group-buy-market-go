package listener

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"group-buy-market-go/internal/infrastructure/data"
)

// TeamSuccessEventListener 团队成功事件监听器
type TeamSuccessEventListener struct {
	rabbitmqClient *data.RabbitMQClient
	logger         *log.Helper
}

// NewTeamSuccessEventListener 创建团队成功事件监听器实例
func NewTeamSuccessEventListener(rabbitmqClient *data.RabbitMQClient, logger log.Logger) *TeamSuccessEventListener {
	return &TeamSuccessEventListener{
		rabbitmqClient: rabbitmqClient,
		logger:         log.NewHelper(logger),
	}
}

// ListenTeamSuccessEvent 监听团队成功事件
func (l *TeamSuccessEventListener) ListenTeamSuccessEvent(ctx context.Context) error {
	// 声明交换机
	err := l.declareExchange()
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}

	// 声明队列
	queueName := "group_buy_market_queue_2_topic_team_success"
	err = l.declareQueue(queueName)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// 绑定队列到交换机
	routingKey := "topic.team_success"
	err = l.bindQueue(queueName, routingKey)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %w", err)
	}

	// 开始消费消息
	msgs, err := l.rabbitmqClient.Consume(queueName)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	l.logger.Info("开始监听团队成功事件...")

	// 处理消息
	go func() {
		for d := range msgs {
			// 处理接收到的消息
			err := l.handleMessage(d)
			if err != nil {
				l.logger.Errorf("Error handling message: %v", err)
				// 拒绝消息，不重新入队
				d.Nack(false, false)
			} else {
				// 确认消息已处理
				d.Ack(false)
			}
		}
	}()

	return nil
}

// declareExchange 声明交换机
func (l *TeamSuccessEventListener) declareExchange() error {
	return l.rabbitmqClient.GetChannel().ExchangeDeclare(
		"group_buy_market_exchange", // name
		"topic",                     // type
		true,                        // durable
		false,                       // auto-deleted
		false,                       // internal
		false,                       // no-wait
		nil,                         // arguments
	)
}

// declareQueue 声明队列
func (l *TeamSuccessEventListener) declareQueue(queueName string) error {
	_, err := l.rabbitmqClient.GetChannel().QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}

// bindQueue 绑定队列到交换机
func (l *TeamSuccessEventListener) bindQueue(queueName, routingKey string) error {
	return l.rabbitmqClient.GetChannel().QueueBind(
		queueName,                   // queue name
		routingKey,                  // routing key
		"group_buy_market_exchange", // exchange
		false,                       // no-wait
		nil,                         // arguments
	)
}

// handleMessage 处理接收到的消息
func (l *TeamSuccessEventListener) handleMessage(delivery amqp.Delivery) error {
	l.logger.Infof("接收消息: %s", string(delivery.Body))
	return nil
}
