package core

import (
	"group-buy-market-go/internal/domain/activity/model"
)

// AbstractGroupBuyMarketSupport 是拼团营销支撑的抽象基类
// 它提供了所有策略节点的公共功能和基础实现
type AbstractGroupBuyMarketSupport struct {
	defaultStrategyHandler StrategyHandler
}

// SetDefaultStrategyHandler 设置默认策略处理器
func (s *AbstractGroupBuyMarketSupport) SetDefaultStrategyHandler(handler StrategyHandler) {
	s.defaultStrategyHandler = handler
}

// GetDefaultStrategyHandler 获取默认策略处理器
func (s *AbstractGroupBuyMarketSupport) GetDefaultStrategyHandler() StrategyHandler {
	if s.defaultStrategyHandler == nil {
		return nil // 或者返回一个默认处理器
	}
	return s.defaultStrategyHandler
}

// Router 路由策略
func (s *AbstractGroupBuyMarketSupport) Router(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	strategyHandler, err := s.Get(requestParameter, dynamicContext)
	if err != nil || strategyHandler == nil {
		if s.defaultStrategyHandler != nil {
			return s.defaultStrategyHandler.Apply(requestParameter, dynamicContext)
		}
		// 如果没有默认处理器，返回空结果
		return &model.TrialBalanceEntity{}, nil
	}

	return strategyHandler.Apply(requestParameter, dynamicContext)
}

// Apply 应用策略
func (s *AbstractGroupBuyMarketSupport) Apply(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	// 异步加载数据
	multiThreadErr := s.multiThread(requestParameter, dynamicContext)
	if multiThreadErr != nil {
		return &model.TrialBalanceEntity{}, multiThreadErr
	}

	// 业务流程受理
	return s.doApply(requestParameter, dynamicContext)
}

// multiThread 异步加载数据 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) multiThread(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) error {
	// 子类需要实现此方法
	return nil
}

// doApply 业务流程受理 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) doApply(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (*model.TrialBalanceEntity, error) {
	// 子类需要实现此方法
	return &model.TrialBalanceEntity{}, nil
}

// Get 获取待执行策略 - 需要子类实现
func (s *AbstractGroupBuyMarketSupport) Get(requestParameter *model.MarketProductEntity, dynamicContext *DynamicContext) (StrategyHandler, error) {
	// 子类需要实现此方法
	return nil, nil
}
