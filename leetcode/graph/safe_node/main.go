package safe_node

import "sort"

func eventualSafeNodes(graph [][]int) (ans []int) {
	safe := make([]bool, len(graph))
	for {
		found := false
		for i := 0; i < len(graph); i++ {
			if safe[i] {
				continue
			}
			s := true
			for _, e := range graph[i] {
				s = s && safe[e]
			}
			if s {
				safe[i] = true
				found = true
				ans = append(ans, i)
			}
		}
		if !found {
			break
		}
	}
	sort.Ints(ans)
	return ans
}
