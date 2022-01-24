package com.wzy.note.list.findOrder_210;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {
    public int[] findOrder(int numCourses, int[][] prerequisites) {

        List<Set<Integer>> preSet = new ArrayList<>(numCourses);
        for (int i = 0; i < numCourses; i++) {
            preSet.add(new HashSet<>());
        }

        for (int[] p : prerequisites) {
            preSet.get(p[0]).add(p[1]);
        }

        Set<Integer> studied = new HashSet<>();
        int[] rst = new int[numCourses];
        int idx = 0;

        while (true) {
            int found = -1;
            for (int i = 0; i < numCourses; i++) {
                if (studied.contains(i)) {
                    continue;
                }
                boolean success = true;
                for (Integer pre : preSet.get(i)) {
                    if (!studied.contains(pre)) {
                        success = false;
                        break;
                    }
                }
                if (success) {
                    found = i;
                    studied.add(i);
                    break;
                }
            }
            if (found == -1) {
                break;
            }
            rst[idx++] = found;
        }
        if (idx != numCourses) {
            return new int[]{};
        }
        return rst;
    }
}