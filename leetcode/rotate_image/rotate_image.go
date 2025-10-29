/*
https://leetcode.com/problems/rotate-image/
48. Rotate Image
Medium

Share
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.


Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:
n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/
package main

import "fmt"

/*
[ 1, 2, 3, 4, 5]
[ 6, 7, 8, 9,10]
[11,12,13,14,15]
[16,17,18,19,20]
[21,22,23,24,25]

[21,16,11, 6, 1]
[22,17,12, 7, 2]
[23,18,13, 8, 3]
[24,19,14, 9, 4]
[25,20,15,10, 5]
*/
func rotate(matrix [][]int) {
	cols := len(matrix[0])
	colstart, colend := 0, cols-1
	rowstart := 0
	rowend := cols - 1
	for colend-colstart > 0 {
		// fmt.Println(rowstart, colstart, rowend, colend)
		j := 1
		for i := colstart + 1; i < colend; i++ {
			temp := matrix[rowstart][i]                                      // 위쪽 값 복사두기
			matrix[rowstart][i] = matrix[rowend-j][colstart]                 // 위쪽 <- 왼쪽
			matrix[rowend-j][colstart] = matrix[rowend][colend-(i-colstart)] // 왼쪽 <- 아래쪽
			matrix[rowend][colend-(i-colstart)] = matrix[rowstart+j][colend] // 아래쪽 <- 오른쪽
			matrix[rowstart+j][colend] = temp                                // 오른쪽 <- 위쪽
			j++
		}
		// 네 꼭지는 2변에 속하는 케이스로 별도 처리
		temp := matrix[rowstart][colstart]
		matrix[rowstart][colstart] = matrix[rowend][colstart]
		matrix[rowend][colstart] = matrix[rowend][colend]
		matrix[rowend][colend] = matrix[rowstart][colend]
		matrix[rowstart][colend] = temp

		colstart++
		colend--
		rowstart++
		rowend--
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	rotate(matrix)
	fmt.Println(matrix)

}
