package settlement

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"group-buy-market-go/internal/domain/trade/biz/settlement/filter"
	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TradeSettlementOrderService 拼团交易结算服务
type TradeSettlementOrderService struct {
	log                              *log.Helper
	repository                       *repository.TradeRepository
	tradeSettlementRuleFilterFactory *filter.TradeSettlementRuleFilterFactory
}

// NewTradeSettlementOrderService 创建拼团交易结算服务实例
func NewTradeSettlementOrderService(logger log.Logger, repository *repository.TradeRepository, tradeSettlementRuleFilterFactory *filter.TradeSettlementRuleFilterFactory) *TradeSettlementOrderService {
	return &TradeSettlementOrderService{
		log:                              log.NewHelper(logger),
		repository:                       repository,
		tradeSettlementRuleFilterFactory: tradeSettlementRuleFilterFactory,
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
	}

	// 3. 构建聚合对象
	groupBuyTeamSettlementAggregate := &model.GroupBuyTeamSettlementAggregate{
		UserEntity:            &model.UserEntity{UserId: tradePaySuccessEntity.UserId},
		GroupBuyTeamEntity:    groupBuyTeamEntity,
		TradePaySuccessEntity: tradePaySuccessEntity,
	}

	// 4. 拼团交易结算
	err = s.repository.SettlementMarketPayOrder(ctx, groupBuyTeamSettlementAggregate)
	if err != nil {
		return nil, err
	}

	// 5. 返回结算信息 - 公司中开发这样的流程时候，会根据外部需要进行值的设置
	return &model.TradePaySettlementEntity{
		Source:     tradePaySuccessEntity.Source,
		Channel:    tradePaySuccessEntity.Channel,
		UserId:     tradePaySuccessEntity.UserId,
		TeamId:     teamId,
		ActivityId: groupBuyTeamEntity.ActivityId,
		OutTradeNo: tradePaySuccessEntity.OutTradeNo,
	}, nil
}
