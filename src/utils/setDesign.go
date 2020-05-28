package utils

import (
	"sort"
	"sync"
)

// 实现 set 集合，变相实现 切片去重
// by phpgo.cnblogs.com
type IntSet struct {
	m map[int]bool
	sync.RWMutex
}

func NewIntSet() *IntSet {
	return &IntSet{
		m: map[int]bool{},
	}
}

func (s *IntSet) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *IntSet) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *IntSet) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *IntSet) Len() int {
	return len(s.List())
}

func (s *IntSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *IntSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *IntSet) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *IntSet) SortList() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// 实现 set 集合，变相实现 切片去重
// by 52php.cnblogs.com
type StringSet struct {
	m map[string]bool
	sync.RWMutex
}

func NewStringSet() *StringSet {
	return &StringSet{
		m: map[string]bool{},
	}
}

func (s *StringSet) Add(items ...string) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		s.m[item] = true
	}
}

func (s *StringSet) Remove(items ...string) {
	s.Lock()
	defer s.Unlock()
	if len(items) == 0 {
		return
	}
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *StringSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *StringSet) Len() int {
	return len(s.List())
}

func (s *StringSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *StringSet) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *StringSet) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *StringSet) SortList() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	sort.Strings(list)
	return list
}
