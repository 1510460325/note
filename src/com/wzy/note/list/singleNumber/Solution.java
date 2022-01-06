package com.wzy.note.list.singleNumber;

import java.util.HashMap;

class Solution {

    public static void main(String[] args) {
        Solution s = new Solution();
        int[] a = {1, 2, 3};
        System.out.println(s.singleNumber(a));
    }

    public int singleNumber(int[] nums) {
        HashMap<Integer, Integer> cache = new HashMap<>(nums.length);
        for (int num : nums) {
            cache.merge(num, 1, Integer::sum);
        }
        for (Integer key : cache.keySet()) {
            if (cache.get(key) == 1) {
                return key;
            }
        }
        return -1;
    }
}
