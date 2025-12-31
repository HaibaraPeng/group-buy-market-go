package settlement

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TradeSettlementOrderService 拼团交易结算服务
type TradeSettlementOrderService struct {
	log        *log.Helper
	repository *repository.TradeRepository
}

// NewTradeSettlementOrderService 创建拼团交易结算服务实例
func NewTradeSettlementOrderService(logger log.Logger, repository *repository.TradeRepository) *TradeSettlementOrderService {
	return &TradeSettlementOrderService{
		log:        log.NewHelper(logger),
		repository: repository,
	}
}

// SettlementMarketPayOrder 拼团交易结算
func (s *TradeSettlementOrderService) SettlementMarketPayOrder(ctx context.Context, tradePaySuccessEntity *model.TradePaySuccessEntity) (*model.TradePaySettlementEntity, error) {
	s.log.WithContext(ctx).Infof("拼团交易-支付订单结算: userId=%s outTradeNo=%s", tradePaySuccessEntity.UserId, tradePaySuccessEntity.OutTradeNo)

	// 1. 查询拼团信息
	marketPayOrderEntity, err := s.repository.QueryMarketPayOrderEntityByOutTradeNo(ctx, tradePaySuccessEntity.UserId, tradePaySuccessEntity.OutTradeNo)
	if err != nil {
		return nil, err
	}
	if marketPayOrderEntity == nil {
		s.log.WithContext(ctx).Infof("不存在的外部交易单号或用户已退单，不需要做支付订单结算: userId=%s outTradeNo=%s", tradePaySuccessEntity.UserId, tradePaySuccessEntity.OutTradeNo)
		return nil, nil
	}

	// 2. 查询组团信息
	groupBuyTeamEntity, err := s.repository.QueryGroupBuyTeamByTeamId(ctx, marketPayOrderEntity.TeamId)
	if err != nil {
		return nil, err
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
		TeamId:     marketPayOrderEntity.TeamId,
		ActivityId: groupBuyTeamEntity.ActivityId,
		OutTradeNo: tradePaySuccessEntity.OutTradeNo,
	}, nil
}
