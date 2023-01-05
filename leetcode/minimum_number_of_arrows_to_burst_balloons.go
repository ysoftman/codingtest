/*
https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
452. Minimum Number of Arrows to Burst Balloons
Medium
There are some spherical balloons taped onto a flat wall that represents the XY-plane. The balloons are represented as a 2D integer array points where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between xstart and xend. You do not know the exact y-coordinates of the balloons.

Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis. A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend. There is no limit to the number of arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path.

Given the array points, return the minimum number of arrows that must be shot to burst all balloons.

Example 1:
Input: points = [[10,16],[2,8],[1,6],[7,12]]
Output: 2
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at x = 6, bursting the balloons [2,8] and [1,6].
- Shoot an arrow at x = 11, bursting the balloons [10,16] and [7,12].

Example 2:
Input: points = [[1,2],[3,4],[5,6],[7,8]]
Output: 4
Explanation: One arrow needs to be shot for each balloon for a total of 4 arrows.

Example 3:
Input: points = [[1,2],[2,3],[3,4],[4,5]]
Output: 2
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at x = 2, bursting the balloons [1,2] and [2,3].
- Shoot an arrow at x = 4, bursting the balloons [3,4] and [4,5].

Constraints:
1 <= points.length <= 105
points[i].length == 2
-231 <= xstart < xend <= 231 - 1
*/
package main

import (
	"fmt"
	"sort"
)

/*
[[10,16],[2,8],[1,6],[7,12]]
=>  [[1,6],[2,8],[7,12],[10,16]] Xstart 기준으로 정렬
XEnd 가 next point(Xstart, Xend) 범위에 있으면(겹치면)같은 그룹으로,
next porint가 현재 그룹으로 묶을 수 없으면 새그룹 시작
그룹 개수 = 화살 개수
*/
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(a, b int) bool {
		return points[a][0] < points[b][0]
	})
	// fmt.Println(points)
	arrows := 1
	xEnd := points[0][1]
	for _, p := range points {
		// 같은 그룹(범위가 겹침)
		if xEnd >= p[0] {
			// 그룹내의 모든 좌표는 겹쳐야해서 xEnd 는 최소 범위로 해야 한다.
			if xEnd > p[1] {
				xEnd = p[1]
			}
			continue
		} else { // 새그룹 시작
			arrows++
			xEnd = p[1]
		}
	}
	return arrows
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println(findMinArrowShots([][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}))
}
