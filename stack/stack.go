package stack

import "sync"

type AbstractStack interface {
	Push(string)
	Pop() string
	Peek() string
	IsEmpty() bool
	Size() int
}

type Stack struct {
	slice []string
	size  int
	lock  sync.Mutex
}

func (s *Stack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.slice = append(s.slice, v)
	s.size++
}

func (s *Stack) Pop() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		panic("empty")
	}

	v := s.slice[s.size-1]
	s.slice = s.slice[0 : s.size-1]
	s.size--
	return v
}

func (s *Stack) Peek() string {
	if s.size == 0 {
		panic("empty")
	}

	return s.slice[s.size-1]
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
