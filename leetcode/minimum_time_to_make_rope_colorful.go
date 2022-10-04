/*
https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
1578. Minimum Time to Make Rope Colorful
Medium

Alice has n balloons arranged on a rope. You are given a 0-indexed string colors where colors[i] is the color of the ith balloon.

Alice wants the rope to be colorful. She does not want two consecutive balloons to be of the same color, so she asks Bob for help. Bob can remove some balloons from the rope to make it colorful. You are given a 0-indexed integer array neededTime where neededTime[i] is the time (in seconds) that Bob needs to remove the ith balloon from the rope.

Return the minimum time Bob needs to make the rope colorful.

Example 1:
Input: colors = "abaac", neededTime = [1,2,3,4,5]
Output: 3
Explanation: In the above image, 'a' is blue, 'b' is red, and 'c' is green.
Bob can remove the blue balloon at index 2. This takes 3 seconds.
There are no longer two consecutive balloons of the same color. Total time = 3.

Example 2:
Input: colors = "abc", neededTime = [1,2,3]
Output: 0
Explanation: The rope is already colorful. Bob does not need to remove any balloons from the rope.

Example 3:
Input: colors = "aabaa", neededTime = [1,2,3,4,1]
Output: 2
Explanation: Bob will remove the ballons at indices 0 and 4. Each ballon takes 1 second to remove.
There are no longer two consecutive balloons of the same color. Total time = 1 + 1 = 2.

Constraints:
n == colors.length == neededTime.length
1 <= n <= 105
1 <= neededTime[i] <= 104
colors contains only lowercase English letters.
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minCost(colors string, neededTime []int) int {
	r := 0
	for i := 1; i < len(colors); i++ {
		// 연속으로 같은 색깔의 풍선이면
		if colors[i] == colors[i-1] {
			// 적은 비용(neededTime)을 선택(풍선 제거)
			r += min(neededTime[i], neededTime[i-1])
			// aaaaa 같은 색깔이 연속으로 나오는 경우 가장 큰 비용값만 남겨야 한다.
			// 현재 연속된 비용중 큰값을 [i]에 업데이트해서 다음 비교시에 가장 큰 값만이 계속 유지되도록 한다.
			neededTime[i] = max(neededTime[i], neededTime[i-1])
		}
	}
	return r
}

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(minCost("abc", []int{1, 2, 3}))
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}))
}
