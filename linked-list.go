package main

import "fmt"

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

func (l *LinkedList) Delete(x int) error {
	node := l.Head
	for {

		if x == l.Head.Data {
			l.Head = l.Head.Next
			l.Head.Prev = nil
			return nil
		}

		if x == l.Tail.Data {
			l.Tail = l.Tail.Prev
			l.Tail.Next = nil
			return nil
		}

		if node.Data == x {
			(node.Prev).Next = node.Next
			(node.Next).Prev = node.Prev
			return nil
		}

		node = node.Next
		if node == nil {
			return fmt.Errorf("такого значения нет")
		}
	}

}
