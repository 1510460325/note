package main

func main() {

}

func longestCommonPrefix(strList []string) (maxPrefix string) {
	maxPrefix = strList[0]
	for i := 1; i < len(strList); i++ {
		for j := 0; j < len(maxPrefix); j++ {
			if j == len(strList[i]) || maxPrefix[j] != strList[i][j] {
				maxPrefix = maxPrefix[:j]
				break
			}
		}
	}
	return maxPrefix
}
