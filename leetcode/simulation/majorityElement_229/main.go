package main

func main() {

}

func majorityElement(nums []int) (ans []int) {
	countMap := make(map[int]int)
	for _, i := range nums {
		countMap[i]++
	}
	for k, v := range countMap {
		if v > len(nums)/3 {
			ans = append(ans, k)
		}
	}
	return ans
}
