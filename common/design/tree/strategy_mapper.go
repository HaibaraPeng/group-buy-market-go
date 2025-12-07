package tree

// StrategyMapper 策略映射器接口
// T 入参类型
// D 上下文参数
// R 返参类型
type StrategyMapper[T any, D any, R any] interface {
	// Get 获取待执行策略
	Get(requestParameter T, dynamicContext D) (StrategyHandler[T, D, R], error)
}
