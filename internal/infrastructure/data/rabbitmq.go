package data

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"group-buy-market-go/internal/conf"
)

// RabbitMQClient 包装RabbitMQ连接和通道
type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	config  *conf.Data_RabbitMQ
	logger  *log.Helper
	mu      sync.Mutex
}

// NewRabbitMQClient 创建新的RabbitMQ客户端
func NewRabbitMQClient(config *conf.Data, logger log.Logger) (*RabbitMQClient, func(), error) {
	rmqConfig := config.Rabbitmq

	// 构建连接字符串
	addr := rmqConfig.Addresses
	port := rmqConfig.Port
	username := rmqConfig.Username
	password := rmqConfig.Password

	// 确保端口是有效的
	if port == 0 {
		port = 5672 // 默认端口
	}

	// 构建RabbitMQ URL
	urlStr := fmt.Sprintf("amqp://%s:%s@%s:%d",
		url.QueryEscape(username),
		url.QueryEscape(password),
		addr,
		port)

	conn, err := amqp.Dial(urlStr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	client := &RabbitMQClient{
		conn:    conn,
		channel: channel,
		config:  rmqConfig,
		logger:  log.NewHelper(logger),
	}

	// 声明交换机 - 使用生产者配置中的交换机名称
	if rmqConfig.Producer != nil && rmqConfig.Producer.Exchange != "" {
		err = client.DeclareExchange()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to declare exchange: %w", err)
		}
	}

	// 声明队列 - 使用生产者特定配置
	if rmqConfig.Producer != nil && rmqConfig.Producer.TopicTeamSuccess != nil {
		queueName := rmqConfig.Producer.TopicTeamSuccess.Queue
		if queueName != "" {
			// 声明队列
			err = client.DeclareQueueWithName(queueName)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to declare queue: %w", err)
			}
		}

		routingKey := rmqConfig.Producer.TopicTeamSuccess.RoutingKey
		queueName := rmqConfig.Producer.TopicTeamSuccess.Queue
		if routingKey != "" && queueName != "" {
			// 绑定队列到交换机
			err = client.BindQueueWithParams(queueName, routingKey)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to bind queue: %w", err)
			}
		}
	}

	// 设置QoS
	err = channel.Qos(
		int(rmqConfig.PrefetchCount), // prefetch count
		0,                            // prefetch size
		false,                        // global
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	cleanup := func() {
		client.logger.Info("closing RabbitMQ connection")
		if channel != nil {
			channel.Close()
		}
		if conn != nil {
			conn.Close()
		}
	}

	return client, cleanup, nil
}

// DeclareExchange 声明交换机
func (r *RabbitMQClient) DeclareExchange() error {
	// 使用生产者配置中的交换机名称
	exchangeName := r.config.Producer.Exchange

	return r.channel.ExchangeDeclare(
		exchangeName,        // name
		"direct",            // type
		r.config.Durable,    // durable
		r.config.AutoDelete, // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
}

// DeclareQueue 声明队列（使用配置中的默认队列名）
func (r *RabbitMQClient) DeclareQueue() error {
	return r.DeclareQueueWithName(r.config.Queue)
}

// DeclareQueueWithName 声明指定名称的队列
func (r *RabbitMQClient) DeclareQueueWithName(queueName string) error {
	_, err := r.channel.QueueDeclare(
		queueName,           // name
		r.config.Durable,    // durable
		r.config.AutoDelete, // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	return err
}

// BindQueue 绑定队列到交换机（使用配置中的默认参数）
func (r *RabbitMQClient) BindQueue() error {
	return r.BindQueueWithParams(r.config.Queue, r.config.RoutingKey)
}

// BindQueueWithParams 使用指定参数绑定队列到交换机
func (r *RabbitMQClient) BindQueueWithParams(queueName, routingKey string) error {
	return r.channel.QueueBind(
		queueName,         // queue name
		routingKey,        // routing key
		r.config.Exchange, // exchange
		false,             // no-wait
		nil,               // arguments
	)
}

// Publish 发布消息到队列
func (r *RabbitMQClient) Publish(ctx context.Context, routingKey string, queueName string, body []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.channel.PublishWithContext(
		ctx,
		r.config.Exchange, // exchange
		routingKey,        // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.DeliveryMode(r.config.DeliveryMode),
		},
	)

	if err != nil {
		r.logger.Errorf("Failed to publish message: %v", err)
		return err
	}

	r.logger.Info("Message published successfully")
	return nil
}

// PublishWithDefaultRouting 发布使用默认路由键的消息
func (r *RabbitMQClient) PublishWithDefaultRouting(ctx context.Context, body []byte) error {
	return r.Publish(ctx, r.config.RoutingKey, r.config.Queue, body)
}

// PublishTeamSuccessEvent 发布团队成功事件
func (r *RabbitMQClient) PublishTeamSuccessEvent(ctx context.Context, body []byte) error {
	if r.config.Producer != nil && r.config.Producer.TopicTeamSuccess != nil {
		return r.Publish(ctx, r.config.Producer.TopicTeamSuccess.RoutingKey,
			r.config.Producer.TopicTeamSuccess.Queue, body)
	}
	// 如果没有特定配置，返回错误，因为必须有明确的配置
	return fmt.Errorf("no TopicTeamSuccess configuration found")
}

// Consume 订阅消息
func (r *RabbitMQClient) Consume(queue string) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

// GetChannel 获取通道
func (r *RabbitMQClient) GetChannel() *amqp.Channel {
	return r.channel
}

// GetConnection 获取连接
func (r *RabbitMQClient) GetConnection() *amqp.Connection {
	return r.conn
}
