package main

import "fmt"

func main() {
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}

func compareVersion(version1 string, version2 string) int {
	if version1 == "" && version2 == "" {
		return 0
	}
	num1, e1 := 0, -1
	for i := range version1 {
		if version1[i] == '.' {
			e1 = i
			break
		}
		num1 = num1*10 + int(version1[i]-'0')
	}
	if e1 == -1 {
		version1 = ""
	} else {
		version1 = version1[e1+1:]
	}
	num2, e2 := 0, -1
	for i := range version2 {
		if version2[i] == '.' {
			e2 = i
			break
		}
		num2 = num2*10 + int(version2[i]-'0')
	}
	if e2 == -1 {
		version2 = ""
	} else {
		version2 = version2[e2+1:]
	}
	if num1 < num2 {
		return -1
	} else if num1 > num2 {
		return 1
	}

	return compareVersion(version1, version2)
}
