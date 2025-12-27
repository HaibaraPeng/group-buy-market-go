package model2

import "fmt"

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

// Node 表示链表中的节点
type Node[T any] struct {
	Item T
	Next *Node[T]
	Prev *Node[T]
}

// NewNode 创建一个新的节点
func NewNode[T any](prev *Node[T], item T, next *Node[T]) *Node[T] {
	node := &Node[T]{
		Item: item,
		Next: next,
		Prev: prev,
	}
	return node
}

// LinkedList 是一个双向链表实现，对应Java版本的LinkedList类
type LinkedList[T any] struct {
	Name  string
	Size  int
	First *Node[T]
	Last  *Node[T]
}

// NewLinkedList 创建一个新的链表实例
func NewLinkedList[T any](name string) *LinkedList[T] {
	return &LinkedList[T]{
		Name:  name,
		Size:  0,
		First: nil,
		Last:  nil,
	}
}

// linkFirst 将元素添加到链表头部
func (l *LinkedList[T]) linkFirst(item T) {
	f := l.First
	newNode := NewNode[T](nil, item, f)
	l.First = newNode
	if f == nil {
		l.Last = newNode
	} else {
		f.Prev = newNode
	}
	l.Size++
}

// linkLast 将元素添加到链表尾部
func (l *LinkedList[T]) linkLast(item T) {
	lNode := l.Last
	newNode := NewNode[T](lNode, item, nil)
	l.Last = newNode
	if lNode == nil {
		l.First = newNode
	} else {
		lNode.Next = newNode
	}
	l.Size++
}

// Add 将元素添加到链表尾部
func (l *LinkedList[T]) Add(item T) bool {
	l.linkLast(item)
	return true
}

// AddFirst 将元素添加到链表头部
func (l *LinkedList[T]) AddFirst(item T) bool {
	l.linkFirst(item)
	return true
}

// AddLast 将元素添加到链表尾部
func (l *LinkedList[T]) AddLast(item T) bool {
	l.linkLast(item)
	return true
}

// Remove 从链表中移除指定元素
func (l *LinkedList[T]) Remove(item T) bool {
	node := l.First
	for node != nil {
		// 比较元素是否相等
		if isEqual(node.Item, item) {
			l.unlink(node)
			return true
		}
		node = node.Next
	}
	return false
}

// isEqual 比较两个任意类型的值是否相等
func isEqual[T any](a, b T) bool {
	// 使用反射进行比较，对于基本类型直接比较
	return any(a) == any(b)
}

// unlink 从链表中移除指定节点
func (l *LinkedList[T]) unlink(node *Node[T]) T {
	element := node.Item
	next := node.Next
	prev := node.Prev

	if prev == nil {
		l.First = next
	} else {
		prev.Next = next
		node.Prev = nil
	}

	if next == nil {
		l.Last = prev
	} else {
		next.Prev = prev
		node.Next = nil
	}

	node.Item = *new(T) // 将节点的Item设置为零值
	l.Size--
	return element
}

// Get 获取指定索引处的元素
func (l *LinkedList[T]) Get(index int) T {
	node := l.node(index)
	return node.Item
}

// node 获取指定索引处的节点
func (l *LinkedList[T]) node(index int) *Node[T] {
	if index < (l.Size >> 1) { // 如果索引在前半部分，从头开始遍历
		x := l.First
		for i := 0; i < index; i++ {
			x = x.Next
		}
		return x
	} else { // 如果索引在后半部分，从尾开始遍历
		x := l.Last
		for i := l.Size - 1; i > index; i-- {
			x = x.Prev
		}
		return x
	}
}

// PrintLinkList 打印链表内容
func (l *LinkedList[T]) PrintLinkList() {
	if l.Size == 0 {
		fmt.Println("链表为空")
	} else {
		temp := l.First
		fmt.Printf("目前的列表，头节点：%v 尾节点：%v 整体：", l.First.Item, l.Last.Item)
		for temp != nil {
			fmt.Printf("%v，", temp.Item)
			temp = temp.Next
		}
		fmt.Println()
	}
}

// GetName 获取链表名称
func (l *LinkedList[T]) GetName() string {
	return l.Name
}
