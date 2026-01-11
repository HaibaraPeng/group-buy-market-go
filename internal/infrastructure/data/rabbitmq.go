package data

import (
	"context"
	"fmt"
	"net/url"
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
		queueName = rmqConfig.Producer.TopicTeamSuccess.Queue
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
		exchangeName, // name
		"topic",      // topic
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

// DeclareQueueWithName 声明指定名称的队列
func (r *RabbitMQClient) DeclareQueueWithName(queueName string) error {
	// 如果队列名为空，则使用默认队列名
	if queueName == "" {
		queueName = "default_queue" // 或者从配置中获取默认队列名
	}
	_, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}

// BindQueue 绑定队列到交换机（使用配置中的默认参数）
func (r *RabbitMQClient) BindQueue() error {
	// 这里需要提供具体的队列名和路由键，目前不适用
	return fmt.Errorf("BindQueue not applicable with current config structure")
}

// BindQueueWithParams 使用指定参数绑定队列到交换机
func (r *RabbitMQClient) BindQueueWithParams(queueName, routingKey string) error {
	exchangeName := r.config.Producer.Exchange
	return r.channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}

// Publish 发布消息到队列
func (r *RabbitMQClient) Publish(ctx context.Context, routingKey string, queueName string, body []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.channel.PublishWithContext(
		ctx,
		r.config.Producer.Exchange, // exchange
		routingKey,                 // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent,
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
	// 这里需要更具体的实现，可能需要配置中的默认值
	// 使用默认的路由键和队列名，这里我们使用TopicTeamSuccess的配置
	if r.config.Producer != nil && r.config.Producer.TopicTeamSuccess != nil {
		return r.Publish(ctx, r.config.Producer.TopicTeamSuccess.RoutingKey,
			r.config.Producer.TopicTeamSuccess.Queue, body)
	}
	return fmt.Errorf("no default routing configuration found")
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
