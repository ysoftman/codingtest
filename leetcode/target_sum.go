/*
https://leetcode.com/problems/target-sum/
494. Target Sum
Medium
You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.

Example 1:
Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

Example 2:
Input: nums = [1], target = 1
Output: 1

Constraints:
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/

package main

import "fmt"

// brute-force : O(2^n)
func recursiveFindSum(nums []int, target, presum, idx int) int {
	if idx == len(nums)-1 {
		cnt := 0
		if presum+(-1*nums[idx]) == target {
			cnt++
		}
		if presum+(nums[idx]) == target {
			cnt++
		}
		return cnt
	}

	a := recursiveFindSum(nums, target, presum+(1*nums[idx]), idx+1)
	b := recursiveFindSum(nums, target, presum+(-1*nums[idx]), idx+1)
	return a + b
}

// func findTargetSumWays(nums []int, target int) int {
// 	return recursiveFindSum(nums, target, 0, 0)
// }

/*
p : 양수로 부분들의 합
n : 음수로 부분들의 합

다음 2개의 식을 만들 수 있다.
p - n = target
p + n = sum(nums)

위 2개을 더하면
2p = target + sum(nums)
p = ( target + sum(nums) ) / 2

p 는 target+sum(num)이 짝수 일때만 유효, p 를 만들 수 있는 sub sum 카운팅
*/
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// O(n*s)
func findTargetSumWays(nums []int, target int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}

	// 모두 다 더해도 target 보다 작으면 답이 없다.
	if s < abs(target) {
		return 0
	}
	// 홀수면 답이 없다.
	if (target+s)%2 == 1 {
		return 0
	}

	p := (target + s) / 2
	dp := make([]int, p+1)
	// 합이 0이 되는 경우는 아무것도 선택하지 않는 경우 1가지
	dp[0] = 1

	for _, n := range nums {
		for i := p; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}

	return dp[p]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, -1000))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
	fmt.Println(findTargetSumWays([]int{1, 0}, 1))
}
