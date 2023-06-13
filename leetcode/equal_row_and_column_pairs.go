/*
2352. Equal Row and Column Pairs
Medium
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.
A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

Example 1:
Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
Output: 1
Explanation: There is 1 equal row and column pair:
- (Row 2, Column 1): [2,7,7]

Example 2:
Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
Output: 3
Explanation: There are 3 equal row and column pairs:
- (Row 0, Column 0): [3,1,2,2]
- (Row 2, Column 2): [2,4,2,2]
- (Row 3, Column 2): [2,4,2,2]

Constraints:
n == grid.length == grid[i].length
1 <= n <= 200
1 <= grid[i][j] <= 105
*/
package main

import "fmt"

func equalPairs(grid [][]int) int {
	row := make(map[int]string)
	col := make(map[int]string)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 11, 1 등을 구분하기 위해 - 구분자 추가
			row[i] += fmt.Sprintf("%d-", grid[i][j])
			col[j] += fmt.Sprintf("%d-", grid[i][j])
		}
	}
	// fmt.Println(row)
	// fmt.Println(col)
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if row[i] == col[j] {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(equalPairs([][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}))
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
	fmt.Println(equalPairs([][]int{{11, 1}, {1, 11}}))
}
