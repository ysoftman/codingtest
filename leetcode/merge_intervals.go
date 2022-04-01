/*
https://leetcode.com/problems/merge-intervals/
56. Merge Intervals
Medium

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

Constraints:
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

package main

import (
	"fmt"
	"sort"
)

func isOverlapped(a, b, c, d int) (bool, []int) {
	i1, i2, j1, j2 := a, b, c, d
	if a > b {
		i1, i2, j1, j2 = c, d, a, b
	}

	merge := make([]int, 2)
	if i1 <= j1 && j1 <= i2 {
		merge[0] = i1
		if j2 > i2 {
			merge[1] = j2 // ex) [2,5] [4,6]
		} else {
			merge[1] = i2 // ex) [2,7] [4,6]
		}
		return true, merge
	}
	return false, merge
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	r := [][]int{}
	cur := []int{intervals[0][0], intervals[0][1]}
	for _, v := range intervals {
		if yes, merge := isOverlapped(cur[0], cur[1], v[0], v[1]); yes {
			cur = merge
		} else {
			r = append(r, []int{cur[0], cur[1]})
			cur[0] = v[0]
			cur[1] = v[1]
		}
	}
	r = append(r, cur)
	return r

}
func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 5}, {2, 3}, {4, 6}, {8, 10}, {2, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
