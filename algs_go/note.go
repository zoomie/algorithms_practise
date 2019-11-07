package main

import "fmt"

func generateHashMap(magazine []string) map[string]int {
	wordsToUse := make(map[string]int)
	for i := 0; i < len(magazine); i++ {
		char := magazine[i]
		_, ok := wordsToUse[char]
		if ok {
			wordsToUse[char]++
		} else {
			wordsToUse[char] = 1
		}
	}
	return wordsToUse
}

func checkMagazine(magazine []string, note []string) {
	answer := true
	wordsToUse := generateHashMap(magazine)
	for i := 0; i < len(note); i++ {
		char := note[i]
		num, ok := wordsToUse[char]
		if ok && (num >= 0) {
			wordsToUse[char]--
		} else {
			answer = false
		}
	}
	fmt.Println(answer)
}

func main_old() {
	magazine := []string{"t", "e", "s", "t"}
	note := []string{"t", "e", "x"}
	checkMagazine(magazine, note)
}
