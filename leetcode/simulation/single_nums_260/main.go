package main

func main() {

}

func singleNumber(nums []int) (ans []int) {
	countMap := make(map[int]int)
	for _, i := range nums {
		countMap[i]++
		if countMap[i] == 2 {
			delete(countMap, i)
		}
	}
	for k := range countMap {
		ans = append(ans, k)
	}
	return ans
}