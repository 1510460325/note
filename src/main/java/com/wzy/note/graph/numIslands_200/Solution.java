package com.wzy.note.graph.numIslands_200;

class Solution {
    public int numIslands(char[][] grid) {
        int rst = 0;
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (destroy(grid, i, j, 'x')) {
                    rst++;
                }
            }
        }
        return rst;
    }

    private boolean destroy(char[][] grid, int i, int j, char owner) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != '1') {
            return false;
        }
        grid[i][j] = owner;
        destroy(grid, i - 1, j, owner);
        destroy(grid, i + 1, j, owner);
        destroy(grid, i, j - 1, owner);
        destroy(grid, i, j + 1, owner);
        return true;
    }
}