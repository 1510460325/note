package com.wzy.note.graph.canFinish_207;

import java.util.*;

class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {

        List<Set<Integer>> preSet = new ArrayList<>(numCourses);
        for (int i = 0; i < numCourses; i++) {
            preSet.add(new HashSet<>());
        }

        for (int[] prerequisite : prerequisites) {
            int s = prerequisite[0], e = prerequisite[1];
            preSet.get(s).add(e);
        }

        Map<Integer, Boolean> studied = new HashMap<>();
        while (true) {
            boolean study = false;
            for (int i = 0; i < preSet.size(); i++) {
                Set<Integer> set = preSet.get(i);
                Boolean ok = studied.get(i);
                if (ok != null) {
                    continue;
                }
                boolean success = true;
                for (Integer p : set) {
                    ok = studied.get(p);
                    if (ok == null) {
                        success = false;
                        break;
                    }
                }
                if (success) {
                    studied.put(i, true);
                    study = true;
                    break;
                }
            }
            if (!study) {
                break;
            }
        }
        return studied.size() == numCourses;
    }
}