/*
https://leetcode.com/problems/battleships-in-a-board/
419. Battleships in a Board
Medium

Given an m x n matrix board where each cell is a battleship 'X' or empty '.', return the number of the battleships on board.

Battleships can only be placed horizontally or vertically on board. In other words, they can only be made of the shape 1 x k (1 row, k columns) or k x 1 (k rows, 1 column), where k can be of any size. At least one horizontal or vertical cell separates between two battleships (i.e., there are no adjacent battleships).

Example 1:
Input: board = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]]
Output: 2

Example 2:
Input: board = [["."]]
Output: 0

Constraints:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] is either '.' or 'X'.

Follow up: Could you do it in one-pass, using only O(1) extra memory and without modifying the values board?
*/
package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
func countBattleships(board [][]byte) int {
	cnt := 0
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}

			// row == X && row-1 == x 라면 둘중 하나는 shape, 결국 1개의 battleship 으로 취급해야 한다.
			if i > 0 && board[i-1][j] == 'X' {
				continue
			}
			// col == X && col-1 == x 라면 둘중 하나는 shape, 결국 1개의 battleship 으로 취급해야 한다.
			if j > 0 && board[i][j-1] == 'X' {
				continue
			}
			cnt++
		}
	}
	return cnt
}

func main() {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	fmt.Println(countBattleships(board))
	board = [][]byte{{'.'}}
	fmt.Println(countBattleships(board))
}
