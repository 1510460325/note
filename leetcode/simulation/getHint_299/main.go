package main

import "fmt"

func main() {
	fmt.Println(getHint("1122","1222"))
}

func getHint(secret string, guess string) string {
	a, b := 0, 0
	countMap := make(map[uint8]int)
	for i := range secret {
		countMap[secret[i]]++
	}
	for i := range secret {
		s, g := secret[i], guess[i]
		if s == g {
			countMap[s]--
			a++
			continue
		}
	}
	for i := range secret {
		s, g := secret[i], guess[i]
		if s == g {
			continue
		}
		if countMap[g] > 0 {
			countMap[g]--
			b++
		}
	}
	return fmt.Sprintf("%vA%vB", a, b)
}
