/*
https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
1155. Number of Dice Rolls With Target Sum
Medium

You have n dice and each die has k faces numbered from 1 to k.

Given three integers n, k, and target, return the number of possible ways (out of the kn total ways) to roll the dice so the sum of the face-up numbers equals target. Since the answer may be too large, return it modulo 109 + 7.

Example 1:
Input: n = 1, k = 6, target = 3
Output: 1
Explanation: You throw one die with 6 faces.
There is only one way to get a sum of 3.

Example 2:
Input: n = 2, k = 6, target = 7
Output: 6
Explanation: You throw two dice, each with 6 faces.
There are 6 ways to get a sum of 7: 1+6, 2+5, 3+4, 4+3, 5+2, 6+1.

Example 3:
Input: n = 30, k = 30, target = 500
Output: 222616187
Explanation: The answer must be returned modulo 109 + 7.

Constraints:
1 <= n, k <= 30
1 <= target <= 1000
*/
package main

import "fmt"

/*
ex) n = 2, k = 6, target = 7
dp[주사위개수][target] = target을 만들 수 있는 경우의 수
dp[1][1] = 1 = 1
dp[1][2] = 2 = 1
dp[1][3] = 3 = 1
dp[1][4] = 4 = 1
dp[1][5] = 5 = 1
dp[1][6] = 6 = 1
dp[1][7] = 0
dp[2][1] = 0
dp[2][2] = 1+1 = 1 => dp[2-1][2-1]...[2-1]
dp[2][3] = 1+2,2+1 = 2 => dp[2-1][3-2]...[3-1]
dp[2][4] = 1+3,3+1,2+2 = 3 => dp[2-1][4-3]...[4-1]
dp[2][5] = 1+4,4+1,2+3,3+2 = 4 => dp[2-1][5-4]...[5-1]
dp[2][6] = 1+5,5+1,3+3,2+4,4+2 = 5 => dp[2-1][6-5]...[6-1]
dp[2][7] = 1+5,5+1,3+3,3+3,2+4,4+2 = 6 => dp[2-1][7-6]...[7-1]
*/
func numRollsToTarget(n int, k int, target int) int {
	// 주사위를 굴렸을때 0~target 까지의 결과를 저장하기 위한 dp
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	m := 1_000_000_000 + 7
	// n 개의 주사위(0개의 주사위는 없음으로 1부터 시작)
	for i := 1; i <= n; i++ {
		// 1~target 이 될때까지
		for j := 1; j <= target; j++ {
			// 주사위의 1~k 면
			for f := 1; f <= k; f++ {
				// 주사위 면은 target 보다 작아야 한다.
				if f <= j {
					dp[i][j] += (dp[i-1][j-f] % m)
					dp[i][j] %= m
				}
			}
			// fmt.Println(i, j, dp[i][j])
		}
	}
	// n 개의 주사위 합이 target 이 되는 결과 리턴
	return dp[n][target]
}

func main() {
	fmt.Println(numRollsToTarget(1, 6, 3))
	fmt.Println(numRollsToTarget(2, 6, 7))
	fmt.Println(numRollsToTarget(30, 30, 500))
}
