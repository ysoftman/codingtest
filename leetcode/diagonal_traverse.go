/*
https://leetcode.com/problems/diagonal-traverse/
498. Diagonal Traverse
Medium
Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.

Example 1:
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]

Example 2:
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]

Constraints:
m == mat.length
n == mat[i].length
1 <= m, n <= 104
1 <= m * n <= 104
-105 <= mat[i][j] <= 105
*/

package main

import (
	"fmt"
)

/*
1,2,3
4,5,6
7,8,9

1,2,3,6,9 에 사선(왼쪽 아래)방향의 배열 수집, 총 5개의 배열
0~4 배열을 짝수번째 배열을 reverse 해서 최종 결과로 추가
*/
func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	r := []int{}
	// 사선 방향으로 배열을 수집
	for i := 0; i < m+n-1; i++ {
		temp := []int{}
		row := 0
		col := i
		if col >= n {
			col = n - 1
			row = i - col
		}
		for row < m && col >= 0 {
			temp = append(temp, mat[row][col])
			row++
			col--
		}
		// fmt.Println(temp)
		if i%2 == 0 {
			for i := len(temp) - 1; i >= 0; i-- {
				r = append(r, temp[i])
			}
		} else {
			r = append(r, temp...)
		}
	}
	return r
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
