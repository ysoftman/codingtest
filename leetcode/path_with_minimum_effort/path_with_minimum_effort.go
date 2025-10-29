/*
https://leetcode.com/problems/path-with-minimum-effort/
1631. Path With Minimum Effort
Medium
You are a hiker preparing for an upcoming hike. You are given heights, a 2D array of size rows x columns, where heights[row][col] represents the height of cell (row, col). You are situated in the top-left cell, (0, 0), and you hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e., 0-indexed). You can move up, down, left, or right, and you wish to find a route that requires the minimum effort.

A route's effort is the maximum absolute difference in heights between two consecutive cells of the route.

Return the minimum effort required to travel from the top-left cell to the bottom-right cell.

Example 1:
Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
Output: 2
Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.

Example 2:
Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
Output: 1
Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].

Example 3:
Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
Output: 0
Explanation: This route does not require any effort.

Constraints:
rows == heights.length
columns == heights[i].length
1 <= rows, columns <= 100
1 <= heights[i][j] <= 106
*/

package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func findPath(heights [][]int, from []int, v map[[2]int]struct{}, max int) bool {
	y, x := from[0], from[1]
	if _, ok := v[[2]int{y, x}]; ok {
		return false
	}
	v[[2]int{y, x}] = struct{}{}
	gx := len(heights[0]) - 1
	gy := len(heights) - 1
	if x == gx && y == gy {
		return true
	}
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, dir := range directions {
		dy, dx := dir[0], dir[1]
		if 0 <= x+dx && x+dx < len(heights[0]) && 0 <= y+dy && y+dy < len(heights) {
			if abs(heights[y+dy][x+dx]-heights[y][x]) <= max {
				if findPath(heights, []int{y + dy, x + dx}, v, max) {
					return true
				}
			}
		}
	}
	return false
}

// dijkstra 경로 찾기 알고리즘 or DFS+BinarySearch 사용해야 됨
// DFS+BinarySearch 사용 참고: https://leetcode.com/problems/path-with-minimum-effort/discuss/1123206/Go%3A-DFS(%2Bmap)-Binary-Search
func minimumEffortPath(heights [][]int) int {
	mid, left, right := 0, 0, 1<<31-1
	for left < right {
		mid = (left + right) / 2
		if findPath(heights, []int{0, 0}, make(map[[2]int]struct{}, 0), mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return (left + right) / 2
}

func main() {
	fmt.Println(minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
	fmt.Println(minimumEffortPath([][]int{{1, 10, 6, 7, 9, 10, 4, 9}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
	fmt.Println(minimumEffortPath([][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}))
}
