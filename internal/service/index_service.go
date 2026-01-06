package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "group-buy-market-go/api/v1"
	"group-buy-market-go/internal/common/consts"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/factory"
	"group-buy-market-go/internal/domain/activity/service/trial/node"
)

// IndexService 营销首页服务
// 提供对外的营销首页服务接口
type IndexService struct {
	v1.UnimplementedIndexHTTPServer
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewIndexService 创建营销首页服务实例
func NewIndexService(rootNode *node.RootNode) *IndexService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &IndexService{
		strategyFactory: strategyFactory,
	}
}

// QueryGroupBuyMarketConfig 查询拼团营销配置
// 对应Java中的queryGroupBuyMarketConfig方法
func (s *IndexService) QueryGroupBuyMarketConfig(ctx context.Context, req *v1.QueryGroupBuyMarketConfigRequest) (*v1.QueryGroupBuyMarketConfigReply, error) {
	log.Infof("查询拼团营销配置开始 userId:%s goodsId:%s", req.GetUserId(), req.GetGoodsId())

	// 参数校验
	if req.GetUserId() == "" || req.GetSource() == "" || req.GetChannel() == "" || req.GetGoodsId() == "" {
		return nil, fmt.Errorf("%s: %s", string(consts.ILLEGAL_PARAMETER), consts.ILLEGAL_PARAMETER.GetErrorMessage())
	}

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
	return &v1.QueryGroupBuyMarketConfigReply{}, nil
}
