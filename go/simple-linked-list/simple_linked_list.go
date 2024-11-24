package linkedlist

import "fmt"

// Define the List and Element types here.
type (
	Element struct {
		value int
		next  *Element
	}

	List struct {
		head *Element
		size int
	}
)

func New(elements []int) *List {
	if len(elements) < 1 {
		return &List{}
	}

	l := &List{nil, 0}
	for _, element := range elements {
		l.Push(element)
	}
	return l
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.head = &Element{value: element, next: l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, fmt.Errorf("empty list")
	}

	value := l.head.value
	l.head = l.head.next
	l.size--

	return value, nil
}

func (l *List) Array() []int {
	if l.size < 1 {
		return []int{}
	}
	array := make([]int, l.size)
	for i, ptr := l.size-1, l.head; ptr != nil; i, ptr = i-1, ptr.next {
		array[i] = ptr.value
	}
	return array
}

func (l *List) Reverse() *List {
	rev := &List{}
	if l.size < 1 {
		return rev
	}
	ptr := l.head
	for ptr != nil {
		rev.Push(ptr.value)
		ptr = ptr.next
	}
	return rev
}
