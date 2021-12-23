package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	idx := 0
	for i := range inorder {
		if postorder[len(postorder)-1] == inorder[i] {
			idx = i
			break
		}
	}
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1]),
	}
}
