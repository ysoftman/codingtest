/*
https://leetcode.com/problems/min-cost-climbing-stairs/
746. Min Cost Climbing Stairs
Easy
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.

Example 1:
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.

Example 2:
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: You will start at index 0.
- Pay 1 and climb two steps to reach index 2.
- Pay 1 and climb two steps to reach index 4.
- Pay 1 and climb two steps to reach index 6.
- Pay 1 and climb one step to reach index 7.
- Pay 1 and climb two steps to reach index 9.
- Pay 1 and climb one step to reach the top.
The total cost is 6.

Constraints:
2 <= cost.length <= 1000
0 <= cost[i] <= 999
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
10,15,20 는 다음의 경우로 계단을 오를 수 있다.
1번째 계단 시작
10 + 15 + 20 = 45
10 + 20 = 30
2번째 계단에서 시작
15 = 15
15 + 20 = 35

계단끝에서 부터 생각해야 함.
계단끝에 올수 있는 경우 한계단전(-1) 또는 두계단전(-2)인데, 이중 작은 cost 를 선택
...
dp[i] = min(dp[i-2] + dp[i-1]) + cost[i]
*/
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{4, 5, 1, 4, 1, 441, 54, 15, 35, 16, 8, 76, 52, 3, 5, 1, 5, 4, 1, 34, 56, 1, 56, 6, 4, 1, 4, 6, 1, 46, 99}))
}
