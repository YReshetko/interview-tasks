package linked_list

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewList[T any](arr []T) Node[T] {
	first := Node[T]{
		value: arr[0],
	}
	last := &first
	for i, t := range arr {
		if i == 0 {
			continue
		}
		l := Node[T]{
			value: t,
		}
		last.next = &l
		last = &l
	}
	return first
}

func (n Node[T]) forEach(fn func(Node[T])) {
	fn(n)
	if n.next != nil {
		n.next.forEach(fn)
	}
}

func (n Node[T]) forEachDown(fn func(Node[T])) {
	if n.next != nil {
		n.next.forEachDown(fn)
	}
	fn(n)
}

func (n Node[T]) ToArray() []T {
	a := []T{}
	n.forEach(func(n Node[T]) {
		a = append(a, n.value)
	})
	return a
}

func Middle[T any](node Node[T]) T {
	middle := node
	ind := 0
	node.forEach(func(Node[T]) {
		ind++
		if ind%2 == 0 {
			middle = *middle.next
		}
	})
	return middle.value
}

func FindFromEnd[T any](node Node[T], index int) T {
	var value T
	ind := 1
	node.forEachDown(func(n Node[T]) {
		if index == ind {
			value = n.value
		}
		ind++
	})
	return value
}

func Reverse[T any](node Node[T]) Node[T] {
	var newNode *Node[T]
	var last *Node[T]
	node.forEachDown(func(n Node[T]) {
		if newNode == nil {
			newNode = &Node[T]{
				value: n.value,
			}
			last = newNode
			return
		}
		l := &Node[T]{
			value: n.value,
		}
		last.next = l
		last = l

	})
	return *newNode
}
