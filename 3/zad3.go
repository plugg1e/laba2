package main

import (
	"fmt"
	"math"
)

func minSubsetDifference(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	dp := make([]bool, totalSum/2+1)
	dp[0] = true

	for _, num := range nums {
		for j := totalSum / 2; j >= num; j-- {
			if dp[j-num] {
				dp[j] = true
			}
		}
	}

	minDiff := math.MaxInt32
	for j := totalSum / 2; j >= 0; j-- {
		if dp[j] {
			minDiff = totalSum - 2*j
			break
		}
	}

	return minDiff
}

func main() {
	nums := []int{5, 8, 1, 14, 7}
	minDiff := minSubsetDifference(nums)
	fmt.Println("Минимальная разница:", minDiff)
}
