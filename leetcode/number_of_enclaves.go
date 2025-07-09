/*
https://leetcode.com/problems/number-of-enclaves/
1020. Number of Enclaves
Medium
You are given an m x n binary matrix grid, where 0 represents a sea cell and 1 represents a land cell.
A move consists of walking from one land cell to another adjacent (4-directionally) land cell or walking off the boundary of the grid.
Return the number of land cells in grid for which we cannot walk off the boundary of the grid in any number of moves.

Example 1:
Input: grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
Output: 3
Explanation: There are three 1s that are enclosed by 0s, and one 1 that is not enclosed because its on the boundary.

Example 2:
Input: grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
Output: 0
Explanation: All 1s are either on the boundary or can reach the boundary.

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] is either 0 or 1.
*/

package main

import "fmt"

func markVisited(grid *[][]int, i, j int) {
	if i < 0 || i == len(*grid) || j < 0 || j == len((*grid)[i]) {
		return
	}
	if (*grid)[i][j] == 0 {
		return
	}
	(*grid)[i][j] = 0
	// recursive mark horizontally and vertically
	markVisited(grid, i-1, j)
	markVisited(grid, i+1, j)
	markVisited(grid, i, j-1)
	markVisited(grid, i, j+1)
}

func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// change edge(boundary) cell to 0
			if i*j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				markVisited(&grid, i, j)
			}
		}
	}

	cnt := 0
	// count remain 1 cell
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// for debugging...
			fmt.Printf("%v ", grid[i][j])
			if grid[i][j] == 1 {
				cnt++
			}
		}
		// for debugging...
		fmt.Println()
	}
	return cnt
}

func main() {
	fmt.Println(numEnclaves([][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}))
	fmt.Println(numEnclaves([][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}))
	fmt.Println(numEnclaves([][]int{
		{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
		{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
	}))
}
