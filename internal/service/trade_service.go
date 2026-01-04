package service

import (
	"context"
	"fmt"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
	"group-buy-market-go/internal/domain/trade/biz/lock"

	"github.com/go-kratos/kratos/v2/log"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/factory"
	trade_model "group-buy-market-go/internal/domain/trade/model"
)

// TradeService 交易服务
// 提供对外的交易服务接口
type TradeService struct {
	v1.UnimplementedTradeHTTPServer
	log             *log.Helper
	tradeLockOrder  *lock.TradeLockOrder
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewTradeService 创建交易服务实例
func NewTradeService(logger log.Logger, tradeLockOrder *lock.TradeLockOrder, rootNode *node.RootNode) *TradeService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)
	return &TradeService{
		log:             log.NewHelper(logger),
		tradeLockOrder:  tradeLockOrder,
		strategyFactory: strategyFactory,
	}
}

// LockMarketPayOrder 锁定营销支付订单
func (s *TradeService) LockMarketPayOrder(ctx context.Context, req *v1.LockMarketPayOrderRequest) (*v1.LockMarketPayOrderReply, error) {
	// 参数
	userId := req.UserId
	source := req.Source
	channel := req.Channel
	goodsId := req.GoodsId
	activityId := req.ActivityId
	outTradeNo := req.OutTradeNo
	teamId := req.TeamId
	notifyUrl := req.NotifyUrl

	s.log.WithContext(ctx).Infof("营销交易锁单:%s LockMarketPayOrderRequest:%+v", userId, req)

	// 检查必要参数
	if userId == "" || source == "" || channel == "" || goodsId == "" || activityId == 0 {
		return nil, fmt.Errorf("非法参数")
	}

	// 查询 outTradeNo 是否已经存在交易记录
	marketPayOrderEntity, err := s.tradeLockOrder.QueryNoPayMarketPayOrderByOutTradeNo(ctx, userId, outTradeNo)
	if err != nil {
		return nil, fmt.Errorf("查询未支付营销订单失败: %w", err)
	}

	// 如果已存在订单记录，直接返回
	if marketPayOrderEntity != nil {
		s.log.WithContext(ctx).Infof("交易锁单记录(存在):%s marketPayOrderEntity:%+v", userId, marketPayOrderEntity)
		return &v1.LockMarketPayOrderReply{
			OrderId:          marketPayOrderEntity.OrderId,
			DeductionPrice:   marketPayOrderEntity.DeductionPrice,
			TradeOrderStatus: int32(marketPayOrderEntity.TradeOrderStatusEnumVO.Code()),
		}, nil
	}

	// 判断拼团锁单是否完成了目标
	if teamId != "" {
		groupBuyProgressVO, err := s.tradeLockOrder.QueryGroupBuyProgress(ctx, teamId)
		if err != nil {
			return nil, fmt.Errorf("查询拼团进度失败: %w", err)
		}

		if groupBuyProgressVO != nil && groupBuyProgressVO.TargetCount == groupBuyProgressVO.LockCount {
			s.log.WithContext(ctx).Infof("交易锁单拦截-拼单目标已达成:%s %s", userId, teamId)
			return nil, fmt.Errorf("拼单目标已达成")
		}
	}

	// 营销优惠试算
	strategyHandler := s.strategyFactory.StrategyHandler()
	dynamicContext := &core.DynamicContext{}

	marketProduct := &model.MarketProductEntity{
		UserId:     userId,
		Source:     source,
		Channel:    channel,
		GoodsId:    goodsId,
		ActivityId: activityId,
	}

	trialBalanceEntity, err := strategyHandler.Apply(ctx, marketProduct, dynamicContext)
	if err != nil {
		return nil, fmt.Errorf("营销优惠试算失败: %w", err)
	}

	// 人群限定检查
	if !trialBalanceEntity.IsVisible || !trialBalanceEntity.IsEnable {
		return nil, fmt.Errorf("人群限定检查失败")
	}

	groupBuyActivityDiscountVO := trialBalanceEntity.GroupBuyActivityDiscountVO

	// 锁单
	marketPayOrderEntity, err = s.tradeLockOrder.LockMarketPayOrder(ctx,
		&trade_model.UserEntity{UserId: userId},
		&trade_model.PayActivityEntity{
			TeamId:       teamId,
			ActivityId:   activityId,
			ActivityName: groupBuyActivityDiscountVO.ActivityName,
			StartTime:    groupBuyActivityDiscountVO.StartTime,
			EndTime:      groupBuyActivityDiscountVO.EndTime,
			TargetCount:  groupBuyActivityDiscountVO.Target,
		},
		&trade_model.PayDiscountEntity{
			Source:         source,
			Channel:        channel,
			GoodsId:        goodsId,
			GoodsName:      trialBalanceEntity.GoodsName,
			OriginalPrice:  trialBalanceEntity.OriginalPrice,
			DeductionPrice: trialBalanceEntity.DeductionPrice,
			PayPrice:       trialBalanceEntity.PayPrice,
			OutTradeNo:     outTradeNo,
			NotifyUrl:      notifyUrl,
		})
	if err != nil {
		return nil, fmt.Errorf("锁单失败: %w", err)
	}

	s.log.WithContext(ctx).Infof("交易锁单记录(新):%s marketPayOrderEntity:%+v", userId, marketPayOrderEntity)

	// 返回结果
	return &v1.LockMarketPayOrderReply{
		OrderId:          marketPayOrderEntity.OrderId,
		DeductionPrice:   marketPayOrderEntity.DeductionPrice,
		TradeOrderStatus: int32(marketPayOrderEntity.TradeOrderStatusEnumVO.Code()),
	}, nil
}
