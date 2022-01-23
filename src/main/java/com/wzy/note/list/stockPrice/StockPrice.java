package com.wzy.note.list.stockPrice;

import java.util.*;

class StockPrice {

    private final SortedMap<Integer, Set<Integer>> map;

    private final Map<Integer, Integer> priceMap;

    private int latestMs;

    public StockPrice() {
        map = new TreeMap<>();
        priceMap = new HashMap<>();
    }

    public void update(int timestamp, int price) {

        latestMs = Integer.max(latestMs, timestamp);

        Integer oldPrice = priceMap.get(timestamp);

        if (oldPrice != null && oldPrice == price) {
            return;
        }
        priceMap.put(timestamp, price);

        if (oldPrice != null) {
            // delete
            Set<Integer> source = map.get(oldPrice);
            source.remove(timestamp);
            if (source.size() == 0) {
                map.remove(oldPrice);
            }
        }
        // add
        Set<Integer> target = map.get(price);
        if (target == null) {
            target = new HashSet<>();
            target.add(timestamp);
            map.put(price, target);
        } else {
            target.add(timestamp);
        }
    }

    public int current() {
        return priceMap.getOrDefault(latestMs, 0);
    }

    public int maximum() {
        Integer val = map.lastKey();
        if (val == null) {
            return 0;
        }
        return val;
    }

    public int minimum() {
        Integer val = map.firstKey();
        if (val == null) {
            return 0;
        }
        return val;
    }

}