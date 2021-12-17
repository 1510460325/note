package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var dfs func(n *TreeNode)
	list := make([]*TreeNode, 0)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		list = append(list, n)
		dfs(n.Right)
	}
	dfs(root)
	first, second := -1, -1
	for i := 0; i < len(list)-1; i++ {
		if list[i].Val > list[i+1].Val {
			if first == -1 {
				first = i
			} else {
				second = i
				break
			}
		}
	}
	if second != -1 {
		list[second+1].Val, list[first].Val = list[first].Val, list[second+1].Val
	} else {
		list[first+1].Val, list[first].Val = list[first].Val, list[first+1].Val
	}
}
