/*
https://leetcode.com/problems/number-of-boomerangs/
447. Number of Boomerangs
Medium
You are given n points in the plane that are all distinct, where points[i] = [xi, yi]. A boomerang is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Return the number of boomerangs.

Example 1:
Input: points = [[0,0],[1,0],[2,0]]
Output: 2
Explanation: The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]].

Example 2:
Input: points = [[1,1],[2,2],[3,3]]
Output: 2

Example 3:
Input: points = [[1,1]]
Output: 0

Constraints:
n == points.length
1 <= n <= 500
points[i].length == 2
-104 <= xi, yi <= 104
All the points are unique.
*/

package main

import "fmt"

func distance(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

// time complexity: O(n*n)
func numberOfBoomerangs(points [][]int) int {
	r := 0
	for i := 0; i < len(points); i++ {
		// point[i] 기준으로 부메랑(i,j,k 일때 (i,j) == (i,k) 거리가 같은)을 만들 수 있는 체크
		m := make(map[int]int, 0)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			m[distance(points[i], points[j])]++
		}
		for _, v := range m {
			if v >= 2 {
				// permutation(순열) 개수로 누적
				// ex) point[i] 와 같은 거리의 point 가 3개 가 있는 경우
				// 순열은 3! 로 = 3*2*1 = 6 이지만
				// 0 번째 point[i] 를 고정하고, 나머지 1,2번재 point 2자리만 변경하면 된다.
				// 3*2 = 6로 계산한다.
				r += v * (v - 1)
			}
		}
	}
	return r
}

func main() {
	fmt.Println(numberOfBoomerangs([][]int{{0, 0}, {1, 0}, {2, 0}}))
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}, {2, 2}, {3, 3}}))
	fmt.Println(numberOfBoomerangs([][]int{{1, 1}}))
}
