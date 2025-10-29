/*
https://leetcode.com/problems/word-search/
79. Word Search
Medium
Given an m x n grid of characters board and a string word, return true if word exists in the grid.
The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example 1:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Example 2:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Example 3:
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Constraints:
m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board and word consists of only lowercase and uppercase English letters.
*/
package main

import "fmt"

func findWord(board [][]byte, i, j int, word string) bool {
	// fmt.Println(word)
	if len(word) == 0 {
		return true
	}
	m := len(board)
	n := len(board[0])
	if !(i >= 0 && i < m && j >= 0 && j < n) {
		return false
	}
	if board[i][j] == '_' {
		return false
	}

	if word[0] == board[i][j] {
		temp := board[i][j]
		board[i][j] = '_'
		if findWord(board, i-1, j, word[1:]) {
			return true
		}
		if findWord(board, i+1, j, word[1:]) {
			return true
		}
		if findWord(board, i, j-1, word[1:]) {
			return true
		}
		if findWord(board, i, j+1, word[1:]) {
			return true
		}
		board[i][j] = temp
	}
	return false
}
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findWord(board, i, j, word) {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
