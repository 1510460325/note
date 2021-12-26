package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ret [][]int) {
	if root == nil {
		return nil
	}
	var dfs func(path []int, rest int, head *TreeNode)
	dfs = func(path []int, rest int, head *TreeNode) {
		if head.Left == nil && head.Right == nil && rest == 0 {
			cur := make([]int, len(path))
			copy(cur, path)
			ret = append(ret, cur)
			return
		}
		if head.Left != nil {
			dfs(append(path, head.Left.Val), rest-head.Left.Val, head.Left)
		}
		if head.Right != nil {
			dfs(append(path, head.Right.Val), rest-head.Right.Val, head.Right)
		}
	}
	dfs([]int{root.Val}, targetSum-root.Val, root)
	return ret
}
