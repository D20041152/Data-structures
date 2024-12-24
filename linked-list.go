package main

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(x int) {
	node := &Node{Data: x}
	if l.Head == nil {
		l.Head = node
		return
	}

	if l.Tail == nil {
		l.Tail = node
		l.Tail.Prev = l.Head
		l.Head.Next = l.Tail
		return
	}

	l.Tail.Next = node
	temp := l.Tail
	l.Tail = node
	l.Tail.Prev = temp

}
