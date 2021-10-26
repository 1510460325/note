package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) (ret []int) {
	ans := make(map[int]int)
	nums1Set := make(map[int]bool)
	for _, i := range nums1 {
		nums1Set[i] = true
	}
	isPre := make(map[int]bool)
	for _, i := range nums2 {
		for j := range isPre {
			_, ok := ans[j]
			if !ok && i > j {
				ans[j] = i
			}
		}
		if nums1Set[i] {
			isPre[i] = true
		}
	}
	for _, i := range nums1 {
		a, ok := ans[i]
		if !ok {
			a = -1
		}
		ret = append(ret, a)
	}
	return
}
