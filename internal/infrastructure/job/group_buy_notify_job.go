package job

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/domain/trade/biz/settlement"
)

// NewGroupBuyNotifyJob 创建拼团完结回调通知任务实例
func NewGroupBuyNotifyJob(logger log.Logger, tradeSettlementOrderService *settlement.TradeSettlementOrderService) *Job {
	var opts = []JobOption{
		Logger(logger),
	}
	srv := NewJob(opts...)

	groupBuyNotifyJob := sync.Once{}
	err := srv.AddFunc("*/15 * * * * *", func(ctx context.Context) {
		l := log.NewHelper(log.With(logger, "module", "groupBuyNotifyJob")).WithContext(ctx)
		groupBuyNotifyJob.Do(func() {
			l.Info("回调通知拼团完结任务开始执行")
		})

		result, err := tradeSettlementOrderService.ExecSettlementNotifyJob(ctx)
		if err != nil {
			l.Errorf("定时任务，回调通知拼团完结任务失败: %v", err)
			return
		}

		l.Infof("定时任务，回调通知拼团完结任务 result:%v", result)
	})
	if err != nil {
		panic(fmt.Sprintf("groupBuyNotifyJob err %s", err))
	}
	return srv
}
