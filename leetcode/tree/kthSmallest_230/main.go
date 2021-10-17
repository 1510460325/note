package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}
	//root = &TreeNode{
	//	Val: 1,
	//	Right: &TreeNode{
	//		Val:   2,
	//	},
	//}
	fmt.Println(kthSmallest(root, 2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	count := 0
	node := root
	for node != nil || len(stack) != 0 {
		if node != nil {
			stack = append([]*TreeNode{node}, stack...)
			node = node.Left
		} else {
			node, stack = stack[0], stack[1:]
			count++
			if count == k {
				return node.Val
			}
			fmt.Println(count, node.Val)
			node = node.Right
		}
	}
	return 0
}
