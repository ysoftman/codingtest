/*
https://leetcode.com/problems/beautiful-arrangement/
526. Beautiful Arrangement
Medium
Suppose you have n integers labeled 1 through n. A permutation of those n integers perm (1-indexed) is considered a beautiful arrangement if for every i (1 <= i <= n), either of the following is true:

perm[i] is divisible by i.
i is divisible by perm[i].
Given an integer n, return the number of the beautiful arrangements that you can construct.

Example 1:
Input: n = 2
Output: 2
Explanation:
The first beautiful arrangement is [1,2]:
  - perm[1] = 1 is divisible by i = 1
  - perm[2] = 2 is divisible by i = 2

The second beautiful arrangement is [2,1]:
  - perm[1] = 2 is divisible by i = 1
  - i = 2 is divisible by perm[2] = 1

Example 2:
Input: n = 1
Output: 1

Constraints:
1 <= n <= 15
*/
package main

import "fmt"

func checkPermutation(visited []int, n, cur int, cnt *int) {
	if cur > n {
		// 여기까지 오면 beautiful arrangement 1개 완성
		(*cnt)++
	}
	// index 1 부터 시작함
	for i := 1; i <= n; i++ {
		// 확인했던 지점은 스킵
		if visited[i] == 1 {
			continue
		}
		// beautiful arrangement 의 원소가 될 수 있는지 체크
		if cur%i == 0 || i%cur == 0 {
			visited[i] = 1
			checkPermutation(visited, n, cur+1, cnt)
			visited[i] = 0
		}
	}
}
func countArrangement(n int) int {
	// index 1부터 시작하기 때문에 n+1
	visited := make([]int, n+1)
	cnt := 0
	// 숫자 1부터
	checkPermutation(visited, n, 1, &cnt)
	return cnt
}

func main() {
	fmt.Println(countArrangement(2))
	fmt.Println(countArrangement(1))
}
