package set

import "sync"

type AbstractSet interface {
	Add(int)
	Remove(int)
	Has(int) bool
	Len() int
	IsEmpty() bool
	Clear()
	ToList() []int
}

type Set struct {
	m      map[int]interface{}
	length int
	lock   sync.RWMutex
}

func NewSet(cap int64) *Set {
	m := make(map[int]interface{}, cap)
	return &Set{m: m}
}

func (s *Set) Len() int {
	return s.length
}

func (s *Set) Add(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.m[item] = struct{}{}
	s.length = len(s.m)
}

func (s *Set) Remove(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.length == 0 {
		return
	}

	delete(s.m, item)
	s.length = len(s.m)
}

func (s *Set) Has(item int) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.m[item]
	return ok
}

func (s *Set) IsEmpty() bool {
	return s.length == 0
}

func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.m = make(map[int]interface{}, 0)
	s.length = 0
}

func (s *Set) ToList() []int {
	s.lock.Lock()
	defer s.lock.Unlock()

	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
