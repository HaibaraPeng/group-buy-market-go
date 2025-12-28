package model2

import "context"

// BusinessLinkedList 是一个用于处理责任链逻辑的链表实现，对应Java版本的BusinessLinkedList类
type BusinessLinkedList[T any, D any, R any] struct {
	*LinkedList[ILogicHandler[T, D, R]] // 继承自LinkedList，存储ILogicHandler类型的元素
}

// NewBusinessLinkedList 创建一个新的BusinessLinkedList实例
func NewBusinessLinkedList[T any, D any, R any](name string) *BusinessLinkedList[T, D, R] {
	return &BusinessLinkedList[T, D, R]{
		LinkedList: NewLinkedList[ILogicHandler[T, D, R]](name),
	}
}

// Apply 按顺序执行链表中的每个处理器，直到返回非空结果
func (bll *BusinessLinkedList[T, D, R]) Apply(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	current := bll.First

	for current != nil {
		handler := current.Item
		result, err := handler.Apply(ctx, requestParameter, dynamicContext)
		if err != nil {
			var zero R
			return zero, err
		}
		// 检查结果是否为零值，如果是则继续执行下一个处理器
		if !isZeroValue(result) {
			return result, nil
		}
		current = current.Next
	}

	var zero R
	return zero, nil
}

// isZeroValue 检查一个值是否为零值
func isZeroValue[T any](v T) bool {
	var zero T
	return any(zero) == any(v)
}

// Next 是另一种执行方式，与Apply类似但调用Next方法
func (bll *BusinessLinkedList[T, D, R]) Next(ctx context.Context, requestParameter T, dynamicContext D) (R, error) {
	current := bll.First

	for current != nil {
		handler := current.Item
		result, err := handler.Next(ctx, requestParameter, dynamicContext)
		if err != nil {
			var zero R
			return zero, err
		}
		// 检查结果是否为零值，如果是则继续执行下一个处理器
		if !isZeroValue(result) {
			return result, nil
		}
		current = current.Next
	}

	var zero R
	return zero, nil
}
