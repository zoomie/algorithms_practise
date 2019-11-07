package main

import "fmt"

type Set struct {
	nums []int
}

func (s *Set) add(i int) {
	s.nums = append(s.nums, i)
}

func (s *Set) pop() int {
	result := s.nums[0]
	s.nums = s.nums[1:]
	return result
}

func main() {
	set := Set{}
	set.add(10)
	set.add(2)
	fmt.Println(set)
	result := set.pop()
	fmt.Println(result)
	fmt.Println(set)
}
