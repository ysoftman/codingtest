/*
https://leetcode.com/problems/toeplitz-matrix/
766. Toeplitz Matrix
Easy

Given an m x n matrix, return true if the matrix is Toeplitz. Otherwise, return false.
A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same elements.

Example 1:
Input: matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
Output: true
Explanation:
In the above grid, the diagonals are:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
In each diagonal all elements are the same, so the answer is True.

Example 2:
Input: matrix = [[1,2],[2,2]]
Output: false
Explanation:
The diagonal "[1, 2]" has different elements.

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 20
0 <= matrix[i][j] <= 99

Follow up:
What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?
What if the matrix is so large that you can only load up a partial row into the memory at once?
*/
package main

import "fmt"

// x=0, y=0 일때 대각선 확인하는 방법
func isToeplitzMatrix2(matrix [][]int) bool {
	y := len(matrix)
	x := len(matrix[0])
	for i := 0; i < x; i++ {
		cur := matrix[0][i]
		a := 0
		b := i
		for a < y && b < x {
			if matrix[a][b] != cur {
				return false
			}
			a++
			b++
		}
	}
	for i := 0; i < y; i++ {
		cur := matrix[i][0]
		a := i
		b := 0
		for a < y && b < x {
			if matrix[a][b] != cur {
				return false
			}
			a++
			b++
		}
	}
	return true
}

// 좀더 간단하게 전체 스캐하면서 대각선위 값을 확인하는 방법
func isToeplitzMatrix(matrix [][]int) bool {
	for y := 1; y < len(matrix); y++ {
		for x := 1; x < len(matrix[0]); x++ {
			if matrix[y-1][x-1] != matrix[y][x] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}
