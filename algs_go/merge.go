package main

import "fmt"

func gen(nums ...int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, num := range nums {
			c <- num
		}
	}()
	return c
}
func merge(chans ...<-chan int) <-chan int {
	output := make(chan int)
	go func() {
		// defer close(output)
		for _, c := range chans {
			go func(c <-chan int) {
				for val := range c {
					output <- val
				}
			}(c)
		}
	}()
	return output
}
func main() {
	a := gen(6, 7, 8, 9)
	b := gen(5, 6, 6, 7)
	c := merge(a, b)
	for i := range c {
		fmt.Println(i)
	}
}
