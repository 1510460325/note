package main

func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	end := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[end] {
			end++
			nums[end] = nums[i]
		}
	}
	return end + 1
}
