package main

import "fmt"

func main() {
	fmt.Println(pathSum(&TreeNode{Val: 1}, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	return pathSumWithStarted(root, targetSum, false)
}

func pathSumWithStarted(root *TreeNode, targetSum int, started bool) int {
	if root == nil {
		return 0
	}
	if started {
		targetSum -= root.Val
		now := 0
		if targetSum == 0 {
			now = 1
		}
		return pathSumWithStarted(root.Left, targetSum, true) +
			pathSumWithStarted(root.Right, targetSum, true) + now
	} else {
		now := 0
		if targetSum-root.Val == 0 {
			now = 1
		}
		return pathSumWithStarted(root.Left, targetSum-root.Val, true) +
			pathSumWithStarted(root.Right, targetSum-root.Val, true) +
			pathSumWithStarted(root.Left, targetSum, false) +
			pathSumWithStarted(root.Right, targetSum, false) + now
	}
}
