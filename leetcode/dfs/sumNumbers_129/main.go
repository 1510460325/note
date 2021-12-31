package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) (rst int) {
	if root == nil {
		return 0
	}
	var dfs func(cur int, p *TreeNode)
	dfs = func(cur int, p *TreeNode) {
		if p.Left == nil && p.Right == nil {
			rst += cur
		}
		if p.Left != nil {
			dfs(cur*10+p.Left.Val, p.Left)
		}
		if p.Right != nil {
			dfs(cur*10+p.Right.Val, p.Right)
		}
	}
	dfs(root.Val, root)
	return rst
}
