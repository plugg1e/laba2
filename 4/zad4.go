package main

import (
	"fmt"
)

func findSubarrayWithSum(arr []int, target int) [][]int {
	var result [][]int
	n := len(arr)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == target {
				result = append(result, arr[i:j+1])
			}
		}
	}

	return result
}

func main() {
	arr := []int{4, -7, 1, 5, -4, 0, -3, 2, 4, 1}
	target := 6
	subarrays := findSubarrayWithSum(arr, target)
	fmt.Println("подмассиивы с суммойй", target, ":", subarrays)
}
