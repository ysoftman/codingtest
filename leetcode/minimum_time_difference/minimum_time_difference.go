/*
https://leetcode.com/problems/minimum-time-difference
539. Minimum Time Difference
Medium

Given a list of 24-hour clock time points in "HH:MM" format, return the minimum minutes difference between any two time-points in the list.

Example 1:
Input: timePoints = ["23:59","00:00"]
Output: 1

Example 2:
Input: timePoints = ["00:00","23:59","00:00"]
Output: 0

Constraints:
2 <= timePoints.length <= 2 * 104
timePoints[i] is in the format "HH:MM".
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func HHMMtoMinutes(s string) int {
	t := strings.Split(s, ":")
	h, _ := strconv.Atoi(t[0])
	m, _ := strconv.Atoi(t[1])
	return (h * 60) + m
}

func findMinDifference(timePoints []string) int {
	delta1 := []int{}
	delta2 := []int{}
	for _, tp := range timePoints {
		m := HHMMtoMinutes(tp)
		delta1 = append(delta1, m)
		if m < 12*60 {
			m += (24 * 60)
		}
		delta2 = append(delta2, m)
	}
	min := 1<<32 - 1
	sort.Slice(delta1, func(a, b int) bool {
		return delta1[a] > delta1[b]
	})
	sort.Slice(delta2, func(a, b int) bool {
		return delta2[a] > delta2[b]
	})
	// fmt.Println(delta1)
	// fmt.Println(delta2)
	for i := 0; i < len(delta1)-1; i++ {
		d := abs(delta1[i] - delta1[i+1])
		if min > d {
			min = d
		}
	}
	for i := 0; i < len(delta2)-1; i++ {
		d := abs(delta2[i] - delta2[i+1])
		if min > d {
			min = d
		}
	}
	return min
}

func main() {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))
	fmt.Println(findMinDifference([]string{"23:59", "01:58"}))
	fmt.Println(findMinDifference([]string{"11:59", "01:58"}))
	fmt.Println(findMinDifference([]string{"23:59", "00:00", "02:19", "12:23", "11:39", "15:30", "20:08"}))
}
