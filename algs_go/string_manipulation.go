package main

import "fmt"

func alternatingCharacters(s string) int {
	As := 0
	Bs := 0
	for _, value := range s {
		if string(value) == "A" {
			As++
		} else if string(value) == "B" {
			Bs++
		}
	}
	return 1
}

func main() {
	s := "AABBAA"
	result := alternatingCharacters(s)
	fmt.Println(result)
}
