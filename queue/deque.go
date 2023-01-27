package queue

import "sync"

type AbstractDeque interface {
	AddFront(string)
	AddRear(string)
	RemoveFront() string
	RemoveRear() string
	Size() int
	IsEmpty() bool
}

type Deque struct {
	slice []string
	size  int
	lock  sync.Mutex
}

func (d *Deque) AddFront(s string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.slice = append([]string{s}, d.slice...)
	d.size++
}

func (d *Deque) AddRear(s string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.slice = append(d.slice, s)
	d.size++
}

func (d *Deque) RemoveFront() string {
	d.lock.Lock()
	defer d.lock.Unlock()

	v := d.slice[0]
	d.slice = d.slice[1:]
	d.size--
	return v
}

func (d *Deque) RemoveRear() string {
	d.lock.Lock()
	defer d.lock.Unlock()

	v := d.slice[d.size-1]
	d.slice = d.slice[0 : d.size-1]
	d.size--
	return v
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) Size() int {
	return d.size
}
