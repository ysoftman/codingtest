/*
https://leetcode.com/problems/last-stone-weight
1046. Last Stone Weight
Easy
Topics
premium lock icon
Companies
Hint
You are given an array of integers stones where stones[i] is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose the heaviest two stones and smash them together. Suppose the heaviest two stones have weights x and y with x <= y. The result of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no stones left, return 0.

Example 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

Example 2:
Input: stones = [1]
Output: 1

Constraints:
1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// sort 사용
func lastStoneWeight2(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		n := len(stones)
		diff := stones[n-1] - stones[n-2]
		stones = stones[:n-2]
		if diff > 0 {
			stones = append(stones, diff)
		}
		sort.Ints(stones)
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

type MaxHeap []int

func (mh MaxHeap) Len() int           { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool { return mh[i] > mh[j] }
func (mh MaxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

// priority queue(maxheap) 사용
func lastStoneWeight(stones []int) int {
	mh := &MaxHeap{}
	heap.Init(mh)
	for _, v := range stones {
		heap.Push(mh, v)
	}
	for mh.Len() > 1 {
		v1 := heap.Pop(mh).(int)
		v2 := heap.Pop(mh).(int)
		if v1-v2 > 0 {
			heap.Push(mh, v1-v2)
		}
	}
	if mh.Len() == 1 {
		return heap.Pop(mh).(int)
	}
	return 0
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{2, 2}))
	fmt.Println(lastStoneWeight([]int{1, 3}))
}
