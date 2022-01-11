package com.wzy.note.list.insertionSortList;

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
    public ListNode insertionSortList(ListNode head) {
        ListNode rst = null;
        while (head != null) {
            ListNode next = head.next;
            head.next = null;
            if (rst == null) {
                rst = head;
            } else if (rst.val > head.val) {
                head.next = rst;
                rst = head;
            } else {
                ListNode cur = rst;
                while (cur.next != null) {
                    if (cur.next.val > head.val) {
                        break;
                    }
                    cur = cur.next;
                }
                head.next = cur.next;
                cur.next = head;
            }
            head = next;
        }
        return rst;
    }
}