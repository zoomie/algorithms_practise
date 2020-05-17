package main

import "fmt"

func testFunc(a int) {
	fmt.Println(a)
}

func shuffle(a []int) []int {
	lastIndex := len(a) - 1
	end := a[:lastIndex]
	start := []int{a[lastIndex]}
	result := append(start, end...)
	return result
}

func rotLeft(a []int, d int) []int {
	for i := 0; i < d; i++ {
		a = shuffle(a)
	}
	return a
}

func main_old() {
	var a = []int{1, 2, 3, 4, 5}
	var d int = 3
	a = rotLeft(a, d)
	fmt.Println(a)
}
