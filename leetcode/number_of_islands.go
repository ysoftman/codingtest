/*
https://leetcode.com/problems/number-of-islands/
200. Number of Islands
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1


Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3

*/

package main

import "fmt"

func markVisited(grid *[][]byte, i, j int) {

	if i < 0 || i >= len(*grid) || j < 0 || j >= len((*grid)[i]) {
		return
	}
	if (*grid)[i][j] == '0' || (*grid)[i][j] == '2' {
		return
	}

	(*grid)[i][j] = '2'
	// recursive mark horizontallly and vertically
	markVisited(grid, i-1, j)
	markVisited(grid, i+1, j)
	markVisited(grid, i, j-1)
	markVisited(grid, i, j+1)
}

func numIslands(grid [][]byte) int {
	countIsland := 0

	// O(mn)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// first visit
			if grid[i][j] == '1' {
				countIsland++
				markVisited(&grid, i, j)
			}

		}
	}
	// for i:=0; i<len(grid); i++ {
	//     for j:=0; j<len(grid[i]); j++ {
	//         fmt.Printf("%s ", string(grid[i][j]))
	//     }
	//     fmt.Println()
	// }
	return countIsland
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}
