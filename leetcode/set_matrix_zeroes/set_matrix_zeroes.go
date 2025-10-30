/*
https://leetcode.com/problems/set-matrix-zeroes/
73. Set Matrix Zeroes
Medium

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1
*/

package main

import "github.com/ysoftman/ysoftmancommon"

// O(mn) time complexity
// O(m+n) space complexity
func setZeroes2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// O(mn) time complexity
// O(1) space complexity
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	firstColIsZero := false
	for i := 0; i < m; i++ {
		// 마킹된 0 이 아니고 원래 첫번째 컬럼중에 0 이 있는 경우
		if matrix[i][0] == 0 {
			firstColIsZero = true
		}
		// [rows][0] 은 firstColIsZero 로 판단하기 때문에 1번째 col 부터 파악
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// marking(set first of cols and rows to 0)
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// [0][0] 의 0 은 [0][cols] 판단용으로 사용한다.([rows][0] 에 판단에 사용하지 못함)
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	// 마킹된 0 이 아니고 원래 첫번째 컬럼중에 0 이 있는 경우
	if firstColIsZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	ysoftmancommon.PrintMatrix(matrix)
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	ysoftmancommon.PrintMatrix(matrix)
	matrix = [][]int{{0, 1, 0, 3, 0}, {3, 7, 4, 5, 2}, {0, 3, 1, 9, 5}, {0, 12, 6, 2, 1}}
	setZeroes(matrix)
	ysoftmancommon.PrintMatrix(matrix)
}
