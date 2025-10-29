/*
https://leetcode.com/problems/surrounded-regions/
130. Surrounded Regions
Medium
Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example 1:
Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
Explanation: Notice that an 'O' should not be flipped if:
- It is on the border, or
- It is adjacent to an 'O' that should not be flipped.
The bottom 'O' is on the border, so it is not flipped.
The other three 'O' form a surrounded region, so they are flipped.

Example 2:
Input: board = [["X"]]
Output: [["X"]]

Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is 'X' or 'O'.
*/
package main

import "fmt"

func printBoard(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Println()
	}
}

/*
XXXX
XOOX
XXOX
XOXX

4방향 검사해서 X 막혀있지 않은곳 a 로 표시
XXXX
XOOX
XXOX
XaXX

a -> O 변경하고 나머지는 모두 X
XXXX
XXXX
XXXX
XOXX
*/
func checkOpen(board [][]byte, m, n, i, j int) {
	// 뚫린 경우
	if board[i][j] == 'O' {
		board[i][j] = 'a'
		// 위쪽 뚫였는지 확인
		if i > 1 {
			checkOpen(board, m, n, i-1, j)
		}
		// 오른쪽 뚫였는지 확인
		if j > 1 {
			checkOpen(board, m, n, i, j-1)
		}
		// 아래쪽 뚫였는지 확인
		if i+1 < m {
			checkOpen(board, m, n, i+1, j)
		}
		// 왼쪽 뚫였는지 확인
		if j+1 < n {
			checkOpen(board, m, n, i, j+1)
		}
	}
}
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		// 왼쪽
		checkOpen(board, m, n, i, 0)
		if n > 1 {
			// 오른쪽
			checkOpen(board, m, n, i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		// 위쪽
		checkOpen(board, m, n, 0, j)
		if m > 1 {
			// 아래쪽
			checkOpen(board, m, n, m-1, j)
		}
	}
	// a -> O  나머진 X 표시
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'a' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	printBoard(board)
	board = [][]byte{{'X'}}
	solve(board)
	printBoard(board)
	board = [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	solve(board)
	printBoard(board)
}
