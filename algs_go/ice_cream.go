package main

import "fmt"

func whatFlavors(cost []int, money int) {
	var ids [][]int
	temp := make([]int, 2)
	for i := 0; i < len(cost); i++ {
		for j := 0; j < len(cost); j++ {
			if i != j {
				fmt.Println(i, j)
				sum := cost[i] + cost[j]
				if sum == money {
					temp[0] = i
					temp[1] = j
					ids = append(ids, temp)
				}
			}
		}
	}
	fmt.Println(ids)
}

func main() {
	cost := []int{1, 2, 4}
	money := 3
	whatFlavors(cost, money)
}
