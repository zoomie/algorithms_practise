package main

import "fmt"

type myStack []int

func (s *myStack) add(num int) {
	*s = append(*s, num)
	fmt.Println(s)
}

func (s *myStack) pop() int {
	endIndex := len(*s) - 1
	result := (*s)[endIndex]
	*s = (*s)[:endIndex]
	return result
}
func main() {
	// stack := make(myStack, 10)
	var stack myStack
	stack.add(2)
	stack.add(3)
	result := stack.pop()
	fmt.Println(result)
	fmt.Println(stack)
}
