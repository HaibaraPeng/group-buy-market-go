package lock

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"group-buy-market-go/internal/domain/trade/biz/lock/filter"

	"group-buy-market-go/internal/domain/trade/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
)

// TradeLockOrder 交易订单用例
type TradeLockOrder struct {
	repo                   *repository.TradeRepository
	tradeRuleFilterFactory *filter.TradeRuleFilterFactory
	log                    *log.Helper
}

// NewTradeLockOrder 创建交易订单用例实例
func NewTradeLockOrder(repo *repository.TradeRepository, tradeRuleFilterFactory *filter.TradeRuleFilterFactory, logger log.Logger) *TradeLockOrder {
	return &TradeLockOrder{
		repo:                   repo,
		tradeRuleFilterFactory: tradeRuleFilterFactory,
		log:                    log.NewHelper(logger),
	}
}

// QueryNoPayMarketPayOrderByOutTradeNo 查询未支付的营销订单
func (uc *TradeLockOrder) QueryNoPayMarketPayOrderByOutTradeNo(ctx context.Context, userId string, outTradeNo string) (*model.MarketPayOrderEntity, error) {
	uc.log.Infof("拼团交易-查询未支付营销订单:%s outTradeNo:%s", userId, outTradeNo)
	return uc.repo.QueryMarketPayOrderEntityByOutTradeNo(ctx, userId, outTradeNo)
}

// QueryGroupBuyProgress 查询拼团进度
func (uc *TradeLockOrder) QueryGroupBuyProgress(ctx context.Context, teamId string) (*model.GroupBuyProgressVO, error) {
	uc.log.Infof("拼团交易-查询拼单进度:%s", teamId)
	return uc.repo.QueryGroupBuyProgress(ctx, teamId)
}

// LockMarketPayOrder 锁定营销支付订单
func (uc *TradeLockOrder) LockMarketPayOrder(ctx context.Context, userEntity *model.UserEntity, payActivityEntity *model.PayActivityEntity, payDiscountEntity *model.PayDiscountEntity) (*model.MarketPayOrderEntity, error) {
	uc.log.Infof("拼团交易-锁定营销优惠支付订单:%s activityId:%d goodsId:%s", userEntity.UserId, payActivityEntity.ActivityId, payDiscountEntity.GoodsId)

	// 交易规则过滤
	tradeRuleFilterBackEntity, err := uc.tradeRuleFilterFactory.Execute(ctx,
		&model.TradeRuleCommandEntity{
			ActivityId: payActivityEntity.ActivityId,
			UserId:     userEntity.UserId,
		},
		&filter.DynamicContext{},
	)
	if err != nil {
		return nil, err
	}

	// 已参与拼团量 - 用于构建数据库唯一索引使用，确保用户只能在一个活动上参与固定的次数
	userTakeOrderCount := tradeRuleFilterBackEntity.UserTakeOrderCount

	// 构建聚合对象
	groupBuyOrderAggregate := &model.GroupBuyOrderAggregate{
		UserEntity:         userEntity,
		PayActivityEntity:  payActivityEntity,
		PayDiscountEntity:  payDiscountEntity,
		UserTakeOrderCount: userTakeOrderCount,
	}

	// 锁定聚合订单 - 这会用户只是下单还没有支付。后续会有2个流程；支付成功、超时未支付（回退）
	return uc.repo.LockMarketPayOrder(ctx, groupBuyOrderAggregate)
}
