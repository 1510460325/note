package com.wzy.note.list.detectCycle_142;


import java.util.HashMap;
import java.util.HashSet;

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
        next = null;
    }
}

public class Solution {
    public ListNode detectCycle(ListNode head) {
        HashSet<ListNode> found = new HashSet<>();
        while (head != null) {
            if (found.contains(head)) {
                return head;
            }
            found.add(head);
            head = head.next;
        }
        return null;
    }
}