package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return gen(1, n)
}

func gen(start, end int) (ans []*TreeNode) {

	if start > end {
		return nil
	}
	if start == end {
		return []*TreeNode{{Val: start}}
	}
	for i := start; i <= end; i++ {
		left := gen(start, i-1)
		right := gen(i+1, end)
		ans = append(ans, multi(i, left, right)...)
	}
	return ans
}

func multi(val int, left, right []*TreeNode) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if len(left) == 0 && len(right) == 0 {
		ret = append(ret, &TreeNode{Val: val})
	} else if len(left) == 0 {
		for _, rightHead := range right {
			ret = append(ret, &TreeNode{val, nil, rightHead})
		}
	} else if len(right) == 0 {
		for _, leftHead := range left {
			ret = append(ret, &TreeNode{val, leftHead, nil})
		}
	} else {
		for _, leftHead := range left {
			for _, rightHead := range right {
				ret = append(ret, &TreeNode{val, leftHead, rightHead})
			}
		}
	}
	return ret
}
