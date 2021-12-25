package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue, reverse := []*TreeNode{root}, false
	for len(queue) > 0 {
		next := make([]*TreeNode, 0, len(queue)*2)
		cur := make([]int, 0, len(queue))
		for _, item := range queue {
			if item.Val%2 == 1 && reverse ||
				item.Val%2 == 0 && !reverse {
				return false
			}
			if len(cur) > 0 {
				if reverse && item.Val >= cur[len(cur)-1] { // 奇数下标：偶数递减
					return false
				} else if !reverse && item.Val <= cur[len(cur)-1] { // 偶数下标：奇数递增
					return false
				}
			}
			cur = append(cur, item.Val)
			if item.Left != nil {
				next = append(next, item.Left)
			}
			if item.Right != nil {
				next = append(next, item.Right)
			}
		}
		queue, reverse = next, !reverse
	}
	return true
}
