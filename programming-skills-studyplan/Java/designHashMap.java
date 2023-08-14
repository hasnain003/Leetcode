/*
Design a HashMap without using any built-in hash table libraries.
https://leetcode.com/problems/design-hashmap/description/?envType=study-plan-v2&envId=programming-skills.
*/
class Node {
    public int key,value;
    public Node(int key, int value) {
        this.key = key;
        this.value = value;
    }
}
class MyHashMap {
    private ArrayList<ArrayList<Node>> hashMap;
    private final int hash = 1009;
    public MyHashMap() {
        this.hashMap = new ArrayList<>();
        for (int i = 0; i < hash;i++) {
            hashMap.add(new ArrayList<Node>());
        }
    }
    
    public void put(int key, int value) {
        int index = key % hash;
        for (int i = 0; i < this.hashMap.get(index).size();i++) {
            if (this.hashMap.get(index).get(i).key == key) {
                this.hashMap.get(index).set(i,new Node(key,value));
                return;
            }
        }
        this.hashMap.get(index).add(new Node(key,value));
    }
    
    public int get(int key) {
        int index = key % hash;
        for (int i = 0; i < this.hashMap.get(index).size();i++) {
            if (this.hashMap.get(index).get(i).key == key) {
                return this.hashMap.get(index).get(i).value;
            }
        }
        return -1;
    }
    
    public void remove(int key) {
        int index = key % hash;
        for (int i = 0; i < this.hashMap.get(index).size();i++) {
            if (this.hashMap.get(index).get(i).key == key) {
                this.hashMap.get(index).remove(i);
                return;
            }
        }
    }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * MyHashMap obj = new MyHashMap();
 * obj.put(key,value);
 * int param_2 = obj.get(key);
 * obj.remove(key);
 */