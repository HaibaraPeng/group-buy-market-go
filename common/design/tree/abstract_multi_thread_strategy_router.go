package tree

// AbstractMultiThreadStrategyRouter 异步资源加载策略路由器
type AbstractMultiThreadStrategyRouter[T any, D any, R any] struct {
	AbstractStrategyRouter[T, D, R]
}

// Apply 应用策略
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) Apply(requestParameter T, dynamicContext D) (R, error) {
	// 异步加载数据
	err := r.MultiThread(requestParameter, dynamicContext)
	if err != nil {
		var zero R
		return zero, err
	}

	// 业务流程受理
	return r.DoApply(requestParameter, dynamicContext)
}

// MultiThread 异步加载数据 - 需要子类实现
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) MultiThread(requestParameter T, dynamicContext D) error {
	// 子类需要实现此方法
	// 这里可以设置超时时间
	return nil
}

// DoApply 业务流程受理 - 需要子类实现
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) DoApply(requestParameter T, dynamicContext D) (R, error) {
	// 子类需要实现此方法
	var zero R
	return zero, nil
}
