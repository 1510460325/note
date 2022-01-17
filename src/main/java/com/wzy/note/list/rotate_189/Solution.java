package com.wzy.note.list.rotate_189;

import java.util.Arrays;

class Solution {
    public void rotate(int[] nums, int k) {
        k %= nums.length;
        int[] tmp = new int[k];
        System.arraycopy(nums, nums.length - k, tmp, 0, k);
        System.arraycopy(nums, 0, nums, k, nums.length - k);
        System.arraycopy(tmp, 0, nums, 0, k);
    }

    public static void main(String[] args) {
        int[] p = new int[]{-1};
        new Solution().rotate(p, 2);
        System.out.println(Arrays.toString(p));
    }
}