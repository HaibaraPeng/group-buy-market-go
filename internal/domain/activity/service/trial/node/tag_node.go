package node

import (
	"context"
	"group-buy-market-go/common/design/tree"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"group-buy-market-go/internal/infrastructure/adapter/repository"

	"github.com/go-kratos/kratos/v2/log"
)

// TagNode 人群标签判断节点
// 负责判断用户是否在指定的人群标签范围内
type TagNode struct {
	core.AbstractGroupBuyMarketSupport
	activityRepository *repository.ActivityRepository
	endNode            *EndNode
	log                *log.Helper
}

// NewTagNode 创建人群标签节点
func NewTagNode(activityRepository *repository.ActivityRepository, endNode *EndNode, logger log.Logger) *TagNode {
	tagNode := &TagNode{
		activityRepository: activityRepository,
		endNode:            endNode,
		log:                log.NewHelper(logger),
	}

	// 设置自定义方法实现
	tagNode.SetDoApplyFunc(tagNode.doApply)
	tagNode.SetMultiThreadFunc(tagNode.multiThread)
	tagNode.SetDoGet(tagNode.Get)

	return tagNode
}

// multiThread 异步加载数据 - 标签节点不需要异步加载
func (t *TagNode) multiThread(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 标签节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法，判断用户是否在人群标签范围内
func (t *TagNode) doApply(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	t.log.Infow("拼团商品查询试算服务-TagNode", "userId", requestParameter.UserId, "requestParameter", requestParameter)

	// 获取拼团活动配置
	groupBuyActivityDiscountVO := dynamicContext.GetGroupBuyActivityDiscountVO()

	tagId := groupBuyActivityDiscountVO.TagId
	visible := groupBuyActivityDiscountVO.IsVisible()
	enable := groupBuyActivityDiscountVO.IsEnable()

	// 人群标签配置为空，则走默认值
	if tagId == "" {
		dynamicContext.SetVisible(true)
		dynamicContext.SetEnable(true)
		return t.Router(ctx, requestParameter, dynamicContext)
	}

	// 是否在人群范围内；visible、enable 如果值为 true 则表示没有配置拼团限制，那么就直接保证为 true 即可
	isWithin, err := t.activityRepository.IsTagCrowdRange(ctx, tagId, requestParameter.UserId)
	if err != nil {
		t.log.Errorf("检查用户是否在人群标签范围内失败: %v", err)
		// 出错时默认允许用户访问
		isWithin = true
	}

	dynamicContext.SetVisible(visible || isWithin)
	dynamicContext.SetEnable(enable || isWithin)

	return t.Router(ctx, requestParameter, dynamicContext)
}

// Get 获取下一个策略处理器
func (t *TagNode) Get(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	return t.endNode, nil
}

// 确保 TagNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*TagNode)(nil)
