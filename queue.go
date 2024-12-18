package main

import "fmt"

type Queue struct {
	mas  []int
	head int
	tail int
}

func NewQueue(N int) *Queue {
	return &Queue{mas: make([]int, N), head: 0, tail: 0}
}

func (q *Queue) Enqueue(a int) error {
	if q.tail == len(q.mas) {
		if q.tail % len(q.mas) == q.head {
			return fmt.Errorf("очередь переполнена")
		} else {
			q.tail = q.tail % len(q.mas)
		}

	}

	q.mas[q.tail] = a
	q.tail = q.tail + 1
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.tail == q.head {
		return -100, fmt.Errorf("очередь опусташена")
	}

	if q.head == len(q.mas) {
		q.head = q.head % len(q.mas)
	}
	res := q.head
	q.head = q.head + 1
	return res, nil
}
