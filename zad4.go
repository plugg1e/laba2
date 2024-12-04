package main

import (
	"fmt"
	"strconv"
)

type Arr struct {
	number   []string
	sizes    int
	capacity int
}

func NewArr(cap int) *Arr {
	return &Arr{
		number:   make([]string, cap),
		sizes:    0,
		capacity: cap,
	}
}

func (a *Arr) reSize() {
	a.capacity *= 2
	newNumber := make([]string, a.capacity)
	copy(newNumber, a.number)
	a.number = newNumber
}

func (a *Arr) AddArr(value string) {
	if a.sizes >= a.capacity {
		a.reSize()
	}
	a.number[a.sizes] = value
	a.sizes++
}

func (a *Arr) Size() int {
	return a.sizes
}

func findSubarrayWithSum(arr *Arr, target int) [][]int {
	var result [][]int
	n := arr.Size()

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			num, _ := strconv.Atoi(arr.number[j])
			sum += num
			if sum == target {
				var subarray []int
				for k := i; k <= j; k++ {
					num, _ := strconv.Atoi(arr.number[k])
					subarray = append(subarray, num)
				}
				result = append(result, subarray)
			}
		}
	}

	return result
}

func main() {
	arr := NewArr(10)
	elements := []string{"4", "-7", "1", "5", "-4", "0", "-3", "2", "4", "1"}
	for _, elem := range elements {
		arr.AddArr(elem)
	}

	target := 6
	subarrays := findSubarrayWithSum(arr, target)
	fmt.Println("подмассиивы с суммойй", target, ":", subarrays)
}
