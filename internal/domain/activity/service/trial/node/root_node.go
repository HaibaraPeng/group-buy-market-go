package node

import (
	"errors"
	"group-buy-market-go/internal/domain/activity/model"
	"group-buy-market-go/internal/domain/activity/service/trial/core"
	"log"
)

// RootNode 根节点
// 策略树的起始节点，负责初始化处理流程
type RootNode struct {
	core.AbstractGroupBuyMarketSupport
}

// NewRootNode 创建根节点
func NewRootNode() *RootNode {
	return &RootNode{}
}

// doApply 业务流程受理
// 对应Java中的doApply方法
func (r *RootNode) doApply(requestParameter *model.MarketProductEntity, dynamicContext *core.DynamicContext) (*model.TrialBalanceEntity, error) {
	log.Printf("拼团商品查询试算服务-RootNode userId:%d requestParameter:%+v", dynamicContext.UserID, requestParameter)

	// 参数判断
	if requestParameter == nil || dynamicContext == nil {
		return nil, errors.New("参数不能为空")
	}

	// 注意：Go版本的实体结构与Java版本略有不同，这里根据Go的实际情况进行校验
	if dynamicContext.UserID <= 0 || requestParameter.ID <= 0 {
		return nil, errors.New("非法参数: 用户ID和商品ID不能为空")
	}

	return r.Router(requestParameter, dynamicContext)
}

// 确保 RootNode 实现了 StrategyHandler 接口
var _ core.StrategyHandler = (*RootNode)(nil)
