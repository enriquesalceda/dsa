package linkedlist

type GenericItem interface {
	~string | ~int | ~float64
}

type Node[T GenericItem] struct {
	Item T
	next *Node[T]
}

type List[T GenericItem] struct {
	first *Node[T]
	size  int
}

func NewList[T GenericItem]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Append(node Node[T]) {
	if l.size == 0 {
		l.first = &node
		l.size++
	}
}

func (l *List[T]) Size() int {
	return l.size
}
