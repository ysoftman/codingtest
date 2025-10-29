/*
https://leetcode.com/problems/design-hashmap/
706. Design HashMap
Easy
Design a HashMap without using any built-in hash table libraries.
Implement the MyHashMap class:
MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

Example 1:
Input
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output
[null, null, null, 1, -1, null, 1, null, -1]

Explanation
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]


Constraints:
0 <= key, value <= 106
At most 104 calls will be made to put, get, and remove.
*/
package main

import "fmt"

type MyHashMap struct {
	hm []int
}

func Constructor() MyHashMap {
	mhm := MyHashMap{
		hm: make([]int, 1000001),
	}
	for i := range mhm.hm {
		mhm.hm[i] = -1
	}
	return mhm
}

func (this *MyHashMap) Put(key int, value int) {
	this.hm[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.hm[key]
}

func (this *MyHashMap) Remove(key int) {
	this.hm[key] = -1
}
func (this *MyHashMap) Print() {
	fmt.Print("[")
	for i := range this.hm {
		if this.hm[i] != -1 {
			fmt.Printf("[%v,%v] ", i, this.hm[i])
		}
	}
	fmt.Print("]\n")
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	myHashMap := Constructor()
	myHashMap.Put(1, 1) // The map is now [[1,1]]
	myHashMap.Print()
	myHashMap.Put(2, 2) // The map is now [[1,1], [2,2]]
	myHashMap.Print()
	fmt.Println(myHashMap.Get(1)) // return 1, The map is now [[1,1], [2,2]]
	fmt.Println(myHashMap.Get(3)) // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	myHashMap.Put(2, 1)           // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	myHashMap.Print()
	fmt.Println(myHashMap.Get(2)) // return 1, The map is now [[1,1], [2,1]]
	myHashMap.Remove(2)           // remove the mapping for 2, The map is now [[1,1]]
	myHashMap.Print()
	fmt.Println(myHashMap.Get(2)) // return -1 (i.e., not found), The map is now [[1,1]]
}
