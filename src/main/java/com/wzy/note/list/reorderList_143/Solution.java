package com.wzy.note.list.reorderList_143;


class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}

class Solution {
    public void reorderList(ListNode head) {
        if (head == null) {
            return;
        }
        ListNode middle = split(head);
        ListNode head1 = reverse(middle);

        while (head1 != null) {
            ListNode next = head.next;
            ListNode next1 = head1.next;
            head.next = head1;
            head1.next = next;
            if (next != null) {
                head = next;
            } else {
                head = head1;
            }
            head1 = next1;
        }
    }

    private static ListNode split(ListNode head) {
        if (head == null) {
            return null;
        }
        ListNode pre = null;
        ListNode slow = head;
        ListNode fast = head;
        while (fast != null && fast.next != null) {
            pre = slow;
            slow = slow.next;
            fast = fast.next.next;
        }
        if (pre != null) {
            pre.next = null;
        }
        return slow;
    }

    private static ListNode reverse(ListNode head) {
        ListNode rst = null;
        // 头插法
        while (head != null) {
            ListNode next = head.next;
            head.next = rst;
            rst = head;
            head = next;
        }
        return rst;
    }
}