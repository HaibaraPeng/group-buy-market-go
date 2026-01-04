package job

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/domain/trade/biz/settlement"
)

// GroupBuyNotifyJob 拼团完结回调通知任务
type GroupBuyNotifyJob struct {
	log                         *log.Helper
	tradeSettlementOrderService *settlement.TradeSettlementOrderService
}

// NewGroupBuyNotifyJob 创建拼团完结回调通知任务实例
func NewGroupBuyNotifyJob(logger log.Logger, tradeSettlementOrderService *settlement.TradeSettlementOrderService) *GroupBuyNotifyJob {
	return &GroupBuyNotifyJob{
		log:                         log.NewHelper(logger),
		tradeSettlementOrderService: tradeSettlementOrderService,
	}
}

// Exec 执行定时任务，回调通知拼团完结
func (g *GroupBuyNotifyJob) Exec() {
	ctx := context.Background()

	g.log.Info("定时任务，回调通知拼团完结任务开始执行")

	result, err := g.tradeSettlementOrderService.ExecSettlementNotifyJob(ctx)
	if err != nil {
		g.log.Errorf("定时任务，回调通知拼团完结任务失败: %v", err)
		return
	}

	g.log.Infof("定时任务，回调通知拼团完结任务 result:%v", result)
}

// Run 启动定时任务
func (g *GroupBuyNotifyJob) Run() {
	// 每15秒执行一次，对应Java中的cron表达式 "0/15 * * * * ?"
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			g.Exec()
		}
	}
}
