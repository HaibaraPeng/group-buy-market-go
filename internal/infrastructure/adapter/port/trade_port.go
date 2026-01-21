package port

import (
	"context"
	"time"

	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v8"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/data"
	"group-buy-market-go/internal/infrastructure/event/publish"
	"group-buy-market-go/internal/infrastructure/gateway"
)

// TradePort 交易端口实现
// 对应Java中的TradePort类
type TradePort struct {
	groupBuyNotifyService *gateway.GroupBuyNotifyService
	data                  *data.Data
	redsync               *redsync.Redsync
	publisher             publish.RabbitMQEventPublisher
}

// NewTradePort 创建新的交易端口实例
func NewTradePort(groupBuyNotifyService *gateway.GroupBuyNotifyService, data *data.Data, publisher publish.RabbitMQEventPublisher) *TradePort {
	pool := goredis.NewPool(data.Rdb(context.Background()))
	redsync := redsync.New(pool)

	return &TradePort{
		groupBuyNotifyService: groupBuyNotifyService,
		data:                  data,
		redsync:               redsync,
		publisher:             publisher,
	}
}

// GroupBuyNotify 拼团回调方法
func (t *TradePort) GroupBuyNotify(ctx context.Context, notifyTask *model.NotifyTaskEntity) (string, error) {
	// 创建分布式锁
	mutex := t.redsync.NewMutex(notifyTask.LockKey(), redsync.WithExpiry(3*time.Second), redsync.WithTries(1))

	// 尝试获取锁
	if err := mutex.Lock(); err != nil {
		// 获取锁失败，返回NULL表示未获取到锁
		return string(model.NULL), nil
	}
	defer mutex.Unlock()

	// HTTP回调方式
	if notifyTask.NotifyType == model.HTTP {
		// 无效的 notifyUrl 则直接返回成功
		if notifyTask.NotifyUrl == "" || notifyTask.NotifyUrl == "暂无" {
			return string(model.SUCCESS), nil
		}

		// 调用拼团回调服务
		result, err := t.groupBuyNotifyService.GroupBuyNotify(ctx, notifyTask.NotifyUrl, notifyTask.ParameterJson)
		if err != nil {
			return string(model.NULL), err
		}

		return result, nil
	}

	// MQ回调方式
	if notifyTask.NotifyType == model.MQ {
		// 发布MQ消息
		err := t.publisher.Publish(ctx, notifyTask.NotifyMQ, notifyTask.ParameterJson)
		if err != nil {
			return string(model.NULL), err
		}

		return string(model.SUCCESS), nil
	}

	return string(model.SUCCESS), nil
}
