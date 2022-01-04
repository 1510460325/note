package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cache := make(map[int]*Node)
	var clone func(cur *Node) *Node
	clone = func(cur *Node) *Node {
		if cache[cur.Val] != nil {
			return cache[cur.Val]
		}
		c := &Node{Val: cur.Val}
		cache[cur.Val] = c
		for _, item := range cur.Neighbors {
			c.Neighbors = append(c.Neighbors, clone(item))
		}
		return c
	}
	return clone(node)
}
