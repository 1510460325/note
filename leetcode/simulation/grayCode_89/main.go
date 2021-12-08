package main

import "fmt"

func main() {
	fmt.Println(grayCode(8))
}

func grayCode(n int) (ans []int) {
	used := make(map[int]bool)
	cur := 0
	count := 1 << n
	for len(used) < count {
		used[cur] = true
		ans = append(ans, cur)
		for i := 0; i <= n; i++ {
			if cur&(1<<i) > 0 && !used[cur-1<<i] {
				cur -= 1 << i
				break
			}
			if cur&(1<<i) == 0 && !used[cur+1<<i] {
				cur += 1 << i
				break
			}
		}
	}
	return ans
}
