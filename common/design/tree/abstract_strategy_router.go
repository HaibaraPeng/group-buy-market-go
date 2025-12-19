package tree

import "context"

// AbstractStrategyRouter 抽象策略路由器
type AbstractStrategyRouter[T any, D any, R any] struct {
	defaultStrategyHandler StrategyHandler[T, D, R]
}

// SetDefaultStrategyHandler 设置默认策略处理器
func (r *AbstractStrategyRouter[T, D, R]) SetDefaultStrategyHandler(handler StrategyHandler[T, D, R]) {
	r.defaultStrategyHandler = handler
}

// GetDefaultStrategyHandler 获取默认策略处理器
func (r *AbstractStrategyRouter[T, D, R]) GetDefaultStrategyHandler() StrategyHandler[T, D, R] {
	if r.defaultStrategyHandler == nil {
		return NewDefaultStrategyHandler[T, D, R]()
	}
	return r.defaultStrategyHandler
}

// Router 路由策略
func (r *AbstractStrategyRouter[T, D, R]) Router(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	strategyHandler, err := r.Get(requestParameter, dynamicContext)
	if err != nil {
		return r.GetDefaultStrategyHandler().Apply(ctx, requestParameter, dynamicContext)
	}

	if strategyHandler != nil {
		return strategyHandler.Apply(ctx, requestParameter, dynamicContext)
	}

	return r.GetDefaultStrategyHandler().Apply(ctx, requestParameter, dynamicContext)
}

// Get 获取待执行策略 - 需要子类实现
func (r *AbstractStrategyRouter[T, D, R]) Get(requestParameter T, dynamicContext D) (StrategyHandler[T, D, R], error) {
	// 子类需要实现此方法
	var handler StrategyHandler[T, D, R]
	return handler, nil
}
