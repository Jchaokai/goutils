package set

import (
	"sync"
)

// 通过原生map来实现set，因为map的key唯一，但有个问题value可能占用内存
// 可以用空结构体struct{}填充value

type (
	empty struct{}

	UnSafeSet map[interface{}]empty

	SafeSet struct {
		s UnSafeSet
		sync.RWMutex
	}

	// Set
	// todo how to iterate
	Set interface {
		Add(items ...interface{}) bool
		Clear()
		IsSubOf(other *Set) bool
		Contains(item interface{}) bool
		Equal(other *Set) bool
		Size() int
	}
)

func NewUnSafeSet(items ...interface{}) Set {
	s := make(UnSafeSet)
	s.Add(items...)
	return &s
}
func NewSet(item ...interface{}) Set {
	s := &SafeSet{s: make(UnSafeSet)}
	s.Add(item...)
	return s
}

func (s *UnSafeSet) Add(items ...interface{}) bool {
	for _, item := range items {
		(*s)[item] = empty{}
	}
	return true
}

func (s *UnSafeSet) Clear() {
	*s = make(UnSafeSet)
}

func (s *UnSafeSet) IsSubOf(other *Set) bool {
	if s.Size() > (*other).Size() {
		return false
	}
	for k := range *s {
		if !(*other).Contains(k) {
			return false
		}
	}
	return true
}

func (s *UnSafeSet) Contains(item interface{}) bool {
	_, ok := (*s)[item]
	return ok
}

func (s *UnSafeSet) Equal(other *Set) bool {
	if s.Size() != (*other).Size() {
		return false
	}
	for k := range *s {
		if !(*other).Contains(k) {
			return false
		}
	}
	return true
}

func (s *UnSafeSet) Size() int {
	return len(*s)
}

func (s *SafeSet) Add(items ...interface{}) bool {
	s.Lock()
	defer s.Unlock()
	return s.s.Add(items...)
}

func (s *SafeSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.s = make(UnSafeSet)
}

func (s *SafeSet) IsSubOf(other *Set) bool {
	s.RLock()
	defer s.RUnlock()
	return s.s.IsSubOf(other)
}

func (s *SafeSet) Contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	return s.s.Contains(item)
}

func (s *SafeSet) Equal(other *Set) bool {
	s.RLock()
	defer s.RUnlock()
	return s.s.Equal(other)
}

func (s *SafeSet) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.s.Size()
}
