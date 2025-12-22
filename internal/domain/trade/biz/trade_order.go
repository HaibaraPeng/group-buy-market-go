package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TradeOrder 交易订单用例
type TradeOrder struct {
	repo *repository.TradeRepository
	log  *log.Helper
}

// NewTradeOrder 创建交易订单用例实例
func NewTradeOrder(repo *repository.TradeRepository, logger log.Logger) *TradeOrder {
	return &TradeOrder{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// QueryNoPayMarketPayOrderByOutTradeNo 查询未支付的营销订单
func (uc *TradeOrder) QueryNoPayMarketPayOrderByOutTradeNo(ctx context.Context, userId string, outTradeNo string) (*model.MarketPayOrderEntity, error) {
	uc.log.Infof("拼团交易-查询未支付营销订单:%s outTradeNo:%s", userId, outTradeNo)
	return uc.repo.QueryMarketPayOrderEntityByOutTradeNo(ctx, userId, outTradeNo)
}

// QueryGroupBuyProgress 查询拼团进度
func (uc *TradeOrder) QueryGroupBuyProgress(ctx context.Context, teamId string) (*model.GroupBuyProgressVO, error) {
	uc.log.Infof("拼团交易-查询拼单进度:%s", teamId)
	return uc.repo.QueryGroupBuyProgress(ctx, teamId)
}

// LockMarketPayOrder 锁定营销支付订单
func (uc *TradeOrder) LockMarketPayOrder(ctx context.Context, userEntity *model.UserEntity, payActivityEntity *model.PayActivityEntity, payDiscountEntity *model.PayDiscountEntity) (*model.MarketPayOrderEntity, error) {
	uc.log.Infof("拼团交易-锁定营销优惠支付订单:%s activityId:%d goodsId:%s", userEntity.UserId, payActivityEntity.ActivityId, payDiscountEntity.GoodsId)

	// 构建聚合对象
	groupBuyOrderAggregate := &model.GroupBuyOrderAggregate{
		UserEntity:        userEntity,
		PayActivityEntity: payActivityEntity,
		PayDiscountEntity: payDiscountEntity,
	}

	// 锁定聚合订单 - 这会用户只是下单还没有支付。后续会有2个流程；支付成功、超时未支付（回退）
	return uc.repo.LockMarketPayOrder(ctx, groupBuyOrderAggregate)
}
