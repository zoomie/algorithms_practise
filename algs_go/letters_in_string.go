package main

import "fmt"

func countChars(arr []string) map[string]int {
	result := make(map[string]int)
	for _, char := range arr {
		_, ok := result[char]
		if ok {
			result[char]++
		} else {
			result[char] = 1
		}
	}
	return result
}

func main() {
	arr1 := []string{"a", "b", "c", "a"}
	arr2 := []string{"a", "b", "c", "a"}
	result1 := countChars(arr1)
	result2 := countChars(arr2)
	for key, value := range result1 {
		if 
	}
}
