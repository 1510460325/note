package com.wzy.note.bfs.highestPeak_1765;

import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Queue;

class Solution {

    class loc {
        private int x, y;

        public loc(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public boolean isValid(int xx, int yy) {
            return x >= 0 && x < xx && y >= 0 && y < yy;
        }
    }

    private int[][] dirs = new int[][]{{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

    public int[][] highestPeak(int[][] isWater) {

        int m = isWater.length, n = isWater[0].length;
        int[][] rst = new int[m][n];

        Queue<loc> queue = new ArrayDeque<>();

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (isWater[i][j] == 1) {
                    rst[i][j] = 0;
                    queue.offer(new loc(i, j));
                } else {
                    rst[i][j] = -1;
                }
            }
        }

        while (!queue.isEmpty()) {
            loc p = queue.poll();
            for (int[] dir : dirs) {
                loc next = new loc(p.x + dir[0], p.y + dir[1]);
                if (next.isValid(m, n) && rst[next.x][next.y] == -1) {
                    rst[next.x][next.y] = rst[p.x][p.y] + 1;
                    queue.offer(next);
                }
            }
        }

        return rst;
    }

    public static void main1(String[] args) {
        int[][] graph = new int[][]{
                {0, 1},
                {0, 0}
        };
        System.out.println(Arrays.deepToString(new Solution().highestPeak(graph)));
    }
}