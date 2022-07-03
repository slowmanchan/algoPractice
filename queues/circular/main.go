package main

import "fmt"

type CircularQueue struct {
	Data []int
	Head int
	Tail int
	Size int
}

func NewCircularQueue(k int) *CircularQueue {
	q := new(CircularQueue)

	q.Data = make([]int, k)
	q.Head = -1
	q.Tail = -1
	q.Size = k

	return q
}

func (q *CircularQueue) EnQueue(val int) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.Head = 0
	}

	q.Tail = (q.Tail + 1) % q.Size
	q.Data[q.Tail] = val
	return true
}

func (q *CircularQueue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.Head == q.Tail {
		q.Head = -1
		q.Tail = -1
		return true
	}

	q.Head = (q.Head + 1) % q.Size
	return true
}

func (q *CircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Data[q.Head]
}

func (q *CircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Data[q.Tail]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.Head == -1
}

func (q *CircularQueue) IsFull() bool {
	return ((q.Tail + 1) % q.Size) == q.Head
}

func main() {
	q := NewCircularQueue(2)
	q.EnQueue(1)
	q.EnQueue(2)

	q.Dequeue()
	q.Dequeue()

	q.EnQueue(1)
	q.EnQueue(1)
	fmt.Printf("%+v\n", q)
}
