/*
https://leetcode.com/problems/min-cost-to-connect-all-points
1584. Min Cost to Connect All Points
Medium
You are given an array points representing integer coordinates of some points on a 2D-plane, where points[i] = [xi, yi].
The cost of connecting two points [xi, yi] and [xj, yj] is the manhattan distance between them: |xi - xj| + |yi - yj|, where |val| denotes the absolute value of val.
Return the minimum cost to make all points connected. All points are connected if there is exactly one simple path between any two points.

Example 1:
Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Explanation:
We can connect the points as shown above to get the minimum cost of 20.
Notice that there is a unique path between every pair of points.

Example 2:
Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18

Constraints:
1 <= points.length <= 1000
-106 <= xi, yi <= 106
All pairs (xi, yi) are distinct.
*/
package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// 현재 노드의 부모 노드 찾기
func findRootNode(nodeGroup []int, node int) int {
	// 현재 인데스와 값이 같을때가 최종 root 노드
	if node != nodeGroup[node] {
		// nodeGroup[node] 에 저장된 인덱스(노드)로 가서 다시 찾는다.
		nodeGroup[node] = findRootNode(nodeGroup, nodeGroup[node])
	}
	return nodeGroup[node]
}

func joinNode(nodeGroup, nodeRank []int, node1, node2 int) bool {
	n1 := findRootNode(nodeGroup, node1)
	n2 := findRootNode(nodeGroup, node2)
	// n1,n2 가 같은 부모노르를 가지고 있다면 이미 연결되어 있느 상태
	if n1 == n2 {
		return false
	}
	if nodeRank[n1] > nodeRank[n2] {
		nodeGroup[n2] = n1
	} else if nodeRank[n1] < nodeRank[n2] {
		nodeGroup[n1] = n2
	} else {
		nodeGroup[n1] = n2
		nodeRank[n2]++
	}
	return true
}

// Minimum Spanning Tree (MST) problem
func minCostConnectPoints(points [][]int) int {
	type edge struct {
		distance int
		node1    int
		node2    int
	}
	edges := make([]edge, 0)
	// 모든 엣지(노드1,노드2,노드1-2거리)를 구한다.
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges = append(edges, edge{
				distance: dist,
				node1:    i,
				node2:    j,
			})
		}
	}
	// 엣지를 거리순으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})
	// fmt.Println(edges)

	// 인덱스를 노드로 보고 값은 부모노드의 인덱스로 본다
	nodeGroup := make([]int, len(points))
	// 현재노드의 부모노드는 자신으로 초기화
	for i := 0; i < len(nodeGroup); i++ {
		nodeGroup[i] = i
	}
	nodeRank := make([]int, len(points))

	sumMinCost := 0
	for i := 0; i < len(edges); i++ {
		// node1와 node2를 연결가능한 경우 minCost 에 추가
		if joinNode(nodeGroup, nodeRank, edges[i].node1, edges[i].node2) {
			sumMinCost += edges[i].distance
		}
	}

	return sumMinCost
}
func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}
