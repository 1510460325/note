package com.wzy.note.list.copyRandomList;


import java.util.HashMap;

// Definition for a Node.
class Node {
    int val;
    Node next;
    Node random;

    public Node(int val) {
        this.val = val;
        this.next = null;
        this.random = null;
    }
}


class Solution {


    public static void main(String[] args) {

    }

    private final HashMap<Node, Node> cache = new HashMap<>();

    public Node copyRandomList(Node head) {
        if (head == null) {
            return null;
        }
        Node val = cache.get(head);
        if (val != null) {
            return val;
        }
        val = new Node(head.val);
        cache.put(head, val);
        val.next = copyRandomList(head.next);
        val.random = copyRandomList(head.random);
        return cache.get(head);
    }

}