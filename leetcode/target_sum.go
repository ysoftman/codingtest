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

func findTargetSumWays(nums []int, target int) int {
	return recursiveFindSum(nums, target, 0, 0)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
	fmt.Println(findTargetSumWays([]int{1, 0}, 1))
}
