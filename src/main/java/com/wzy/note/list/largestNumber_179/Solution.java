package com.wzy.note.list.largestNumber_179;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution {
    public String largestNumber(int[] param) {
        List<String> nums = new ArrayList<>();
        for (int j : param) {
            nums.add(String.valueOf(j));
        }
        nums.sort((a, b) -> Long.valueOf(b + a).compareTo(Long.valueOf(a + b)));
        StringBuilder sb = new StringBuilder();
        for (String i : nums) {
            sb.append(i);
        }
        char[] rst = sb.toString().toCharArray();
        for (int i = 0; i < rst.length; i++) {
            if (rst[i] != '0') {
                return String.valueOf(Arrays.copyOfRange(rst, i, rst.length));
            }
        }
        return "0";
    }

    // 9534330
// 9534303
    public static void main(String[] args) {
        int[] p = new int[]{0, 0};
        System.out.println(new Solution().largestNumber(p));
    }
}
