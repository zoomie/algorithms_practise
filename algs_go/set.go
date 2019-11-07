package main

import "fmt"

type Set struct {
	data map[int]bool
}

func (s *Set) add(i int) {
	if s.data == nil {
		s.data = make(map[int]bool)
	}
	_, ok := s.data[i]
	if !ok {
		s.data[i] = true
	}
}

func (s *Set) remove(i int) {
	_, ok := s.data[i]
	if ok {
		delete(s.data, i)
	}
}

func (s *Set) allValues() {
	var values []int
	for key, _ := range s.data {
		values = append(values, key)
	}
	fmt.Println(values)
}

func main() {
	s := Set{}
	s.add(10)
	s.add(10)
	s.add(20)
	s.allValues()
}
