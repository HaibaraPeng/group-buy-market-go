package settlement

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/infrastructure/adapter/port"
	"group-buy-market-go/internal/infrastructure/gateway"

	"group-buy-market-go/internal/domain/trade/biz/settlement/filter"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TradeSettlementOrderService 拼团交易结算服务
type TradeSettlementOrderService struct {
	log                              *log.Helper
	repository                       *repository.TradeRepository
	port                             *port.TradePort
	tradeSettlementRuleFilterFactory *filter.TradeSettlementRuleFilterFactory
	notifyService                    *gateway.GroupBuyNotifyService
}

// NewTradeSettlementOrderService 创建拼团交易结算服务实例
func NewTradeSettlementOrderService(logger log.Logger, repository *repository.TradeRepository, port *port.TradePort,
	tradeSettlementRuleFilterFactory *filter.TradeSettlementRuleFilterFactory, notifyService *gateway.GroupBuyNotifyService) *TradeSettlementOrderService {
	return &TradeSettlementOrderService{
		log:                              log.NewHelper(logger),
		repository:                       repository,
		port:                             port,
		tradeSettlementRuleFilterFactory: tradeSettlementRuleFilterFactory,
		notifyService:                    notifyService,
	}
}

// SettlementMarketPayOrder 拼团交易结算
func (s *TradeSettlementOrderService) SettlementMarketPayOrder(ctx context.Context, tradePaySuccessEntity *model.TradePaySuccessEntity) (*model.TradePaySettlementEntity, error) {
	s.log.WithContext(ctx).Infof("拼团交易-支付订单结算: userId=%s outTradeNo=%s", tradePaySuccessEntity.UserId, tradePaySuccessEntity.OutTradeNo)

	// 1. 结算规则过滤
	tradeSettlementRuleFilterBackEntity, err := s.tradeSettlementRuleFilterFactory.Execute(ctx,
		&model.TradeSettlementRuleCommandEntity{
			Source:       tradePaySuccessEntity.Source,
			Channel:      tradePaySuccessEntity.Channel,
			UserId:       tradePaySuccessEntity.UserId,
			OutTradeNo:   tradePaySuccessEntity.OutTradeNo,
			OutTradeTime: tradePaySuccessEntity.OutTradeTime,
		},
		&filter.DynamicContext{},
	)
	if err != nil {
		return nil, err
	}

	teamId := tradeSettlementRuleFilterBackEntity.TeamId

	// 2. 查询组团信息
	groupBuyTeamEntity := &model.GroupBuyTeamEntity{
		TeamId:         tradeSettlementRuleFilterBackEntity.TeamId,
		ActivityId:     tradeSettlementRuleFilterBackEntity.ActivityId,
		TargetCount:    tradeSettlementRuleFilterBackEntity.TargetCount,
		CompleteCount:  tradeSettlementRuleFilterBackEntity.CompleteCount,
		LockCount:      tradeSettlementRuleFilterBackEntity.LockCount,
		Status:         tradeSettlementRuleFilterBackEntity.Status,
		ValidStartTime: tradeSettlementRuleFilterBackEntity.ValidStartTime,
		ValidEndTime:   tradeSettlementRuleFilterBackEntity.ValidEndTime,
		NotifyUrl:      tradeSettlementRuleFilterBackEntity.NotifyUrl,
	}

	// 3. 构建聚合对象
	groupBuyTeamSettlementAggregate := &model.GroupBuyTeamSettlementAggregate{
		UserEntity:            &model.UserEntity{UserId: tradePaySuccessEntity.UserId},
		GroupBuyTeamEntity:    groupBuyTeamEntity,
		TradePaySuccessEntity: tradePaySuccessEntity,
	}

	// 4. 拼团交易结算
	isCompleted, err := s.repository.SettlementMarketPayOrder(ctx, groupBuyTeamSettlementAggregate)
	if err != nil {
		return nil, err
	}

	// 5. 组队回调处理 - 处理失败也会有定时任务补偿，通过这样的方式，可以减轻任务调度，提高时效性
	if isCompleted {
		notifyResultMap, err := s.ExecSettlementNotifyJobByTeamId(ctx, teamId)
		if err != nil {
			return nil, err
		}
		s.log.WithContext(ctx).Infof("回调通知拼团完结 result:%v", notifyResultMap)
	}

	// 6. 返回结算信息 - 公司中开发这样的流程时候，会根据外部需要进行值的设置
	return &model.TradePaySettlementEntity{
		Source:     tradePaySuccessEntity.Source,
		Channel:    tradePaySuccessEntity.Channel,
		UserId:     tradePaySuccessEntity.UserId,
		TeamId:     teamId,
		ActivityId: groupBuyTeamEntity.ActivityId,
		OutTradeNo: tradePaySuccessEntity.OutTradeNo,
	}, nil
}

// ExecSettlementNotifyJob 执行结算通知任务
func (s *TradeSettlementOrderService) ExecSettlementNotifyJob(ctx context.Context) (map[string]int, error) {
	s.log.Info("拼团交易-执行结算通知任务")

	// 查询未执行任务
	notifyTaskEntityList, err := s.repository.QueryUnExecutedNotifyTaskList(context.Background())
	if err != nil {
		return nil, err
	}

	return s.execSettlementNotifyJob(ctx, notifyTaskEntityList)
}

// ExecSettlementNotifyJobByTeamId 执行结算通知回调，指定 teamId
func (s *TradeSettlementOrderService) ExecSettlementNotifyJobByTeamId(ctx context.Context, teamId string) (map[string]int, error) {
	s.log.Infof("拼团交易-执行结算通知回调，指定 teamId:%s", teamId)

	notifyTaskEntity, err := s.repository.QueryUnExecutedNotifyTaskByTeamId(context.Background(), teamId)
	if err != nil {
		return nil, err
	}

	if notifyTaskEntity == nil {
		return s.execSettlementNotifyJob(ctx, []*model.NotifyTaskEntity{})
	}

	return s.execSettlementNotifyJob(ctx, []*model.NotifyTaskEntity{notifyTaskEntity})
}

// execSettlementNotifyJob 执行结算通知任务的具体实现
func (s *TradeSettlementOrderService) execSettlementNotifyJob(ctx context.Context, notifyTaskEntityList []*model.NotifyTaskEntity) (map[string]int, error) {
	successCount := 0
	errorCount := 0
	retryCount := 0

	for _, notifyTask := range notifyTaskEntityList {
		// 回调处理 success 成功，error 失败
		response, err := s.port.GroupBuyNotify(ctx, notifyTask)
		if err != nil {
			s.log.Errorf("执行回调通知失败: %v", err)
			response = "error"
		}

		// 更新状态判断&变更数据库表回调任务状态
		if response == "success" {
			err = s.repository.UpdateNotifyTaskStatusSuccess(context.Background(), notifyTask.TeamId)
			if err == nil {
				successCount += 1
			}
		} else if response == "error" {
			if notifyTask.NotifyCount < 5 {
				err = s.repository.UpdateNotifyTaskStatusError(context.Background(), notifyTask.TeamId)
				if err == nil {
					errorCount += 1
				}
			} else {
				err = s.repository.UpdateNotifyTaskStatusRetry(context.Background(), notifyTask.TeamId)
				if err == nil {
					retryCount += 1
				}
			}
		}
	}

	resultMap := map[string]int{
		"waitCount":    len(notifyTaskEntityList),
		"successCount": successCount,
		"errorCount":   errorCount,
		"retryCount":   retryCount,
	}

	return resultMap, nil
}
