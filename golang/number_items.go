package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) > 1 {
		partition := arr[0]
		size := len(arr)
		larger := make([]int, 0, size)
		smaller := make([]int, 0, size)
		equal := make([]int, 1, size)
		equal[0] = partition
		for i := 1; i < len(arr); i++ {
			if arr[i] > partition {
				larger = append(larger, arr[i])
			} else if arr[i] < partition {
				smaller = append(smaller, arr[i])
			} else {
				equal = append(equal, arr[i])
			}
		}
		return append(append(mergeSort(smaller), equal...), mergeSort(larger)...)

	} else {
		return arr
	}
}

func maximumToys(prices []int, k int) []int {
	prices = mergeSort(prices)
	fmt.Println(k)
	// fmt.Println(prices)
	// var maxNum int
	// maxSum := 0
	// for i := 0; i < len(prices); i++ {
	// 	maxSum += prices[i]
	// 	if maxSum > k {
	// 		maxNum = i
	// 		break
	// 	}
	// }
	return prices
}

func main() {
	prices := []int{11, 2, 5, 5, 1}
	k := 8
	number := maximumToys(prices, k)
	fmt.Println(number)
}
