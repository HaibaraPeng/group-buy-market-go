package node

import (
	"fmt"
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/discount"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/domain/activity/service/trial/thread"
	"group-buy-market-go/internal/infrastructure/adapter/repository"
	"log"
	"math/big"
)

// MarketNode 营销优惠节点
// 负责计算各种营销优惠
type MarketNode struct {
	core.AbstractGroupBuyMarketSupport
	activityRepository          *repository.ActivityRepository
	endNode                     *EndNode
	discountCalculateServiceMap map[model.MarketPlanEnum]discount.IDiscountCalculateService
	zjCalculateService          *discount.ZJCalculateService
	zkCalculateService          *discount.ZKCalculateService
	mjCalculateService          *discount.MJCalculateService
	nCalculateService           *discount.NCalculateService
}

// NewMarketNode 创建营销节点
func NewMarketNode(
	endNode *EndNode,
	activityRepository *repository.ActivityRepository,
	zjCalculateService *discount.ZJCalculateService,
	zkCalculateService *discount.ZKCalculateService,
	mjCalculateService *discount.MJCalculateService,
	nCalculateService *discount.NCalculateService,
) *MarketNode {
	marketNode := &MarketNode{
		activityRepository: activityRepository,
		endNode:            endNode,
		zjCalculateService: zjCalculateService,
		zkCalculateService: zkCalculateService,
		mjCalculateService: mjCalculateService,
		nCalculateService:  nCalculateService,
	}

	// 初始化折扣计算服务映射
	marketNode.discountCalculateServiceMap = map[model.MarketPlanEnum]discount.IDiscountCalculateService{
		model.ZJ: zjCalculateService,
		model.ZK: zkCalculateService,
		model.MJ: mjCalculateService,
		model.N:  nCalculateService,
	}

	// 设置自定义方法实现
	marketNode.SetDoApplyFunc(marketNode.doApply)
	marketNode.SetMultiThreadFunc(marketNode.multiThread)
	marketNode.SetDoGet(marketNode.Get)

	return marketNode
}

// multiThread 异步加载数据
// 对应Java中的multiThread方法
func (m *MarketNode) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 异步查询活动配置
	queryActivityTask := thread.NewQueryGroupBuyActivityDiscountVOThreadTask(
		requestParameter.Source,
		requestParameter.Channel,
		m.activityRepository,
	)

	// 异步查询商品信息
	querySkuTask := thread.NewQuerySkuVOFromDBThreadTask(
		requestParameter.GoodsId,
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

	// 写入上下文
	if activityVO != nil {
		dynamicContext.SetGroupBuyActivityDiscountVO(activityVO)
	}
	if skuVO != nil {
		dynamicContext.SetSkuVO(skuVO)
	}

	log.Printf("拼团商品查询试算服务-MarketNode异步线程加载数据完成")
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (m *MarketNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-MarketNode requestParameter:%+v", requestParameter)

	groupBuyActivityDiscountVO := dynamicContext.GetGroupBuyActivityDiscountVO()
	if groupBuyActivityDiscountVO == nil {
		return nil, fmt.Errorf("拼团活动配置为空")
	}

	groupBuyDiscount := groupBuyActivityDiscountVO.GroupBuyDiscount
	if groupBuyDiscount == nil {
		return nil, fmt.Errorf("拼团折扣配置为空")
	}

	skuVO := dynamicContext.GetSkuVO()
	if skuVO == nil {
		return nil, fmt.Errorf("商品信息为空")
	}

	discountCalculateService, exists := m.discountCalculateServiceMap[groupBuyDiscount.MarketPlan]
	if !exists {
		log.Printf("不存在%s类型的折扣计算服务，支持类型为:%v", groupBuyDiscount.MarketPlan, m.getSupportedMarketPlans())
		return nil, fmt.Errorf("不支持的折扣类型: %s", groupBuyDiscount.MarketPlan)
	}

	// 折扣价格
	originalPrice := big.NewFloat(skuVO.OriginalPrice)
	deductionPrice := discountCalculateService.Calculate(requestParameter.UserId, originalPrice, groupBuyDiscount)
	dynamicContext.SetDeductionPrice(deductionPrice)

	return m.Router(requestParameter, dynamicContext)
}

// getSupportedMarketPlans 获取支持的营销计划类型
func (m *MarketNode) getSupportedMarketPlans() []model.MarketPlanEnum {
	plans := make([]model.MarketPlanEnum, 0, len(m.discountCalculateServiceMap))
	for plan := range m.discountCalculateServiceMap {
		plans = append(plans, plan)
	}
	return plans
}

// Get 获取下一个策略处理器
// 营销节点处理完成后进入结束节点
func (m *MarketNode) Get(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	log.Printf("营销节点处理完成，进入结束节点")

	// 返回结束节点作为下一个处理器
	return m.endNode, nil
}

// 确保 MarketNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*MarketNode)(nil)
