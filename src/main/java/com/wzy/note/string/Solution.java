package com.wzy.note.string;

import java.util.List;

class Solution {

    public String reverseWords(String s) {
        StringBuilder cur = new StringBuilder();
        StringBuilder rst = new StringBuilder();
        char[] chs = s.toCharArray();
        for (int i = 0; i < chs.length; i++) {
            if (chs[i] != ' ') {
                cur.append(chs[i]);
            }
            if ((chs[i] == ' ' || i == chs.length - 1) && cur.length() != 0) {
                if (rst.length() > 0) {
                    rst.insert(0, cur + " ");
                } else {
                    rst.append(cur);
                }
                cur = new StringBuilder();
            }
        }
        return rst.toString();
    }

    public static void main(String[] args) {
        System.out.println(new Solution().reverseWords("   the    sky  is  blue  "));
    }
}