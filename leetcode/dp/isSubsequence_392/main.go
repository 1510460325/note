package main

func main() {

}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	matchIndex := 0
	for i := range t {
		if t[i] == s[matchIndex] {
			matchIndex++
			if matchIndex >= len(s) {
				return true
			}
		}
	}
	return false
}