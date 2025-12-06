package trial

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/factory"
	"group-buy-market-go/internal/domain/service/trial/types"
	"log"
)

// GroupBuyMarketService 拼团营销服务
// 提供对外的营销试算服务接口
type GroupBuyMarketService struct {
	strategyFactory *factory.DefaultActivityStrategyFactory
}

// NewGroupBuyMarketService 创建拼团营销服务实例
func NewGroupBuyMarketService() *GroupBuyMarketService {
	// 构建策略树：根节点 -> 开关节点 -> 营销节点 -> 结束节点
	rootNode := NewRootNode()
	strategyFactory := factory.NewDefaultActivityStrategyFactory(rootNode)

	return &GroupBuyMarketService{
		strategyFactory: strategyFactory,
	}
}

// CalculateTrialBalance 计算试算平衡
// 执行完整的策略树流程，返回最终的试算结果
func (s *GroupBuyMarketService) CalculateTrialBalance(product *model.MarketProductEntity, context *types.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("开始执行营销试算流程，商品ID: %d, 用户ID: %d", product.ID, context.UserID)

	// 获取策略树的根节点
	handler := s.strategyFactory.StrategyHandler()

	var result *model.TrialBalanceEntity
	var err error

	// 循环执行策略树中的各个节点
	for handler != nil {
		// 应用当前节点的策略
		result, err = handler.Apply(product, context)
		if err != nil {
			log.Printf("节点处理出错: %v", err)
			return nil, err
		}

		// 如果处理失败，提前结束
		if !result.Success {
			log.Printf("节点处理失败: %s", result.Message)
			return result, nil
		}

		// 获取下一个处理器
		handler, err = handler.Get(product, context)
		if err != nil {
			log.Printf("获取下一节点出错: %v", err)
			return nil, err
		}
	}

	log.Printf("营销试算流程执行完成")
	return result, nil
}
