package main

func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	n := total / 2
	pre, now := -1, -1
	for i := 0; i <= n; i++ {
		pre = now
		if len(nums1) == 0 {
			now = nums2[0]
			nums2 = nums2[1:]
			continue
		}
		if len(nums2) == 0 {
			now = nums1[0]
			nums1 = nums1[1:]
			continue
		}
		if nums1[0] < nums2[0] {
			now = nums1[0]
			nums1 = nums1[1:]
		} else {
			now = nums2[0]
			nums2 = nums2[1:]
		}
	}
	if total % 2 == 0 {
		return float64(pre + now) / 2
	}
	return float64(now)
}
