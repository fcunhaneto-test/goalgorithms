package lists

import (
	"errors"
	"fmt"
)

type Queue struct {
	a   []int
	ini int
	end int
}

/*
InitQueue starts the queue

param:
list: slice to start queue, if don't have a slice pass nil
length: int the queue lenght
*/
func InitQueue(list []int, length int) Queue {
	l := make([]int, length)
	n := len(list) - 1
	if list != nil {
		for i := 0; i <= n; i++ {
			l[i] = list[i]
		}
	}

	q := Queue{
		a:   l,
		ini: 0,
		end: n,
	}

	return q
}

/*
Enqueue inserting element in the queue
*/
func (q *Queue) Enqueue(num int) (Queue, error) {
	q.end++
	if q.isQueueFull() {
		return *q, errors.New("Queue is full")
	}

	q.a[q.end] = num

	return *q, nil
}

/*
Dequeue get the first element inserted in the queue
*/
func (q *Queue) Dequeue() (int, error) {
	fmt.Println("dequeue =", q.ini, q.end)
	if q.isQueueEmpty() {
		return 0, errors.New("Queue is empty")
	}

	q.ini++
	return q.a[q.ini-1], nil
}

func (q *Queue) isQueueFull() bool {
	if q.end < len(q.a) {
		return false
	}

	return true
}

func (q *Queue) isQueueEmpty() bool {
	fmt.Println("isQueueEmpty =", q.ini, q.end)
	if q.ini <= q.end {
		return false
	}

	return true
}
