package main

func main() {

}

func loudAndRich(richer [][]int, quiet []int) (ret []int) {
	richSet := make([]map[int]bool, len(quiet))
	for i := range richSet {
		richSet[i] = make(map[int]bool)
	}
	for _, item := range richer {
		a, b := item[0], item[1]
		richSet[b][a] = true
	}
	done := make([]bool, len(quiet))
	var dfs func(now int)
	dfs = func(now int) {
		if done[now] {
			return
		}
		for k := range richSet[now] {
			dfs(k)
		}
		tmp := make(map[int]bool)
		for k := range richSet[now] {
			tmp[k] = true
			for kk := range richSet[k] {
				tmp[kk] = true
			}
		}
		richSet[now] = tmp
		done[now] = true
	}
	for i := 0; i < len(quiet); i++ {
		dfs(i)
		min := i
		for j := range richSet[i] {
			if quiet[min] > quiet[j] {
				min = j
			}
		}
		ret = append(ret, min)
	}
	return ret
}
