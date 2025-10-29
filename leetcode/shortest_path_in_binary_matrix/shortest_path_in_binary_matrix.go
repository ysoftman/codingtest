/*
https://leetcode.com/problems/shortest-path-in-binary-matrix/
1091. Shortest Path in Binary Matrix
Medium
Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.
A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:
All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path.

Example 1:
Input: grid = [[0,1],[1,0]]
Output: 2

Example 2:
Input: grid = [[0,0,0],[1,1,0],[1,1,0]]
Output: 4

Example 3:
Input: grid = [[1,0,0],[1,1,0],[1,1,0]]
Output: -1

Constraints:
n == grid.length
n == grid[i].length
1 <= n <= 100
grid[i][j] is 0 or 1
*/

package main

import "fmt"

type node struct {
	x int
	y int
}

// DP 로 풀면 안된다. BFS로 풀어야 한다.
func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	if grid[0][0] != 0 {
		return -1
	}
	q := []node{}
	q = append(q, node{0, 0})

	visited := map[string]bool{}
	visited["0-0"] = true
	cnt := 0
	for len(q) > 0 {
		cnt++
		slibingQ := []node{}
		for len(q) > 0 {
			head := q[0]
			q = q[1:]

			// check 8-direction (sibling nodes in BFS)
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if head.y+i >= 0 && head.y+i < rows && head.x+j >= 0 && head.x+j < cols {
						// result(bottom-right)
						if head.y == rows-1 && head.x == cols-1 {
							return cnt
						}
						if grid[head.y+i][head.x+j] == 1 {
							continue
						}
						k := fmt.Sprintf("%d-%d", head.x+j, head.y+i)
						if visited[k] == false {
							slibingQ = append(slibingQ, node{head.x + j, head.y + i})
							visited[k] = true
						}
					}
				}
			}
		}
		q = append(q, slibingQ...)
	}
	return -1
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
}
