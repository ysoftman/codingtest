/*
https://leetcode.com/problems/all-paths-from-source-to-target/
797. All Paths From Source to Target
Medium
Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1 and return them in any order.

The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

Example 1:
Input: graph = [[1,2],[3],[3],[]]
Output: [[0,1,3],[0,2,3]]
Explanation: There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.

Example 2:
Input: graph = [[4,3,1],[3,2,4],[3],[4],[]]
Output: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]

Constraints:
n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i (i.e., there will be no self-loops).
All the elements of graph[i] are unique.
The input graph is guaranteed to be a DAG.
*/
package main

import "fmt"

// DFS(Depth First Search)
// dfs 로 탐색해 마지막 노드(마지막 인덱스)에 도착하면 결과로 추가
func dfs(graph [][]int, idx int, path []int, r *[][]int) {
	if idx == len(graph)-1 {
		temp := make([]int, len(path))
		copy(temp, path)
		*r = append(*r, temp)
		return
	}
	// 현재 노드(인덱스)에서 갈 수 있는 노드(인덱스)를 path 로 추가하고 dfs 재귀 수행
	for _, i := range graph[idx] {
		dfs(graph, i, append(path, i), r)
	}
}
func allPathsSourceTarget(graph [][]int) [][]int {
	r := [][]int{}
	path := []int{0}
	dfs(graph, 0, path, &r)
	return r
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}))
}
