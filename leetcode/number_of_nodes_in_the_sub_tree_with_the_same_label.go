/*
https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
1519. Number of Nodes in the Sub-Tree With the Same Label
Medium
You are given a tree (i.e. a connected, undirected graph that has no cycles) consisting of n nodes numbered from 0 to n - 1 and exactly n - 1 edges. The root of the tree is the node 0, and each node of the tree has a label which is a lower-case character given in the string labels (i.e. The node with the number i has the label labels[i]).

The edges array is given on the form edges[i] = [ai, bi], which means there is an edge between nodes ai and bi in the tree.

Return an array of size n where ans[i] is the number of nodes in the subtree of the ith node which have the same label as node i.

A subtree of a tree T is the tree consisting of a node in T and all of its descendant nodes.

Example 1:
Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], labels = "abaedcd"
Output: [2,1,1,1,1,1,1]
Explanation: Node 0 has label 'a' and its sub-tree has node 2 with label 'a' as well, thus the answer is 2. Notice that any node is part of its sub-tree.
Node 1 has a label 'b'. The sub-tree of node 1 contains nodes 1,4 and 5, as nodes 4 and 5 have different labels than node 1, the answer is just 1 (the node itself).

Example 2:
Input: n = 4, edges = [[0,1],[1,2],[0,3]], labels = "bbbb"
Output: [4,2,1,1]
Explanation: The sub-tree of node 2 contains only node 2, so the answer is 1.
The sub-tree of node 3 contains only node 3, so the answer is 1.
The sub-tree of node 1 contains nodes 1 and 2, both have label 'b', thus the answer is 2.
The sub-tree of node 0 contains nodes 0, 1, 2 and 3, all with label 'b', thus the answer is 4.

Example 3:
Input: n = 5, edges = [[0,1],[0,2],[1,3],[0,4]], labels = "aabab"
Output: [3,2,1,1,1]

Constraints:
1 <= n <= 105
edges.length == n - 1
edges[i].length == 2
0 <= ai, bi < n
ai != bi
labels.length == n
labels is consisting of only of lowercase English letters.
*/
package main

import "fmt"

// 현재 노드 하위에서 labels 의 개수 파악
func countLabels(node, parent int, graph [][]int, labels string, result *[]int) map[byte]int {
	cntLabels := make(map[byte]int, 26)
	cntLabels[labels[node]] = 1
	for i := 0; i < len(graph[node]); i++ {
		// 그래프라 부모 노드를 다음과 같이 파악해서 스킵한다.
		if graph[node][i] == parent {
			continue
		}
		// 현재 노드 레이블 개수 업데이트(하위 노드 레이블 파악해서 추가)
		childLabels := countLabels(graph[node][i], node, graph, labels, result)
		for j := 'a'; j <= 'z'; j++ {
			cntLabels[byte(j)] += childLabels[byte(j)]
		}
	}
	(*result)[node] += cntLabels[labels[node]]
	return cntLabels
}

// n : node 개수
// edges : node 연결 정보
// labels : node(번호순)의 레이블 정보
func countSubTrees(n int, edges [][]int, labels string) []int {
	// graph 데이터 구조 생성
	graph := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	// fmt.Println(graph)
	result := make([]int, n)
	countLabels(0, -1, graph, labels, &result)
	return result
}

func main() {
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
	fmt.Println(countSubTrees(4, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb"))
	fmt.Println(countSubTrees(5, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab"))
}
