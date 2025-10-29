/*
https://leetcode.com/problems/find-the-town-judge/
997. Find the Town Judge
Easy

In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:
The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.

Example 1:
Input: n = 2, trust = [[1,2]]
Output: 2

Example 2:
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

Example 3:
Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

Constraints:
1 <= n <= 1000
0 <= trust.length <= 104
trust[i].length == 2
All the pairs of trust are unique.
ai != bi
1 <= ai, bi <= n
*/
package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}
	cnt := make([]int, n+1)
	// 모든 사람은 town judge 를 믿음
	// 0번째 -> 1번째 사람 믿음,
	// 0번째 사람은 town judge 일 가능성 없음
	// 1번째 사람은 town judge 일 가능성 있음
	for _, t := range trust {
		cnt[t[0]]--
		cnt[t[1]]++
	}
	// 모든 사람이 믿는 사람(label)이 town judge
	for i := 0; i < len(cnt); i++ {
		if cnt[i] == n-1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
