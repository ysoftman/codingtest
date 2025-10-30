/*
https://leetcode.com/problems/range-sum-query-2d-immutable/
304. Range Sum Query 2D - Immutable
Medium

Given a 2D matrix matrix, handle multiple queries of the following type:

Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
Implement the NumMatrix class:

NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
You must design an algorithm where sumRegion works on O(1) time complexity.

Example 1:
Input
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Output
[null, 8, 11, 12]
Explanation
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-104 <= matrix[i][j] <= 104
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
At most 104 calls will be made to sumRegion.
*/

package main

import (
	"fmt"

	"github.com/ysoftman/ysoftmancommon"
)

type NumMatrix struct {
	mat [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	nm := NumMatrix{}
	rows := len(matrix) + 1
	cols := len(matrix[0]) + 1
	nm.mat = make([][]int, rows)
	for i := 0; i < rows; i++ {
		nm.mat[i] = make([]int, cols)
	}

	// for i := 1; i < rows; i++ {
	// 	for j := 1; j < rows; j++ {
	// 		nm.mat[i][j] = matrix[i-1][j-1]
	// 	}
	// }

	// 0,0 각 위치까지의 합을 저장해둔다.
	/*
	 0  0  0  0  0  0           0  0  0  0  0  0
	 0  3  0  1  4  2           0  3  3  4  8 10
	 0  5  6  3  2  1     -->   0  8 14 18 24 27
	 0  1  2  0  1  5           0  9 17 21 28 36
	 0  4  1  0  1  7           0 13 22 26 34 49
	 0  1  0  3  0  5           0 14 23 30 38 58
	*/
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			nm.mat[i][j] = nm.mat[i-1][j] + nm.mat[i][j-1] - nm.mat[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return nm
}

// time complexity : O(1)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// (가장 큰 sum값 bottom,right) - (bottom,left-1 위치 sum) - (top-1,right 위치 sum) + (top-1,left-1 위치 sum)
	return this.mat[row2+1][col2+1] - this.mat[row2+1][col1] - this.mat[row1][col2+1] + this.mat[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
func main() {
	// ["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
	// [[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
	mat := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(mat)
	ysoftmancommon.PrintMatrix(obj.mat)
	fmt.Println(obj.SumRegion(2, 1, 4, 3))
	fmt.Println(obj.SumRegion(1, 1, 2, 2))
	fmt.Println(obj.SumRegion(1, 2, 2, 4))

	// ["NumMatrix","sumRegion","sumRegion","sumRegion"]
	// [[[[-4,-5]]],[0,0,0,0],[0,0,0,1],[0,1,0,1]]
	mat = [][]int{
		{-4, -5},
	}
	obj = Constructor(mat)
	ysoftmancommon.PrintMatrix(obj.mat)
	fmt.Println(obj.SumRegion(0, 0, 0, 0))
	fmt.Println(obj.SumRegion(0, 0, 0, 1))
	fmt.Println(obj.SumRegion(0, 1, 0, 1))
}
