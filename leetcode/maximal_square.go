/*
https://leetcode.com/problems/maximal-square/
221. Maximal Square
Medium
Given an m x n binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.


Example 1:
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4

Example 2:
Input: matrix = [["0","1"],["1","0"]]
Output: 1

Example 3:
Input: matrix = [["0"]]
Output: 0

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] is '0' or '1'.
*/

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
정삭각형 확인방법(i,j) == 1 일때
1

위에서 i+0,j+0 까지는 확인했으니, i+1,j+1 들의 값만 확인하면 된다.
 [j+1][row] 에서의 row 1 1 체크
 ^
 |
11
11--> [col][i+1] 에서의 컬럼 1 1 체크

위에서 i+1,j+1 까지는 확인했으니, i+2,j+2 들의 값만 확인하면 된다.
  [j+2][row] 에서의 row 1 1 1체크
  ^
  |
111
111
111--> [col][i+2] 에서의 컬럼 1 1 1 체크
*/
// brute force
func maximalSquare2(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	maxLength := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			length := 1
			// 길이를 1씩 늘려가면서 square(정사각형)이 될 수 있는지 확인한다.
			for i+length < m && j+length < n {
				stop := false
				// j+length col 에서의 row 값들이 1 되는지 확인
				for k := i; k <= i+length; k++ {
					if matrix[k][j+length] == '0' {
						stop = true
						break
					}
				}
				if stop {
					break
				}
				// i+length row 에서의 col 값들이 1 되는지 확인
				for k := j; k <= j+length; k++ {
					if matrix[i+length][k] == '0' {
						stop = true
						break
					}
				}
				if stop {
					break
				}
				length++

			}
			maxLength = max(maxLength, length)
		}
	}
	return maxLength * maxLength
}

/*
dp 에 현재 위치에 될 가장 큰 정사각형이 될 수 있는 길이를 기록해 나가면서 가장 큰 길이를 찾는다.
이전 dp 값들에서 가장 작은 값을 선택+1
min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1

dp 값들이 다음과 같을때 현재 위치 x는
1, 1
0, x
min(1,1,0)+1 = 1
*/
// dynamic programming
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	maxLength := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				maxLength = max(maxLength, dp[i][j])
			}
		}
	}

	return maxLength * maxLength
}

func main() {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalSquare([][]byte{{'0', '1'}, {'1', '0'}}))
	fmt.Println(maximalSquare([][]byte{{'0'}}))
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '1', '1'}, {'0', '1', '0', '1', '0'}, {'1', '1', '0', '1', '1'}, {'1', '1', '0', '1', '1'}, {'0', '1', '1', '1', '1'}}))
}
