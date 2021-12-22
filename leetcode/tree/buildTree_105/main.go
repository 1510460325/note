package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	e := 0
	for i := range inorder {
		if preorder[0] == inorder[i] {
			e = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:e+1], inorder[:e]),
		Right: buildTree(preorder[e+1:], inorder[e+1:]),
	}
}
