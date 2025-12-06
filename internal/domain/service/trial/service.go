package trial

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/service/trial/factory"
	"group-buy-market-go/internal/domain/service/trial/node"
)

// GroupBuyMarketService 拼团营销服务
// 提供对外的营销试算服务接口
type GroupBuyMarketService struct {
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewGroupBuyMarketService 创建拼团营销服务实例
func NewGroupBuyMarketService() *GroupBuyMarketService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	rootNode := node.NewRootNode()
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &GroupBuyMarketService{
		strategyFactory: strategyFactory,
	}
}

// CalculateTrialBalance 计算试算平衡
// 执行完整的策略树流程，返回最终的试算结果
func (s *GroupBuyMarketService) CalculateTrialBalance(product *model.MarketProductEntity, context *DynamicContext) (*model.TrialBalanceEntity, error) {
	// 获取策略树的根节点
	handler := s.strategyFactory.StrategyHandler()

	var result *model.TrialBalanceEntity
	var err error

	// 循环执行策略树中的各个节点
	for handler != nil {
		// 应用当前节点的策略
		result, err = handler.Apply(product, context)
		if err != nil {
			return nil, err
		}

		// 如果处理失败，提前结束
		if !result.Success {
			return result, nil
		}

		// 获取下一个处理器
		handler, err = handler.Get(product, context)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
