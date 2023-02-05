/*
https://leetcode.com/problems/island-perimeter/
463. Island Perimeter
Easy
You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Example 1:
Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
Output: 16
Explanation: The perimeter is the 16 yellow stripes in the image above.

Example 2:
Input: grid = [[1]]
Output: 4

Example 3:
Input: grid = [[1,0]]
Output: 4

Constraints:

row == grid.length
col == grid[i].length
1 <= row, col <= 100
grid[i][j] is 0 or 1.
There is exactly one island in grid.
*/
package main

import "fmt"

// island = cell(1x1)들이 가로,세로 연결됨
// island 1개만 반드시 존재, lake (island 안에 water)는 없음
// island 의 둘레(perimeter) 구하기
func islandPerimeter(grid [][]int) int {
	perimeter := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				continue
			}
			// cell=1 이면 이건 island 의 일부이다.
			perimeter += 4
			// 위쪽 cell==1(island의 일부) 라면
			// 위쪽 아래쪽변과 현재가 위쪽변이 중첩, 각각 1씩 => 2개를 빼야 한다.
			if y-1 >= 0 && grid[y-1][x] == 1 {
				perimeter -= 2
			}
			// 왼쪽 cell==1(island의 일부) 라면
			// 왼쪽 오른쪽변과 현재 왼쪽변이 중첩, 각각 1씩 => 2개를 빼야 한다.
			if x-1 >= 0 && grid[y][x-1] == 1 {
				perimeter -= 2
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1}}))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
}
