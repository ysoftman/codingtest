/*
https://leetcode.com/problems/game-of-life/
289. Game of Life
Medium

According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population.
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously. Given the current state of the m x n grid board, return the next state.

Example 1:
Input: board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
Output: [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

Example 2:
Input: board = [[1,1],[1,0]]
Output: [[1,1],[1,1]]

Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 25
board[i][j] is 0 or 1.

Follow up:
Could you solve it in-place? Remember that the board needs to be updated simultaneously: You cannot update some cells first and then use their updated values to update other cells.
In this question, we represent the board using a 2D array. In principle, the board is infinite, which would cause problems when the active area encroaches upon the border of the array (i.e., live cells reach the border). How would you address these problems?
*/
package main

import "github.com/ysoftman/ysoftmancommon"

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

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			// 3x3 범위(윈도우)안에서 1 카운트
			for k := max(i-1, 0); k < min(i+2, m); k++ {
				for l := max(j-1, 0); l < min(j+2, n); l++ {
					//
					if board[k][l] == 1 || board[k][l] == 3 {
						cnt++
					}
				}
			}
			// 현재 윈도우내에서 1의 개수가 3 이거나
			// 현재 위치의 빼고 나머지가 3이면
			if cnt == 3 || cnt-board[i][j] == 3 {
				// 현재 셀의 값을 2 또는 3 값으로 만든다.
				// 2,3 이며 다음 루프시 위에 1,3 만 카운트한다.
				board[i][j] |= 2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 0,2(10),3(11) 인 상태에서 오른쪽 쉬프트 하면
			// 0,1(01),1(01)
			board[i][j] >>= 1
		}
	}
}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	ysoftmancommon.PrintMatrix(board)
}
