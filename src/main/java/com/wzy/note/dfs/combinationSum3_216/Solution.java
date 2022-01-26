package com.wzy.note.dfs.combinationSum3_216;

import java.util.*;

class Solution {
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> rst = new ArrayList<>();
        dfs(rst, new Stack<>(), n, k);
        return rst;
    }

    private void dfs(List<List<Integer>> rst, Stack<Integer> cur, int target, int rest) {
        if (target < 0) {
            return;
        }
        if (target == 0) {
            if (rest == 0) {
                rst.add(new ArrayList<>(cur));
            }
            return;
        }
        if (rest == 0) {
            return;
        }
        int pre = 0;
        if (cur.size() > 0) {
            pre = cur.peek();
        }
        for (int i = pre + 1; i <= 9; i++) {
            cur.push(i);
            dfs(rst, cur, target - i, rest - 1);
            cur.pop();
        }
    }
}