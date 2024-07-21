package main

import (
	"fmt"
	"math"
)

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	var l, e, h []int
	check := arr[0]
	for _, num := range arr {
		if num < check {
			l = append(l, num)
		} else if num > check {
			h = append(h, num)
		} else {
			e = append(e, num)
		}
	}
	return append(append(sort(l), e...), sort(h)...)
}
func sortpartial(arr []int, index int) []int {
	parr := sort(arr[index:])
	for i := 0; i < len(parr); i++ {
		arr[index+i] = parr[i]
	}
	return arr
}

func getNextMaxI(nums []int, fromIndex int) int {
	foundI := 0
	minDiff := math.MaxInt
	target := nums[fromIndex]
	for i := fromIndex + 1; i < len(nums); i++ {
		diff := nums[i] - target
		if diff > 0 && diff < minDiff {
			minDiff = diff
			foundI = i
		}
	}
	return foundI
}

func order(arr []int) ([]int, bool) {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] < arr[i] {
			nextMaxI := getNextMaxI(arr, i-1)
			arr[i-1], arr[nextMaxI] = arr[nextMaxI], arr[i-1]
			arr = sortpartial(arr, i)
			return arr, false
		}
	}
	return nil, true
}

func gen(n int) {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	finished := false
	for !finished {
		arr, finished = order(arr)
		fmt.Println(arr)
	}
}

func main() {
	gen(3)
}
