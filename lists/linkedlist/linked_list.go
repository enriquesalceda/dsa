package linkedlist

type GenericItem interface {
	~string | ~int | ~float64
}

type Node[T GenericItem] struct {
	Item T
	next *Node[T]
}

func (t *Node[T]) Next() *Node[T] {
	return t.next
}

type List[T GenericItem] struct {
	first *Node[T]
	size  int
}

func New[T GenericItem]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Append(item T) {
	if l.size == 0 {
		l.first = &Node[T]{Item: item, next: nil}
		l.size++
	} else {
		n := l.first
		for range l.size {
			if n.next == nil {
				n.next = &Node[T]{Item: item, next: nil}
				l.size++
			} else {
				n = n.next
			}
		}
	}
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) First() *Node[T] {
	return l.first
}

func (l *List[T]) Items() []T {
	var items []T

	n := l.first
	for range l.size {
		items = append(items, n.Item)
		n = n.next
	}

	return items
}

func (l *List[T]) IndexOf(item T) (bool, int) {
	n := l.first

	for i := 0; i < l.size; i++ {
		if n.Item == item {
			return true, i + 1
		}
		n = n.next
	}

	return false, 0
}
