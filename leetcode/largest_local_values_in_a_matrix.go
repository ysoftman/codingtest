/*
2373. Largest Local Values in a Matrix
Easy
You are given an n x n integer matrix grid.
Generate an integer matrix maxLocal of size (n - 2) x (n - 2) such that:
maxLocal[i][j] is equal to the largest value of the 3 x 3 matrix in grid centered around row i + 1 and column j + 1.
In other words, we want to find the largest value in every contiguous 3 x 3 matrix in grid.

Return the generated matrix.

Example 1:
Input: grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
Output: [[9,9],[8,6]]
Explanation: The diagram above shows the original matrix and the generated matrix.
Notice that each value in the generated matrix corresponds to the largest value of a contiguous 3 x 3 matrix in grid.

Example 2:
Input: grid = [[1,1,1,1,1],[1,1,1,1,1],[1,1,2,1,1],[1,1,1,1,1],[1,1,1,1,1]]
Output: [[2,2,2],[2,2,2],[2,2,2]]
Explanation: Notice that the 2 is contained within every contiguous 3 x 3 matrix in grid.

Constraints:
n == grid.length == grid[i].length
3 <= n <= 100
1 <= grid[i][j] <= 100
*/
package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, len(grid)-2)
	for i := 0; i < len(grid)-2; i++ {
		result[i] = make([]int, len(grid[0])-2)
	}
	for i := 0; i < len(grid)-2; i++ {
		row := []int{}
		for j := 0; j < len(grid[i])-2; j++ {
			// 3x3
			temp := 0
			for a := 0; a < 3; a++ {
				for b := 0; b < 3; b++ {
					if grid[i+a][j+b] > temp {
						temp = grid[i+a][j+b]
					}
				}
			}
			row = append(row, temp)
		}
		result[i] = row
	}
	return result
}

func main() {
	fmt.Println(largestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
	fmt.Println(largestLocal([][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}))
}
