package main

/*
Design a HashMap without using any built-in hash table libraries.
https://leetcode.com/problems/design-hashmap/description/?envType=study-plan-v2&envId=programming-skills.
*/

type Map struct {
	key   int
	value int
}

type MyHashMap struct {
	maps [][]Map
}

const KeySize = 2069

func HashMapConstructor() MyHashMap {
	return MyHashMap{
		maps: make([][]Map, KeySize),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % KeySize
	for i, node := range this.maps[index] {
		if node.key == key {
			this.maps[index][i].value = value
			return
		}
	}
	this.maps[index] = append(this.maps[index], Map{key, value})
}

func (this *MyHashMap) Get(key int) int {
	index := key % KeySize
	for _, pair := range this.maps[index] {
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % KeySize
	for i, pair := range this.maps[index] {
		if pair.key == key {
			this.maps[index] = append(this.maps[index][:i], this.maps[index][i+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := HashMapConstructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
