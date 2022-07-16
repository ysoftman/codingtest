/*
https://leetcode.com/problems/out-of-boundary-paths/
576. Out of Boundary Paths
Medium
There is an m x n grid with a ball. The ball is initially at the position [startRow, startColumn]. You are allowed to move the ball to one of the four adjacent cells in the grid (possibly out of the grid crossing the grid boundary). You can apply at most maxMove moves to the ball.

Given the five integers m, n, maxMove, startRow, startColumn, return the number of paths to move the ball out of the grid boundary. Since the answer can be very large, return it modulo 109 + 7.

Example 1:
Input: m = 2, n = 2, maxMove = 2, startRow = 0, startColumn = 0
Output: 6

Example 2:
Input: m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
Output: 12

Constraints:
1 <= m, n <= 50
0 <= maxMove <= 50
0 <= startRow < m
0 <= startColumn < n
*/
package main

import "fmt"

// time limit exceeded!
// brute force, time complexity: O(4^n), space complexity: O(n) -> recursion depth
func findPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	// 값이 클 수 있어 modulo 모듈 값을 리턴하라고 문제에 써있음
	modulo := 1000000000 + 7
	if startRow < 0 || startRow == m ||
		startColumn < 0 || startColumn == n {
		return 1
	}
	if maxMove == 0 {
		return 0
	}
	// 4방향 각각에 대해서 boundary 를 벗어나는지 경우들을 모두 더한다.
	return findPaths(m, n, maxMove-1, startRow-1, startColumn)%modulo +
		findPaths(m, n, maxMove-1, startRow, startColumn-1)%modulo +
		findPaths(m, n, maxMove-1, startRow+1, startColumn)%modulo +
		findPaths(m, n, maxMove-1, startRow, startColumn+1)%modulo
}

// soulution 보면 이것도 accept 돼야 하는데, time limit exceeded! 된다. resucive call 때문인듯
// recursive + dynamic programming(memoization) 기록 해보자.
// time complexity: O(m*n*maxMove), space complexity: O(m*n*maxMove)
func findPathsWitRecursiveDP(m int, n int, maxMove int, startRow int, startColumn int, dp *[][][]int) int {
	// 값이 클 수 있어 modulo 모듈 값을 리턴하라고 문제에 써있음
	modulo := 1000000000 + 7
	if startRow < 0 || startRow == m ||
		startColumn < 0 || startColumn == n {
		return 1
	}
	if maxMove == 0 {
		return 0
	}
	if (*dp)[startRow][startColumn][maxMove] != -1 {
		return (*dp)[startRow][startColumn][maxMove]
	}
	// 4방향 각각에 대해서 boundary 를 벗어나는지 경우들을 모두 더한다.
	r := findPaths(m, n, maxMove-1, startRow-1, startColumn)%modulo +
		findPaths(m, n, maxMove-1, startRow, startColumn-1)%modulo +
		findPaths(m, n, maxMove-1, startRow+1, startColumn)%modulo +
		findPaths(m, n, maxMove-1, startRow, startColumn+1)%modulo
	(*dp)[startRow][startColumn][maxMove] = r
	return r
}
func findPaths1(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	// dp 기록된적이 없을 표시하기 위해 -1 로 초기화
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			// grid[i][j] 일때의 움직인 개수의 경우도 기록해야 한다.
			// 움직인 개수라 dp 공간은 +1더 많게 잡아 둔다.
			dp[i][j] = make([]int, maxMove+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	return findPathsWitRecursiveDP(m, n, maxMove, startRow, startColumn, &dp)
}

// dynamic programming(memoization) 기록 해보자.
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	// 값이 클 수 있어 modulo 모듈 값을 리턴하라고 문제에 써있음
	modulo := 1000000000 + 7
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1
	cnt := 0
	for moves := 1; moves <= maxMove; moves++ {

		// 현재 move 에서 grid 에 대한 DP
		tempGrid := make([][]int, m)
		for i := range tempGrid {
			tempGrid[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				// 한칸만 더 가면 boundary 를 벗어나는 경우 dp 값 사용
				if i == 0 {
					cnt = (cnt + dp[i][j]) % modulo
				}
				if i == m-1 {
					cnt = (cnt + dp[i][j]) % modulo
				}
				if j == 0 {
					cnt = (cnt + dp[i][j]) % modulo
				}
				if j == n-1 {
					cnt = (cnt + dp[i][j]) % modulo
				}

				// 그밖에 경우
				a := 0
				if i > 0 {
					a = dp[i-1][j]
				}
				b := 0
				if i < m-1 {
					b = dp[i+1][j]
				}
				c := 0
				if j > 0 {
					c = dp[i][j-1]
				}
				d := 0
				if j < n-1 {
					d = dp[i][j+1]
				}

				// 4방향에 대한 dp 합산
				// (문제 조건)더하기 할때마다 modulo 값을 넘어가지 않도록 해야 한다.
				tempGrid[i][j] = ((a+b)%modulo + (c+d)%modulo) % modulo
			}
		}

		dp = tempGrid
	}
	return cnt
}

func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
	fmt.Println(findPaths(1, 3, 3, 0, 1))
	fmt.Println(findPaths(8, 6, 10, 4, 3))
	fmt.Println(findPaths(8, 10, 10, 1, 8))
}
