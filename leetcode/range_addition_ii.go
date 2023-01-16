/*
https://leetcode.com/problems/range-addition-ii/
598. Range Addition II
Easy
You are given an m x n matrix M initialized with all 0's and an array of operations ops, where ops[i] = [ai, bi] means M[x][y] should be incremented by one for all 0 <= x < ai and 0 <= y < bi.

Count and return the number of maximum integers in the matrix after performing all the operations.

Example 1:
Input: m = 3, n = 3, ops = [[2,2],[3,3]]
Output: 4
Explanation: The maximum integer in M is 2, and there are four of it in M. So return 4.

Example 2:
Input: m = 3, n = 3, ops = [[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3],[2,2],[3,3],[3,3],[3,3]]
Output: 4

Example 3:
Input: m = 3, n = 3, ops = []
Output: 9

Constraints:
1 <= m, n <= 4 * 104
0 <= ops.length <= 104
ops[i].length == 2
1 <= ai <= m
1 <= bi <= n
*/

package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minM := 1<<31 - 1
	minN := 1<<31 - 1
	// 0,0 ~ 가장 작은 M,N 이 계속 ops 수행시 계속 겹치게 되어 결국 가장 큰 갑을 가진 영역이 된다.
	for i := 0; i < len(ops); i++ {
		if ops[i][0] < minM {
			minM = ops[i][0]
		}
		if ops[i][1] < minN {
			minN = ops[i][1]
		}
	}
	return minM * minN
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	fmt.Println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}))
	fmt.Println(maxCount(3, 3, [][]int{}))
}
