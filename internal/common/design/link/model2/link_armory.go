package model2

// LinkArmory 是一个用于装配责任链的工具类，对应Java版本的LinkArmory类
type LinkArmory[T any, D any, R any] struct {
	logicLink *BusinessLinkedList[T, D, R]
}

// NewLinkArmory 创建一个新的LinkArmory实例，接受链名称和多个处理器
func NewLinkArmory[T any, D any, R any](linkName string, logicHandlers ...ILogicHandler[T, D, R]) *LinkArmory[T, D, R] {
	logicLink := NewBusinessLinkedList[T, D, R](linkName)

	for _, handler := range logicHandlers {
		logicLink.Add(handler)
	}

	return &LinkArmory[T, D, R]{
		logicLink: logicLink,
	}
}

// GetLogicLink 返回装配好的责任链
func (la *LinkArmory[T, D, R]) GetLogicLink() *BusinessLinkedList[T, D, R] {
	return la.logicLink
}
