package link

import "group-buy-market-go/internal/common/design/link/model1"

// ExampleHandler 是一个示例处理器，演示如何使用责任链模式
type ExampleHandler[T any, D any, R any] struct {
	*model1.AbstractLogicLink[T, D, R]
	name string
}

// NewExampleHandler 创建一个新的示例处理器
func NewExampleHandler[T any, D any, R any](name string) *ExampleHandler[T, D, R] {
	return &ExampleHandler[T, D, R]{
		AbstractLogicLink: &model1.AbstractLogicLink[T, D, R]{},
		name:              name,
	}
}

// Apply 实现了ILogicLink接口的Apply方法
func (h *ExampleHandler[T, D, R]) Apply(requestParameter T, dynamicContext D) (R, error) {
	// 这里是具体的处理逻辑
	// 示例：处理请求并返回结果
	var result R
	// 在实际实现中，这里会包含具体的业务逻辑

	// 如果有下一个节点，则传递给下一个节点处理
	if h.Next() != nil {
		return h.NextLink(requestParameter, dynamicContext)
	}

	// 如果没有下一个节点，则返回结果
	return result, nil
}
