package queue

import "sync"

type AbstractQueue interface {
	EnQueue(string)
	DeQueue() string
	Front() string
	Size() int
	IsEmpty() bool
}

type Queue struct {
	slice []string
	size  int
	lock  sync.Mutex
}

func (q *Queue) EnQueue(s string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.slice = append(q.slice, s)
	q.size++
}

func (q *Queue) DeQueue() string {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		panic("empty")
	}

	v := q.slice[0]
	q.slice = q.slice[1:]
	q.size--
	return v
}

func (q *Queue) Front() string {
	return q.slice[0]
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
