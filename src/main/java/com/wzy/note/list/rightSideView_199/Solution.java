package com.wzy.note.list.rightSideView_199;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

class Solution {
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> rst = new ArrayList<>();
        if (root == null) {
            return rst;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(root);
        while (queue.size() > 0) {
            Queue<TreeNode> next = new LinkedList<>();
            TreeNode cur;
            int last = -1;
            while ((cur = queue.poll()) != null) {
                last = cur.val;
                if (cur.left != null) {
                    next.offer(cur.left);
                }
                if (cur.right != null) {
                    next.offer(cur.right);
                }
            }
            rst.add(last);
            queue = next;
        }
        return rst;
    }
}