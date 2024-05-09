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

func NewEmpty[T GenericItem]() *List[T] {
	return &List[T]{}
}

func NewWithItems[T GenericItem](items ...T) *List[T] {
	list := &List[T]{}

	for _, v := range items {
		list.Append(v)
	}

	return list
}

func (l *List[T]) Append(item T) *List[T] {
	if l.size == 0 {
		l.first = createNode(item, nil)
		l.size++
	} else {
		n := l.first
		for range l.size {
			if n.next == nil {
				n.next = createNode(item, nil)
				l.size++
			} else {
				n = n.next
			}
		}
	}

	return l
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

func (l *List[T]) InsertAt(desiredIndex int, newItem T) {
	if l.Size() < desiredIndex {
		panic("linked queue does not have this index")
	}

	if desiredIndex == 0 {
		panic("linked queue index starts from one (1)")
	}

	n := l.first
	for i := 1; i < l.size; i++ {
		if i == desiredIndex {
			l.first = createNode(newItem, n.next)
			break
		}

		if i == desiredIndex-1 {
			if l.size == i+1 {
				n.next = createNode(newItem, nil)
				break
			}

			n.next = createNode(newItem, n.next)
			l.size++
			break
		}

		n = n.next
	}
}

func createNode[T GenericItem](item T, next *Node[T]) *Node[T] {
	return &Node[T]{Item: item, next: next}
}
