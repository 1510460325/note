package main

func main() {

}

func findRepeatedDnaSequences(s string) (ans []string) {
	countMap := make(map[string]int)
	for i := 10; i <= len(s); i++ {
		countMap[s[i-10:i]]++
	}
	for k, v := range countMap {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
