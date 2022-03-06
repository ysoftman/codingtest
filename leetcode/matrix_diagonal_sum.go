/*
https://leetcode.com/problems/matrix-diagonal-sum/
1572. Matrix Diagonal Sum
Easy
Given a square matrix mat, return the sum of the matrix diagonals.
Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.

Example 1:
Input: mat = [[1,2,3],
              [4,5,6],
              [7,8,9]]
Output: 25
Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
Notice that element mat[1][1] = 5 is counted only once.


Example 2:
Input: mat = [[1,1,1,1],
              [1,1,1,1],
              [1,1,1,1],
              [1,1,1,1]]
Output: 8


Example 3:
Input: mat = [[5]]
Output: 5
*/
package main

import "fmt"

func diagonalSum(mat [][]int) int {
	sum := 0
	if len(mat) == 0 {
		return 0
	}
	if len(mat) == 1 {
		return mat[0][0]
	}

	for i := 0; i < len(mat); i++ {
		sum += mat[i][i] + mat[i][len(mat)-1-i]
	}
	if len(mat)%2 == 1 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}
	return sum
}

func main() {
	fmt.Println(diagonalSum([][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))
	fmt.Println(diagonalSum([][]int{{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1}}))
	fmt.Println(diagonalSum([][]int{{5}}))
}
