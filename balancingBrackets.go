package main

import "fmt"

type amounts int

// func (a *amounts) change() {
// 	for i := 0; i < 6; i++ {
// 		if i < len(*a) {
// 			(*a)[i] = i
// 		} else {
// 			*a = append(*a, i)
// 		}
// 	}
// }

func main() {
	x := 1
	y := 2
	x, y = y, x
	fmt.Println(x, y)
}

// func isBalanced(s string) bool {
// 	openCurly := "{"
// 	closedCurly := "}"
// 	counter := make(map[string]int)
// 	for _, codePoint := range s {
// 		char := string(codePoint)
// 		if char == openCurly {
// 			counter[char]++

// 		}
// 		if char == closedCurly {
// 			counter[char]--

// 		}
// 	}
// 	if counter[openCurly] == counter[closedCurly] {
// 		return true
// 	} else {
// 		return false
// 	}
// }
