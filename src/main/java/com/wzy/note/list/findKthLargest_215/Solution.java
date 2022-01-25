package com.wzy.note.list.findKthLargest_215;

class Solution {
    public int findKthLargest(int[] nums, int k) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int i = l, j = r;
            while (i < j) {
                while (i < j && nums[i] >= nums[j]) {
                    j--;
                }
                if (i == j) {
                    break;
                }
                int tmp = nums[i];
                nums[i] = nums[j];
                nums[j] = tmp;
                while (i < j && nums[i] >= nums[j]) {
                    i++;
                }
                if (i == j) {
                    break;
                }
                tmp = nums[i];
                nums[i] = nums[j];
                nums[j] = tmp;
            }
            if (i == k - 1) {
                return nums[i];
            } else if (i < k) {
                l = i + 1;
            } else {
                r = i - 1;
            }
        }
        return nums[l];
    }
}