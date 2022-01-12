package com.wzy.note.list.sortList;

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
    public ListNode sortList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode fast = head, slow = head;
        while (fast.next != null && fast.next.next != null) {
            fast = fast.next.next;
            slow = slow.next;
        }
        ListNode mid = slow.next;
        slow.next = null;
        return merge(sortList(head), sortList(mid));
    }

    private ListNode merge(ListNode l1, ListNode l2) {
        ListNode rst = new ListNode();
        ListNode tail = rst;
        while (l1 != null || l2 != null) {
            if (l1 == null || (l2 != null && l2.val < l1.val)) {
                tail.next = l2;
                l2 = l2.next;
            } else {
                tail.next = l1;
                l1 = l1.next;
            }
            tail = tail.next;
        }
        return rst.next;
    }
}


