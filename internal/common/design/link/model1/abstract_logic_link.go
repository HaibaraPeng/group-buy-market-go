package model1

import "context"

// AbstractLogicLink 实现了ILogicLink接口的抽象类
type AbstractLogicLink[T any, D any, R any] struct {
	next        ILogicLink[T, D, R]
	doApplyFunc func(ctx context.Context, requestParameter T, dynamicContext D) (R, error)
}

// Next 返回下一个节点
func (l *AbstractLogicLink[T, D, R]) Next() ILogicLink[T, D, R] {
	return l.next
}

// AppendNext 添加下一个节点
func (l *AbstractLogicLink[T, D, R]) AppendNext(next ILogicLink[T, D, R]) ILogicLink[T, D, R] {
	l.next = next
	return next
}

// NextLink 调用下一个节点的Apply方法
func (l *AbstractLogicLink[T, D, R]) NextLink(requestParameter T, dynamicContext D) (R, error) {
	var result R
	if l.next != nil {
		return l.next.Apply(requestParameter, dynamicContext)
	}
	return result, nil
}

// SetDoApplyFunc 设置业务流程受理函数
func (r *AbstractLogicLink[T, D, R]) SetDoApplyFunc(f func(ctx context.Context, requestParameter T, dynamicContext D) (R, error)) {
	r.doApplyFunc = f
}

// DoApply 业务流程受理 - 可以被子类重写或使用设置的函数
func (r *AbstractLogicLink[T, D, R]) Apply(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	if r.doApplyFunc != nil {
		return r.doApplyFunc(ctx, requestParameter, dynamicContext)
	}
	// 默认实现
	var zero R
	return zero, nil
}
