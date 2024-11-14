package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}
type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, v := range elements {
		list.Push(v)
	}

	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	if l.head == nil || l.tail == nil {
		node := &Node{Value: v}
		l.head, l.tail = node, node
		return
	}
	node := &Node{Value: v}
	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *List) Push(v interface{}) {
	if l.head == nil || l.tail == nil {
		node := &Node{Value: v}
		l.head, l.tail = node, node
		return
	}
	node := &Node{Value: v}
	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil || l.tail == nil {
		return nil, errors.New("Empty list")
	}

	node := l.head
	if node == l.tail {
		l.head, l.tail = nil, nil
		return node.Value, nil
	}
	l.head = node.next
	l.head.prev = nil

	return node.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.head == nil || l.tail == nil {
		return nil, errors.New("Empty list")
	}
	node := l.tail
	if node == l.head {
		l.head, l.tail = nil, nil
		return node.Value, nil
	}
	l.tail = node.prev
	l.tail.next = nil

	return node.Value, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.tail == nil || l.head == l.tail {
		return
	}

	l.head, l.tail = l.tail, l.head
	node := l.head
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.next
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
