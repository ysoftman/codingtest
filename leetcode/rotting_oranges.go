/*
https://leetcode.com/problems/rotting-oranges/
994. Rotting Oranges
Medium
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

Example 1:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Example 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
*/
package main

import "fmt"

func checkEdge(i, j, Y, X int) bool {
	if i >= 0 && i < Y && j >= 0 && j < X {
		return true
	}
	return false
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	q := make([][]int, 0)
	// queue all rotten_orange(2)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}

	cnt := 0
	// BFS
	for len(q) > 0 {
		curLen := len(q)
		// step -> flush current queue
		for idx := 0; idx < curLen; idx++ {
			top := q[idx]
			i, j := top[0], top[1]
			if checkEdge(i, j, rows, cols) && grid[i][j] != 2 {
				continue
			}

			if checkEdge(i-1, j, rows, cols) && grid[i-1][j] == 1 {
				grid[i-1][j] = 2
				q = append(q, []int{i - 1, j})
			}
			if checkEdge(i, j-1, rows, cols) && grid[i][j-1] == 1 {
				grid[i][j-1] = 2
				q = append(q, []int{i, j - 1})
			}
			if checkEdge(i+1, j, rows, cols) && grid[i+1][j] == 1 {
				grid[i+1][j] = 2
				q = append(q, []int{i + 1, j})
			}
			if checkEdge(i, j+1, rows, cols) && grid[i][j+1] == 1 {
				grid[i][j+1] = 2
				q = append(q, []int{i, j + 1})
			}
		}

		if len(q) > curLen {
			cnt++
		}
		q = q[curLen:]
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return cnt
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}
