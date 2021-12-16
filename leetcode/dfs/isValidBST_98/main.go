package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTWithLimit(nil, nil, root)
}

func isValidBSTWithLimit(min, max, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil {
		if min.Val >= root.Val {
			return false
		}
	}
	if max != nil {
		if max.Val <= root.Val {
			return false
		}
	}
	return isValidBSTWithLimit(min, root, root.Left) && isValidBSTWithLimit(root, max, root.Right)
}
