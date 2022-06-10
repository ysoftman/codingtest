/*
https://leetcode.com/problems/count-submatrices-with-all-ones/
1504. Count Submatrices With All Ones
Medium
Given an m x n binary matrix mat, return the number of submatrices that have all ones.

Example 1:
Input: mat = [[1,0,1],[1,1,0],[1,1,0]]
Output: 13
Explanation:
There are 6 rectangles of side 1x1.
There are 2 rectangles of side 1x2.
There are 3 rectangles of side 2x1.
There is 1 rectangle of side 2x2.
There is 1 rectangle of side 3x1.
Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.

Example 2:
Input: mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
Output: 24
Explanation:
There are 8 rectangles of side 1x1.
There are 5 rectangles of side 1x2.
There are 2 rectangles of side 1x3.
There are 4 rectangles of side 2x1.
There are 2 rectangles of side 2x2.
There are 2 rectangles of side 3x1.
There is 1 rectangle of side 3x2.
Total number of rectangles = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24.

Constraints:
1 <= m, n <= 150
mat[i][j] is either 0 or 1.
*/

package main

import "fmt"

/*
1,0,1
1,1,0
1,1,0

일때, 0,0 시작점으로 할때 시작점을 포함하는 모두 1로 채워진 submatrix 찾기

(0,0 ~ 0,0)
(0,0 ~ 0,1)
(0,0 ~ 0,2)
(1,1 ~ 1,1) 인데 1,1 시작점일때 카운트되기 때문에 뺀다.
(1,1 ~ 1,2) 인데 1,1 시작점일때 카운트되기 때문에 뺀다.
(2,0 ~ 2,0) 인데 2,0 시작점일때 카운트되기 때문에 뺀다. => 그래서 6개

... 모든 시작점에 대해서 위와 같이 찾기

1 1
1 1
부분에서 0,0 시작점으로 하면
(0,0 ~ 0,0)
(0,0 ~ 0,1)
(0,0 ~ 0,1)
(0,0 ~ 1,1) 인데, 1,1시작점일때 카운트되기 때문에 뺀다. => 그래서 3개

*/

func countSubMat(mat [][]int, starty, startx int) int {
	cnt := 0
	endy := len(mat)
	endx := len(mat[0])
	for i := starty; i < endy; i++ {
		for j := startx; j < endx; j++ {
			if mat[i][j] == 1 {
				cnt++
			} else {
				// 0을 만나면 더이상 찾지 않는다.(다음 시작점에서 계속 찾을 수 있다.)
				endx = j
			}
		}
	}
	return cnt
}

// time complexity : O(m*n * a*b)
func numSubmat(mat [][]int) int {
	cnt := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				cnt += countSubMat(mat, i, j)
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numSubmat([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(numSubmat([][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}))
}
