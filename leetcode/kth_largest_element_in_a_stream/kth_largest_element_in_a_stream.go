/*
https://leetcode.com/problems/kth-largest-element-in-a-stream/
703. Kth Largest Element in a Stream
Easy
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

Example 1:
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]
Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

Constraints:
1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
At most 104 calls will be made to add.
It is guaranteed that there will be at least k elements in the array when you search for the kth element.
*/
package main

import (
	"container/heap"
	"fmt"
)

/*
type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

// time limit exceeded
func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	// sort.Ints(this.nums)
	// return this.nums[len(this.nums)-this.k]
	sort.Slice(this.nums, func(a, b int) bool {
		return this.nums[a] > this.nums[b]
	})
	return this.nums[this.k-1]
}
*/

type MinHeap []int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(a, b int) bool { return mh[a] < mh[b] }
func (mh MinHeap) Swap(a, b int)      { mh[a], mh[b] = mh[b], mh[a] }
func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}
func (mh *MinHeap) Pop() interface{} {
	v := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return v
}

type KthLargest struct {
	k       int
	minHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		minHeap: &MinHeap{},
		k:       k,
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	for this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}
