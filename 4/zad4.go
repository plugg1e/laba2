package main

import "fmt"

func findSubarrayWithSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	prefixSum := make(map[int]int)
	sum := 0

	for i, num := range nums {
		sum += num

		if sum == target {
			result = append(result, nums[:i+1])
		}

		if val, ok := prefixSum[sum-target]; ok {
			result = append(result, nums[val+1:i+1])
		}

		prefixSum[sum] = i
	}

	return result
}

func main() {
	nums := []int{4, -7, 1, 5, -4, 0, -3, 2, 4, 1}
	target := 5
	subarrays := findSubarrayWithSum(nums, target)
	fmt.Println("Подмассивы с суммой", target, ":", subarrays)
}
