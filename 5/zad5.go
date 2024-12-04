package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			index := i
			if !leftToRight {
				index = levelSize - 1 - i
			}

			level[index] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 22,
		Left: &TreeNode{
			Val: 16,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 19,
			},
		},
		Right: &TreeNode{
			Val: 51,
			Left: &TreeNode{
				Val: 43,
			},
			Right: &TreeNode{
				Val: 57,
			},
		},
	}

	result := zigzagLevelOrder(root)
	fmt.Println("Змейка:", result)
}
