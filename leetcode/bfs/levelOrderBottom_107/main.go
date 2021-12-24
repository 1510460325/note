package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) (ret [][]int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		next := make([]*TreeNode, 0, len(queue)*2)
		curRet := make([]int, 0, len(queue))
		for _, item := range queue {
			curRet = append(curRet, item.Val)
			if item.Left != nil {
				next = append(next, item.Left)
			}
			if item.Right != nil{
				next = append(next, item.Right)
			}
		}
		queue = next
		ret = append([][]int{curRet}, ret...)
	}
	return ret
}
