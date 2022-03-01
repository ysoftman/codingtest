/*
https://leetcode.com/problems/max-area-of-island/

695. Max Area of Island
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.
The area of an island is the number of cells with a value 1 in the island.
Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:
Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.

Example 2:
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
*/
package main

import "fmt"

func areaCC(grid *[][]int, y, x int, area *int) {
	lenY := len(*grid)
	lenX := len((*grid)[0])

	if y-1 >= 0 && (*grid)[y-1][x] == 1 {
		(*grid)[y-1][x] = 2
		*area++
		areaCC(grid, y-1, x, area)
	}
	if x-1 >= 0 && (*grid)[y][x-1] == 1 {
		(*grid)[y][x-1] = 2
		*area++
		areaCC(grid, y, x-1, area)
	}
	if y+1 < lenY && (*grid)[y+1][x] == 1 {
		(*grid)[y+1][x] = 2
		*area++
		areaCC(grid, y+1, x, area)
	}
	if x+1 < lenX && (*grid)[y][x+1] == 1 {
		(*grid)[y][x+1] = 2
		*area++
		areaCC(grid, y, x+1, area)
	}

}
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				grid[y][x] = 2
				area := 1
				areaCC(&grid, y, x, &area)
				if max < area {
					max = area
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
	fmt.Println(maxAreaOfIsland([][]int{{0, 0, 0, 0, 0, 0, 0, 0}}))
}
