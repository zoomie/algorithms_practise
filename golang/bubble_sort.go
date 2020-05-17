package main

import (
	"fmt"
)

// func mergeSort(nums []int) []int {
	partition := len(nums) / 2
	if partition == 0 {
		return nums
	}
	checkNum := nums[partition]
	lower := make([]int, len(nums))
	upper := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > partition {
			lower := append(lower, nums[i])
		} else {
			upper := append(upper, nums[i])
		}
	}
	return append(mergeSort(lower), mergeSort(upper)...)
}

func main() {
	nums := []int{4, 2, 3, 5, 1}
	sortedNums := mergeSort(nums)
	fmt.Println()
	fmt.Println(sortedNums)
}
