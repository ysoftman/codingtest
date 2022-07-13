/*
https://leetcode.com/problems/matchsticks-to-square/
473. Matchsticks to Square
Medium
You are given an integer array matchsticks where matchsticks[i] is the length of the ith matchstick. You want to use all the matchsticks to make one square. You should not break any stick, but you can link them up, and each matchstick must be used exactly one time.

Return true if you can make this square and false otherwise.

Example 1:
Input: matchsticks = [1,1,2,2,2]
Output: true
Explanation: You can form a square with length 2, one side of the square came two sticks with length 1.

Example 2:
Input: matchsticks = [3,3,3,3,4]
Output: false
Explanation: You cannot find a way to form a square with all the matchsticks.

Constraints:
1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 108
*/

package main

import (
	"fmt"
	"sort"
)

/*
1,1,2,2,2
1,2      .. 변의 길이1로 했을때, 다음 2가 나와서 안됨
1+1,2,2,2  .. 1+1로 변의 길이를 했을때,

3,3,3,3,4
3,3,3,3 .. 3을 변이 길이로 했을때 4 가 남아서 안됨
3+3,3 .. 6을 변의 길이로 했을때 안됨
3+3,3+3,3 .. 6을 변의 길이로 했을때 안됨
3+3,3+3,3+4 .. 6을 변의 길이로 했을때 안됨

*/
// dfs, time complexity: O(4^n)
func recursive_matchsticks(correctLength, idx int, lengths []int, matchsticks []int) bool {
	// 마지막 성냥이라면 4변의 길이가 모두 같은지 본다.
	if idx == len(matchsticks) {
		for i := 1; i < len(lengths); i++ {
			if lengths[0] != lengths[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < 4; i++ {
		// time limit exceed 방지
		if lengths[i]+matchsticks[idx] > correctLength {
			continue
		}
		// 한변의 길이를 변경(증가)하면서 정사각형이 될 수 있는지 본다.
		lengths[i] += matchsticks[idx]
		// i 번째까지 설정된 변의 길이기준으로 재귀적 호출한다.
		if recursive_matchsticks(correctLength, idx+1, lengths, matchsticks) {
			return true
		}
		// 위 에서 true 가 이면 이전 값으로 복원해 다음 길이로 다시 진행할 수 있다.
		lengths[i] -= matchsticks[idx]
	}
	return false
}
func makesquare(matchsticks []int) bool {
	// 4변이 있어야 한다.
	if len(matchsticks) < 4 {
		return false
	}
	sum := 0
	for i := 0; i < len(matchsticks); i++ {
		sum += matchsticks[i]
	}
	// 4변의 길이가 같아야 한다.
	if sum%4 != 0 {
		return false
	}
	// descending sort 로 큰값으로 변의 길이를 찾으면 시도회수를 줄일 수 있다.
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	// fmt.Println(matchsticks)
	lengths := []int{0, 0, 0, 0}
	// 정사각형이 될 수 있는 한변의 길이
	correctLength := sum / 4
	return recursive_matchsticks(correctLength, 0, lengths, matchsticks)
}

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Println(makesquare([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}))
}
