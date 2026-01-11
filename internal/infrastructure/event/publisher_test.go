//go:build integration

package event

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"group-buy-market-go/internal/conf"
	"group-buy-market-go/internal/infrastructure/data"
)

// setupTestData 创建测试用的Data实例
func setupTestData() *data.Data {
	// 创建测试用的Mysql客户端（使用nil表示不实际连接）
	var db *gorm.DB

	// 创建测试用的Redis客户端（使用nil表示不实际连接）
	var redisClient *redis.Client

	// 创建测试用的RabbitMQ客户端
	cfg := &conf.Data{
		Rabbitmq: &conf.Data_RabbitMQ{
			Addresses:     "localhost",
			Port:          5672,
			Username:      "admin",
			Password:      "admin",
			PrefetchCount: 1,
			Producer: &conf.Data_RabbitMQProducer{
				Exchange: "group_buy_market_exchange",
				TopicTeamSuccess: &conf.Data_RabbitMQTopicConfig{
					RoutingKey: "topic.team_success",
					Queue:      "group_buy_market_queue_2_topic_team_success",
				},
			},
		},
	}

	logger := log.DefaultLogger
	rabbitmqClient, _, err := data.NewRabbitMQClient(cfg, logger)
	if err != nil {
		// 如果无法连接到RabbitMQ，则跳过测试
		return nil
	}

	d := data.NewData(db, redisClient, rabbitmqClient)
	return d
}

// TestRabbitMQEventPublisher_RealConnection 测试真实RabbitMQ连接
func TestRabbitMQEventPublisher_RealConnection(t *testing.T) {
	testData := setupTestData()
	if testData == nil {
		t.Skip("Cannot connect to RabbitMQ, skipping integration test")
	}

	// 创建事件发布器
	publisher := NewRabbitMQEventPublisher(testData.Rmq(context.Background()))

	// 测试发布多个消息，类似Java测试用例
	messages := []string{
		"订单结算：ORD-20231234",
		"订单结算：ORD-20231235",
		"订单结算：ORD-20231236",
		"订单结算：ORD-20231237",
		"订单结算：ORD-20231238",
	}

	for _, msg := range messages {
		err := publisher.Publish(context.Background(), "topic.team_success", msg)
		assert.NoError(t, err, "Publish should not return an error for message: %s", msg)
		time.Sleep(10 * time.Millisecond) // 小延迟，避免消息发送过快
	}
}
