package main

import "fmt"

func davis(n int, memoTable map[int]int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	} else {
		
		return davis(n-3) + davis(n-2) + davis(n-1)
	}

}
func main() {
	n := 4
	memoTable := make(map[int]int)
	result := davis(n, memoTable)
	fmt.Println(result)
}
