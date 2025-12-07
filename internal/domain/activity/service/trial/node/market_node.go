package node

import (
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/thread"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"log"
)

// MarketNode 营销优惠节点
// 负责计算各种营销优惠
type MarketNode struct {
	core.AbstractGroupBuyMarketSupport
	activityRepository repository.ActivityRepository
	endNode            *EndNode
}

// NewMarketNode 创建营销节点
func NewMarketNode(activityRepository repository.ActivityRepository) *MarketNode {
	return &MarketNode{
		activityRepository: activityRepository,
		endNode:            NewEndNode(),
	}
}

// multiThread 异步加载数据
// 对应Java中的multiThread方法
func (m *MarketNode) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 异步查询活动配置
	queryActivityTask := thread.NewQueryGroupBuyActivityDiscountVOThreadTask(
		"", // source参数需要从requestParameter获取，此处简化处理
		"", // channel参数需要从requestParameter获取，此处简化处理
		m.activityRepository,
	)

	// 异步查询商品信息
	querySkuTask := thread.NewQuerySkuVOFromDBThreadTask(
		"", // goodsId参数需要从requestParameter获取，此处简化处理
		m.activityRepository,
	)

	// 启动异步任务
	activityChan := queryActivityTask.AsyncCall()
	skuChan := querySkuTask.AsyncCall()

	// 等待并收集结果
	var activityVO *model.GroupBuyActivityDiscountVO
	var skuVO *model.SkuVO

	// 等待活动查询结果
	activityResult := <-activityChan
	if activityResult.Error != nil {
		log.Printf("查询活动配置失败: %v", activityResult.Error)
		return activityResult.Error
	}
	activityVO = activityResult.Result

	// 等待SKU查询结果
	skuResult := <-skuChan
	if skuResult.Error != nil {
		log.Printf("查询商品信息失败: %v", skuResult.Error)
		return skuResult.Error
	}
	skuVO = skuResult.Result

	// 写入上下文 - 对于一些复杂场景，获取数据的操作，
	// 有时候会在下N个节点获取，这样前置查询数据，可以提高接口响应效率
	// 注意：由于Go中struct无法像Java那样动态添加属性，这里仅作示意
	_ = activityVO
	_ = skuVO

	log.Printf("拼团商品查询试算服务-MarketNode异步线程加载数据完成")
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (m *MarketNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-MarketNode requestParameter:%+v", requestParameter)

	// todo xfg 拼团优惠试算

	return m.Router(requestParameter, dynamicContext)
}

// 确保 MarketNode 实现了 StrategyHandler 接口
var _ core.StrategyHandler = (*MarketNode)(nil)
