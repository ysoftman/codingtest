/*
https://leetcode.com/problems/lru-cache/
146. LRU Cache
Medium
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.

Example 1:
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4


Constraints:
1 <= capacity <= 3000
0 <= key <= 104
0 <= value <= 105
At most 2 * 105 calls will be made to get and put.
*/

package main

import "fmt"

type Node struct {
	key  int
	val  int
	pre  *Node
	next *Node
}

type LRUCache struct {
	capacity int
	m        map[int]*Node // O(1) 검색을 위해
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node, 0),
		tail:     &Node{},
		head:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if _, exist := this.m[key]; !exist {
		return -1
	}
	node := this.m[key]
	// 이미 맨뒤에 있는 경우
	if node == this.tail.pre {
		// for debugging...
		// fmt.Println("[get-case1]", key, this.m)
		// printNode(this.head)
		return this.m[key].val
	}
	// 조회된 노드는 맨뒤로 이동
	node.pre.next = node.next
	node.next.pre = node.pre
	this.tail.pre.next = node
	node.pre = this.tail.pre
	this.tail.pre = node
	node.next = this.tail
	// for debugging...
	// fmt.Println("[get-case2]", key, this.m)
	// printNode(this.head)
	return this.m[key].val
}

func (this *LRUCache) Put(key int, value int) {
	// 이미 있는 값이면 get 에서 노드를 맨뒤로 이동
	if this.Get(key) != -1 {
		// 키는 같지만 값이 변경될 수 있음
		this.m[key].val = value
		return
	}
	// 새로 추가 되는경우 맨뒤에 추가
	newnode := &Node{
		key: key,
		val: value,
	}
	this.tail.pre.next = newnode
	newnode.pre = this.tail.pre
	this.tail.pre = newnode
	newnode.next = this.tail

	this.m[key] = newnode

	// fmt.Println("[put]", key, value, this.m)
	// capacity 보다 크면 맨 앞노드 삭제
	if len(this.m) > this.capacity {
		temp := this.head.next
		this.head.next = temp.next
		temp.next.pre = this.head
		// map 에서도 해당 키 삭제
		delete(this.m, temp.key)
		// fmt.Println("[put] delete", temp.key, temp.val)
	}
	// for debugging...
	// printNode(this.head)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func printNode(head *Node) {
	for head != nil {
		fmt.Printf("(%v:%v) ", head.key, head.val)
		head = head.next
	}
	fmt.Println()
}

func main() {
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

	fmt.Println("-----")
	// ["LRUCache","put","get","put","get","get"]
	// [[1],[2,1],[2],[3,2],[2],[3]]
	obj = Constructor(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Put(3, 2)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))

	fmt.Println("-----")
	// ["LRUCache","put","put","get","put","put","get"]
	// [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	obj = Constructor(2)
	obj.Put(2, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(2))
	obj.Put(1, 1)
	obj.Put(4, 2)
	fmt.Println(obj.Get(2))

	fmt.Println("-----")
	// ["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	// [[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	obj = Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	obj.Put(5, 5)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(5))
}
