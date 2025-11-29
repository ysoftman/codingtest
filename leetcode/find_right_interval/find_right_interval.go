/*
https://leetcode.com/problems/find-right-interval
436. Find Right Interval
Medium
You are given an array of intervals, where intervals[i] = [starti, endi] and each starti is unique.

The right interval for an interval i is an interval j such that startj >= endi and startj is minimized. Note that i may equal j.

Return an array of right interval indices for each interval i. If no right interval exists for interval i, then put -1 at index i.

Example 1:
Input: intervals = [[1,2]]
Output: [-1]
Explanation: There is only one interval in the collection, so it outputs -1.

Example 2:
Input: intervals = [[3,4],[2,3],[1,2]]
Output: [-1,0,1]
Explanation: There is no right interval for [3,4].
The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start that is >= end1 = 3.
The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start that is >= end2 = 2.

Example 3:
Input: intervals = [[1,4],[2,3],[3,4]]
Output: [-1,2,-1]
Explanation: There is no right interval for [1,4] and [3,4].
The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start that is >= end1 = 3.

Constraints:
1 <= intervals.length <= 2 * 104
intervals[i].length == 2
-106 <= starti <= endi <= 106
The start point of each interval is unique.
*/
package main

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	m := make(map[int]int)
	for i := range intervals {
		m[intervals[i][0]] = i
	}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	// fmt.Println(intervals)

	r := make([]int, len(intervals))
	for i := 0; i < len(intervals); i++ {
		orgIdx, _ := m[intervals[i][0]]
		if intervals[i][0] == intervals[i][1] {
			r[orgIdx] = orgIdx
			continue
		}
		r[orgIdx] = -1
		if i == len(intervals)-1 {
			continue
		}
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] <= intervals[j][0] {
				intervalIdx, _ := m[intervals[j][0]]
				r[orgIdx] = intervalIdx
				break
			}
		}
	}
	return r
}

func main() {
	fmt.Println(findRightInterval([][]int{{1, 2}}))
	fmt.Println(findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	fmt.Println(findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	fmt.Println(findRightInterval([][]int{{1, 12}, {2, 9}, {3, 10}, {13, 14}, {15, 16}, {16, 17}}))
	fmt.Println(findRightInterval([][]int{{1, 1}, {3, 4}}))
	fmt.Println(findRightInterval([][]int{{4, 4}}))
}
