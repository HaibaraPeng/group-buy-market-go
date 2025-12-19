package tree

import "context"

// StrategyHandler 策略处理器接口
// T 入参类型
// D 上下文参数
// R 返参类型
type StrategyHandler[T any, D any, R any] interface {
	Apply(ctx context.Context, requestParameter T, dynamicContext D) (R, error)
}
