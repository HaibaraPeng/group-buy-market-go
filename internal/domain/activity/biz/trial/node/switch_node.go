package node

import (
	"context"
	"fmt"
	"group-buy-market-go/internal/common/design/tree"
	"group-buy-market-go/internal/domain/activity/biz/trial/core"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/infrastructure/adapter/repository"

	"github.com/go-kratos/kratos/v2/log"
)

// SwitchNode 开关节点
// 用于判断营销活动是否开启
type SwitchNode struct {
	core.AbstractGroupBuyMarketSupport
	activityRepository *repository.ActivityRepository
	marketNode         *MarketNode
	log                *log.Helper
}

// NewSwitchNode 创建开关节点
func NewSwitchNode(activityRepository *repository.ActivityRepository, marketNode *MarketNode, logger log.Logger) *SwitchNode {
	switchNode := &SwitchNode{
		activityRepository: activityRepository,
		marketNode:         marketNode,
		log:                log.NewHelper(logger),
	}

	// 设置自定义方法实现
	switchNode.SetDoApplyFunc(switchNode.doApply)
	switchNode.SetMultiThreadFunc(switchNode.multiThread)
	switchNode.SetDoGet(switchNode.Get)

	return switchNode
}

// multiThread 异步加载数据 - 开关节点不需要异步加载
func (r *SwitchNode) multiThread(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) error {
	// 开关节点不需要异步加载数据
	return nil
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *SwitchNode) doApply(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	r.log.Infow("拼团商品查询试算服务-SwitchNode", "userId", requestParameter.UserId, "requestParameter", requestParameter)

	// 根据用户ID切量
	userId := requestParameter.UserId

	// 判断是否降级
	if r.activityRepository.DowngradeSwitch() {
		r.log.Infof("拼团活动降级拦截 %s", userId)
		return nil, fmt.Errorf("拼团活动降级拦截: %s", userId)
	}

	// 切量范围判断
	inCutRange, err := r.activityRepository.CutRange(userId)
	if err != nil {
		r.log.Errorf("判断切量范围时出错: %v", err)
		// 出错时默认继续处理
		inCutRange = true
	}

	if !inCutRange {
		r.log.Infof("拼团活动切量拦截 %s", userId)
		return nil, fmt.Errorf("拼团活动切量拦截: %s", userId)
	}

	return r.Router(ctx, requestParameter, dynamicContext)
}

// Get 获取下一个策略处理器
func (r *SwitchNode) Get(ctx context.Context, requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity], error) {
	r.log.Info("开关节点处理完成，进入营销节点")

	// 返回营销节点作为下一个处理器
	return r.marketNode, nil
}

// 确保 SwitchNode 实现了 StrategyHandler 接口
var _ tree.StrategyHandler[*model.MarketProductEntity, *core.DynamicContext, *model.TrialBalanceEntity] = (*SwitchNode)(nil)
