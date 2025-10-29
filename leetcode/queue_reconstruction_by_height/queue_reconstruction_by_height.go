/*
https://leetcode.com/problems/queue-reconstruction-by-height/
406. Queue Reconstruction by Height
Medium
You are given an array of people, people, which are the attributes of some people in a queue (not necessarily in order). Each people[i] = [hi, ki] represents the ith person of height hi with exactly ki other people in front who have a height greater than or equal to hi.

Reconstruct and return the queue that is represented by the input array people. The returned queue should be formatted as an array queue, where queue[j] = [hj, kj] is the attributes of the jth person in the queue (queue[0] is the person at the front of the queue).

Example 1:
Input: people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
Output: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
Explanation:
Person 0 has height 5 with no other people taller or the same height in front.
Person 1 has height 7 with no other people taller or the same height in front.
Person 2 has height 5 with two persons taller or the same height in front, which is person 0 and 1.
Person 3 has height 6 with one person taller or the same height in front, which is person 1.
Person 4 has height 4 with four people taller or the same height in front, which are people 0, 1, 2, and 3.
Person 5 has height 7 with one person taller or the same height in front, which is person 1.
Hence [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] is the reconstructed queue.

Example 2:
Input: people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
Output: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]

Constraints:
1 <= people.length <= 2000
0 <= hi <= 106
0 <= ki < people.length
It is guaranteed that the queue can be reconstructed.
*/
package main

import (
	"fmt"
	"sort"
)

// time complexity: O(nlogn)
func reconstructQueue(people [][]int) [][]int {
	r := [][]int{}
	// hi (키) 큰순으로 정렬, 같은 키는 ki (앞에 키(hi)가 같거나 큰 사람 수)이 작은 순으로 정렬
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] > people[b][0] {
			return true
		}
		if people[a][0] == people[b][0] {
			return people[a][1] < people[b][1]
		}
		return false
	})
	// fmt.Println("sort descending:", people)
	// ki (앞에 키(hi)가 같거나 큰 사람 수) 위치에 insert
	for _, v := range people {
		idx := v[1]
		r = append(r, []int{0, 0})
		copy(r[idx+1:], r[idx:])
		// r[idx] = []int{v[0], v[1]}
		r[idx] = v
		// fmt.Printf("---insert idx(%v) %v\n", idx, r)
	}
	return r
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	fmt.Println(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}))
	fmt.Println(reconstructQueue([][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}}))
}
