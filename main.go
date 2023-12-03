package main

import (
	"fmt"
)

func main() {
	fmt.Println(deleteAndEarn([]int{8, 3, 4, 7, 6, 6, 9, 2, 5, 8, 2, 4, 9, 5, 9, 1, 5, 7, 1, 4}))
}

func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = maxv(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

func rob(nums []int) int {
	first, second := nums[0], 0
	for i := 1; i < len(nums); i++ {
		second = first
		first = maxv()
	}
	return second
}

func maxv(a, b int) int {
	if a > b {
		return a
	}
	return b
}
