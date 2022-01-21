package com.wzy.note.tree.trie;

import java.util.HashMap;
import java.util.Map;

class Trie {

    private final Map<Character, Trie> child;
    private boolean canEnd;

    public Trie() {
        this.canEnd = false;
        this.child = new HashMap<>();
    }

    public void insert(String word) {
        if (word.length() == 0) {
            this.canEnd = true;
            return;
        }
        Character c = word.charAt(0);
        Trie trie = this.child.get(c);
        if (trie == null) {
            trie = new Trie();
            this.child.put(c, trie);
        }
        trie.insert(word.substring(1));
    }

    public boolean search(String word) {
        if (word.length() == 0) {
            return this.canEnd;
        }
        Character c = word.charAt(0);
        Trie trie = this.child.get(c);
        if (trie == null) {
            return false;
        }
        return trie.search(word.substring(1));
    }

    public boolean startsWith(String prefix) {
        if (prefix.length() == 0) {
            return true;
        }
        Character c = prefix.charAt(0);
        Trie trie = this.child.get(c);
        if (trie == null) {
            return false;
        }
        return trie.startsWith(prefix.substring(1));
    }

    public static void main1(String[] args) {
        Trie obj = new Trie();
        obj.insert("apple");
        System.out.println(obj.search("apple"));
        System.out.println(obj.search("app"));
        System.out.println(obj.startsWith("app"));
        obj.insert("app");
        System.out.println(obj.search("app"));
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */