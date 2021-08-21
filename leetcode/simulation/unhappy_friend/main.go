package main

func main() {

}

func unhappyFriends(n int, preferences [][]int, pairs [][]int) (ans int) {
	target := make([]int, n)
	for _, item := range pairs {
		a, b := item[0], item[1]
		target[a] = b
		target[b] = a
	}
	score := make([]map[int]int, n)
	for i, list := range preferences {
		score[i] = make(map[int]int, n)
		for idx, item := range list {
			score[i][item] = idx
		}
	}
	for i := 0; i < n; i++ {
		targetID := target[i]
		for _, j := range preferences[i] {
			if j != targetID {
				jTarget := target[j]
				if score[j][i] < score[j][jTarget] {
					ans++
					break
				}
			} else {
				break
			}
		}
	}
	return ans
}
