package main

func main() {

}

func destCity(paths [][]string) string {
	start := make(map[string]bool)
	for _, item := range paths {
		start[item[0]] = true
	}
	for _, item := range paths {
		if !start[item[1]] {
			return item[1]
		}
	}
	return ""
}