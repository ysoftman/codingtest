/*
https://leetcode.com/problems/possible-bipartition/
886. Possible Bipartition
Medium

We want to split a group of n people (labeled from 1 to n) into two groups of any size. Each person may dislike some other people, and they should not go into the same group.

Given the integer n and the array dislikes where dislikes[i] = [ai, bi] indicates that the person labeled ai does not like the person labeled bi, return true if it is possible to split everyone into two groups in this way.

Example 1:
Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4] and group2 [2,3].

Example 2:
Input: n = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false

Example 3:
Input: n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false

Constraints:
1 <= n <= 2000
0 <= dislikes.length <= 104
dislikes[i].length == 2
1 <= dislikes[i][j] <= n
ai < bi
All the pairs of dislikes are unique.
*/
package main

import "fmt"

/*
n=4, 1~4 가 있음
dislikes = [[1,2],[1,3],[2,4]]
1 : 2, 3 를 싫어함, (같은 그룹이 될 수 없음)
2 : 1, 4  를 싫어함, (같은 그룹이 될 수 없음)
3 : 1  를 싫어함, (같은 그룹이 될 수 없음)
4 : 2  를 싫어함, (같은 그룹이 될 수 없음)

[1,4] 같은 그룹 가능
[2,3] 같은 그룹 가능
---> 여기까지 1,2,3,4 모두 그룹에 포함되어 true
[3,4] 같은 그룹 가능하지만 이미 3,4는 위 그룹에 포함
*/
func possibleBipartition(n int, dislikes [][]int) bool {
	// 숫자들간의 연결성을 나타내는 2차원 공간(그래프로 표현가능)을 만든다.
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	// 두 값이 dislike 임을 표시
	for _, v := range dislikes {
		graph[v[0]-1][v[1]-1] = 1
		graph[v[1]-1][v[0]-1] = 1
	}

	group := make([]int, n)
	for i := 0; i < n; i++ {
		// 현재 i 와 같은 그룹으로 묶일 숫자가 있는지 체크
		if group[i] == 0 && !findNum(graph, &group, i, 1) {
			return false
		}
	}
	return true
}
func findNum(graph [][]int, group *[]int, idx, groupType int) bool {
	(*group)[idx] = groupType
	for i := 0; i < len(graph[idx]); i++ {
		// idx, i 의 2개가 dislike 인경우
		if graph[idx][i] == 1 {
			// idx, i가 같은 그룹인 경우
			if (*group)[i] == (*group)[idx] {
				return false
			}
			// 아직 찾아보지 않은 경우 findNum 으로 찾기
			if (*group)[i] == 0 && !findNum(graph, group, i, -groupType) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	fmt.Println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))
}
