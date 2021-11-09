package main

func main() {

}

func removeElement(nums []int, val int) int {
	end := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[end] = nums[i]
		end++
	}
	return end
}