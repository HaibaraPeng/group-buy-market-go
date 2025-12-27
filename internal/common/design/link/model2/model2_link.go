package model2

// ILink 定义了链表接口，对应Java版本的ILink接口
type ILink[T any] interface {
	Add(T) bool
	AddFirst(T) bool
	AddLast(T) bool
	Remove(T) bool
	Get(int) T
	PrintLinkList()
}

// ILogicHandler 定义了责任链节点的接口
type ILogicHandler[T any, D any, R any] interface {
	Next(T, D) (R, error)
	Apply(T, D) (R, error)
}
