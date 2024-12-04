package main

import (
	"fmt"
)

func minDifference(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	n := len(nums)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, totalSum/2+1)
	}

	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j <= totalSum/2; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	for j := totalSum / 2; j >= 0; j-- {
		if dp[n][j] {
			return totalSum - 2*j
		}
	}

	return totalSum
}

func main() {
	nums := []int{5, 6, 1, 14, 7}
	diff := minDifference(nums)
	fmt.Println("мин разница:", diff)
}
