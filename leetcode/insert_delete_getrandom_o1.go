/*
https://leetcode.com/problems/insert-delete-getrandom-o1/
380. Insert Delete GetRandom O(1)
Medium

Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.

Example 1:
Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


Constraints:
-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m   map[int]int // insert,remove time complexity: O(1) 를 위해
	arr []int       // 랜덤 조회시 time complexity: O(1) 를 위해
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		m:   make(map[int]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.m[val]; !exist {
		this.arr = append(this.arr, val)
		// array 위치(index) 를 저장
		this.m[val] = len(this.arr) - 1
		// fmt.Println("-insert-index-", this.m[val])
		// fmt.Println("---", this.arr)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.m[val]; !exist {
		return false
	}
	index := this.m[val]
	// fmt.Println("-remove-index-", index)
	// index 로 array 원소 삭제
	// index 원소를 삭제하면 index 뒤 원소들의 위치가 바뀌기 때문에
	// 삭제학 원소를 array 의 마지막 원소와 swap 하여 삭제할 원소가 array 의 마지막으로 옮긴 후에 마지막 원소를 삭제해야 한다.
	lastval := this.arr[len(this.arr)-1]
	this.arr[index], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[index]
	this.m[lastval] = index

	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, val)
	// fmt.Println("---", this.arr)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	// 값은 array 에서 랜덤하게 파악
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(2))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())

	fmt.Println("-----")
	rs = Constructor()
	fmt.Println(rs.Insert(0))
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(0))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.GetRandom())
}
