package service

import (
	"context"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/factory"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
)

// IIndexGroupBuyMarketService 拼团营销服务
// 提供对外的营销试算服务接口
type IIndexGroupBuyMarketService struct {
	v1.UnimplementedActivityHTTPServer
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewIIndexGroupBuyMarketService 创建拼团营销服务实例
func NewIIndexGroupBuyMarketService(rootNode *node.RootNode) *IIndexGroupBuyMarketService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &IIndexGroupBuyMarketService{
		strategyFactory: strategyFactory,
	}
}

// IndexMarketTrial 首页营销试算
// 对应Java中的indexMarketTrial方法
func (s *IIndexGroupBuyMarketService) MarketTrial(ctx context.Context, req *v1.MarketTrialRequest) (*v1.MarketTrialReply, error) {
	// 获取策略处理器
	strategyHandler := s.strategyFactory.StrategyHandler()

	// 创建动态上下文
	dynamicContext := &core.DynamicContext{}

	// Create market product entity
	marketProduct := &model.MarketProductEntity{
		UserId:  req.UserId,
		GoodsId: req.GoodsId,
		Source:  req.Source,
		Channel: req.Channel,
	}

	// 应用策略处理器
	trialBalanceEntity, err := strategyHandler.Apply(ctx, marketProduct, dynamicContext)
	if err != nil {
		return nil, err
	}

	// 转换trialBalanceEntity为MarketTrialReply
	reply := &v1.MarketTrialReply{
		TrialResult: &v1.TrialBalanceInfo{
			GoodsId:        trialBalanceEntity.GoodsId,
			GoodsName:      trialBalanceEntity.GoodsName,
			OriginalPrice:  trialBalanceEntity.OriginalPrice,
			DeductionPrice: trialBalanceEntity.DeductionPrice,
			TargetCount:    int32(trialBalanceEntity.TargetCount),
			StartTime:      trialBalanceEntity.StartTime,
			EndTime:        trialBalanceEntity.EndTime,
			IsVisible:      trialBalanceEntity.IsVisible,
			IsEnable:       trialBalanceEntity.IsEnable,
		},
	}

	return reply, nil
}
