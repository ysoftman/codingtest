/*
https://leetcode.com/problems/insert-interval/
57. Insert Interval
Medium

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].


Constraints:
0 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 105
intervals is sorted by starti in ascending order.
newInterval.length == 2
0 <= start <= end <= 105
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isOverlapped(a, b, c, d int) bool {
	if a <= c {
		if c <= b {
			return true
		}
	} else {
		if a <= d {
			return true
		}
	}
	return false
}
func insert(intervals [][]int, newInterval []int) [][]int {
	r := [][]int{}
	starti, endi := newInterval[0], newInterval[1]

	if len(intervals) == 0 || intervals[0][0] > endi {
		r = append(r, []int{starti, endi})
		r = append(r, intervals...)
		return r
	}
	if intervals[len(intervals)-1][1] < starti {
		r = append(r, intervals...)
		r = append(r, []int{starti, endi})
		return r
	}

	for i := 0; i < len(intervals); i++ {
		if isOverlapped(starti, endi, intervals[i][0], intervals[i][1]) {
			// merging...
			starti = min(starti, intervals[i][0])
			endi = max(endi, intervals[i][1])
			if i == len(intervals)-1 {
				r = append(r, []int{starti, endi})
				break
			}
		} else {
			if starti > intervals[i][1] {
				r = append(r, intervals[i])
			}
			if endi < intervals[i][0] {
				r = append(r, []int{starti, endi})
				r = append(r, intervals[i:]...)
				break
			}
		}
	}
	return r
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println(insert([][]int{{1, 2}}, []int{3, 4}))
	fmt.Println(insert([][]int{}, []int{5, 7}))
}
