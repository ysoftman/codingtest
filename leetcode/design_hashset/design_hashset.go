/*
https://leetcode.com/problems/design-hashset/
705. Design HashSet
Easy
Share
Design a HashSet without using any built-in hash table libraries.
Implement MyHashSet class:
void add(key) Inserts the value key into the HashSet.
bool contains(key) Returns whether the value key exists in the HashSet or not.
void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

Example 1:
Input
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Output
[null, null, null, true, false, null, true, null, false]

Explanation
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // set = [1]
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(1); // return True
myHashSet.contains(3); // return False, (not found)
myHashSet.add(2);      // set = [1, 2]
myHashSet.contains(2); // return True
myHashSet.remove(2);   // set = [1]
myHashSet.contains(2); // return False, (already removed)

Constraints:
0 <= key <= 106
At most 104 calls will be made to add, remove, and contains.
*/

package main

import "fmt"

// type MyHashSet struct {
//     set map[int]bool
// }
// func Constructor() MyHashSet {
//     hs := MyHashSet{
//         set: make(map[int]bool, 1),
//     }
//     return hs
// }
// func (this *MyHashSet) Add(key int)  {
//     this.set[key] = true
// }
// func (this *MyHashSet) Remove(key int)  {
//     this.set[key] = false
// }
// func (this *MyHashSet) Contains(key int) bool {
//     if this.set[key] {
//         return true
//     }
//     return false
// }

type MyHashSet struct {
	set [][1]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		set: make([][1]int, 1000001),
	}
}
func (this *MyHashSet) Add(key int) {
	this.set[key][0] = 1
}
func (this *MyHashSet) Remove(key int) {
	this.set[key][0] = 0
}
func (this *MyHashSet) Contains(key int) bool {
	if this.set[key][0] == 1 {
		return true
	}
	return false
}
func (this *MyHashSet) Print() {
	fmt.Print("[")
	for k, v := range this.set {
		if v[0] != 0 {
			fmt.Printf("%v ", k)
		}
	}
	fmt.Print("]\n")
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	myHashSet := Constructor()
	myHashSet.Add(1) // set = [1]
	myHashSet.Print()
	myHashSet.Add(2) // set = [1, 2]
	myHashSet.Print()
	fmt.Println(myHashSet.Contains(1)) // return True
	fmt.Println(myHashSet.Contains(3)) // return False, (not found)
	myHashSet.Add(2)                   // set = [1, 2]
	myHashSet.Print()
	fmt.Println(myHashSet.Contains(2)) // return True
	myHashSet.Remove(2)                // set = [1]
	myHashSet.Print()
	fmt.Println(myHashSet.Contains(2)) // return False, (already removed)
}
