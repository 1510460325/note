package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	tmp := root.Right
	root.Right, root.Left = root.Left, nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}
