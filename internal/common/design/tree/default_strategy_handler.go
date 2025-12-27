package tree

import "context"

// DefaultStrategyHandler 默认策略处理器
type DefaultStrategyHandler[T any, D any, R any] struct{}

// Apply 应用默认策略
func (d *DefaultStrategyHandler[T, D, R]) Apply(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	var r R
	return r, nil
}

// NewDefaultStrategyHandler 创建默认策略处理器
func NewDefaultStrategyHandler[T any, D any, R any]() StrategyHandler[T, D, R] {
	return &DefaultStrategyHandler[T, D, R]{}
}
