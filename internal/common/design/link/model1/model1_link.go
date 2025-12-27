package model1

// ILogicLink 定义了责任链节点的接口
type ILogicLink[T any, D any, R any] interface {
	ILogicChainArmory[T, D, R]
	Apply(T, D) (R, error)
}

// ILogicChainArmory 定义了责任链装配的接口
type ILogicChainArmory[T any, D any, R any] interface {
	Next() ILogicLink[T, D, R]
	AppendNext(ILogicLink[T, D, R]) ILogicLink[T, D, R]
}
