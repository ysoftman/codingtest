/*
https://leetcode.com/problems/01-matrix/
542. 01 Matrix
Medium
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

Example 1:
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]
*/
package main

import "fmt"

// DFS, but time limit exceeded!!!
func findNearestZero(mat [][]int, y, x int) int {
	if mat[y][x] == 0 {
		return 0
	}
	lenY := len(mat)
	lenX := len(mat[0])
	min := 1<<31 - 1
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for _, v := range dirs {
		if y+v[0] >= 0 && y+v[0] < lenY && x+v[1] >= 0 && x+v[1] < lenX {
			if mat[y+v[0]][x+v[1]] == 1 {
				mat[y+v[0]][x+v[1]] = 2 // marked traversal
				temp := findNearestZero(mat, y+v[0], x+v[1]) + 1
				if min > temp {
					min = temp
				}
				mat[y+v[0]][x+v[1]] = 1 // restore for next iteration
			} else if mat[y+v[0]][x+v[1]] == 0 {
				return 1
			}
		}
	}
	return min
}
func updateMatrix2(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		result[i] = make([]int, len(mat[i]))
	}
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			result[y][x] = findNearestZero(mat, y, x)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// dynamic programming
func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		result[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			result[i][j] = 1<<31 - 2
		}
	}
	lenY := len(mat)
	lenX := len(mat[0])

	// check left and up
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if mat[y][x] == 0 {
				result[y][x] = 0
			} else {
				if y-1 >= 0 {
					result[y][x] = min(result[y][x], result[y-1][x]+1)
				}
				if x-1 >= 0 {
					result[y][x] = min(result[y][x], result[y][x-1]+1)
				}
			}
		}
	}
	// check right and bottom
	for y := lenY - 1; y >= 0; y-- {
		for x := lenX - 1; x >= 0; x-- {
			if mat[y][x] == 0 {
				result[y][x] = 0
			} else {
				if y+1 < lenY {
					result[y][x] = min(result[y][x], result[y+1][x]+1)
				}
				if x+1 < lenX {
					result[y][x] = min(result[y][x], result[y][x+1]+1)
				}
			}
		}
	}
	return result
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(updateMatrix(mat))
	mat = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(mat))

	mat = [][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	fmt.Println(updateMatrix(mat))
	mat = [][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 1, 1, 1, 1},
		{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 1, 1, 0, 0, 1},
		{1, 1, 1, 1, 1, 0, 0, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}
	fmt.Println(updateMatrix(mat))
}
