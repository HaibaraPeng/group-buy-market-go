package tree

import "context"

// AbstractMultiThreadStrategyRouter 异步资源加载策略路由器
type AbstractMultiThreadStrategyRouter[T any, D any, R any] struct {
	// 定义接口字段，用于调用子类方法
	multiThreadFunc func(ctx context.Context, requestParameter T, dynamicContext D) error
	doApplyFunc     func(ctx context.Context, requestParameter T, dynamicContext D) (R, error)
	doGet           func(ctx context.Context, requestParameter T, dynamicContext D) (StrategyHandler[T, D, R], error)
}

// SetMultiThreadFunc 设置异步加载数据函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) SetMultiThreadFunc(f func(ctx context.Context, requestParameter T, dynamicContext D) error) {
	r.multiThreadFunc = f
}

// SetDoApplyFunc 设置业务流程受理函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) SetDoApplyFunc(f func(ctx context.Context, requestParameter T, dynamicContext D) (R, error)) {
	r.doApplyFunc = f
}

// SetDoGet 设置获取待执行策略函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) SetDoGet(f func(ctx context.Context, requestParameter T, dynamicContext D) (StrategyHandler[T, D, R], error)) {
	r.doGet = f
}

// Apply 应用策略
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) Apply(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	// 异步加载数据
	err := r.MultiThread(ctx, requestParameter, dynamicContext)
	if err != nil {
		var zero R
		return zero, err
	}

	// 业务流程受理
	return r.DoApply(ctx, requestParameter, dynamicContext)
}

// MultiThread 异步加载数据 - 可以被子类重写或使用设置的函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) MultiThread(ctx context.Context, requestParameter T, dynamicContext D) error {
	if r.multiThreadFunc != nil {
		return r.multiThreadFunc(ctx, requestParameter, dynamicContext)
	}
	// 默认实现
	return nil
}

// DoApply 业务流程受理 - 可以被子类重写或使用设置的函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) DoApply(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	if r.doApplyFunc != nil {
		return r.doApplyFunc(ctx, requestParameter, dynamicContext)
	}
	// 默认实现
	var zero R
	return zero, nil
}

func (r *AbstractMultiThreadStrategyRouter[T, D, R]) Router(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	strategyHandler, err := r.Get(ctx, requestParameter, dynamicContext)

	if strategyHandler != nil {
		return strategyHandler.Apply(ctx, requestParameter, dynamicContext)
	}
	var zero R
	return zero, err
}

// Get 获取待执行策略 - 可以被子类重写或使用设置的函数
func (r *AbstractMultiThreadStrategyRouter[T, D, R]) Get(ctx context.Context, requestParameter T, dynamicContext D) (StrategyHandler[T, D, R], error) {
	if r.doGet != nil {
		return r.doGet(ctx, requestParameter, dynamicContext)
	}
	// 默认实现
	var handler StrategyHandler[T, D, R]
	return handler, nil
}
