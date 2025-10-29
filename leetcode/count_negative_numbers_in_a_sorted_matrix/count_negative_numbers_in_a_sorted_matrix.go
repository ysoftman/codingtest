/*
https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
1351. Count Negative Numbers in a Sorted Matrix
Easy
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

Example 1:
Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.

Example 2:
Input: grid = [[3,2],[1,0]]
Output: 0

Constraints:
m == grid.length
n == grid[i].length
1 <= m, n <= 100
-100 <= grid[i][j] <= 100

Follow up: Could you find an O(n + m) solution?
*/
package main

import "fmt"

// brute-force : O (m*n)
func countNegatives2(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				cnt++
			}
		}
	}
	return cnt
}

/*
[ 4, 3, 2,-1] => binary search
[ 3, 2, 1,-1] => binary search
[ 1, 1,-1,-2] => binary search
[-1,-1,-2,-3] => binary search
*/
// binary search : O(mlogn)
func countNegatives1(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		left, right := 0, len(grid[i])-1
		mid := 0
		for left <= right {
			mid = (right-left)/2 + left
			if grid[i][mid] < 0 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if len(grid[i]) > left && grid[i][left] >= 0 {
			left++
		}
		// fmt.Println("left:", left)
		cnt += len(grid[i]) - (left)
	}
	// fmt.Println("-----")
	return cnt
}

/*
[ 4, 3, 2,-1] => col:3
[ 3, 2, 1,-1] => col:3
[ 1, 1,-1,-2] => col:3->2
[-1,-1,-2,-3] => col:2->1->0
*/
// O(m+n)
func countNegatives(grid [][]int) int {
	cnt := 0
	row := 0
	col := len(grid[0]) - 1
	// 현재 row 에서 찾은 col 이 다음 row 에서 시작 col이 된다.
	for row < len(grid) && col < len(grid[row]) {
		if col >= 0 && grid[row][col] < 0 {
			col--
		} else {
			cnt += len(grid[row]) - (col + 1)
			row++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
}
