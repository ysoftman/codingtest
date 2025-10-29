/*
https://leetcode.com/problems/search-a-2d-matrix/
74. Search a 2D Matrix
Medium
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.

Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	i, j := 0, rows-1
	midrow := 0
	for i <= j {
		midrow = i + (j-i)/2
		if matrix[midrow][0] <= target && target <= matrix[midrow][cols-1] {
			break
		}
		if matrix[midrow][0] < target {
			i = midrow + 1
		} else {
			j = midrow - 1
		}
	}
	if midrow >= rows || midrow < 0 {
		return false
	}

	// fmt.Println("midrow:", midrow)
	i, j = 0, cols-1
	midcol := 0
	for i <= j {
		midcol = i + (j-i)/2
		if matrix[midrow][midcol] == target {
			return true
		}
		if matrix[midrow][midcol] < target {
			i = midcol + 1
		} else {
			j = midcol - 1
		}

	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 10))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 1))
}
