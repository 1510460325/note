package com.wzy.note.list.minSubArrayLen;

class Solution {
    public int minSubArrayLen(int target, int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        final int M = 1 << 31 - 1;
        int rst = M, i = 0, j = 0, sum = nums[0];
        while (j < nums.length) {
            if (sum >= target) {
                if (i == j) {
                    return 1;
                }
                rst = Integer.min(rst, j - i + 1);
                sum -= nums[i];
                i++;
            } else {
                j++;
                if (j != nums.length) {
                    sum += nums[j];
                }
            }
        }
        if (rst == M) {
            return 0;
        }
        return rst;
    }
}