/*
https://leetcode.com/problems/partition-equal-subset-sum/
416. Partition Equal Subset Sum
Medium

Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/
package main

import "fmt"

/*
dynamic programming

[1,5,11,5] =>  sum => 22
22/2 = 11 을 만들 수 있는지 찾기

dp[11] = dp[11] or 11 에서 nums[i] 뺀 dp (dp[11-1], dp[11-5], dp[11-11], dp[11-5])

dp[0] = true 로 시작

nums[0] = 1 일때
dp[11] = dp[11](false) or dp[11-1](false)
dp[10] = dp[10](false) or dp[10-1](false)
dp[9] = dp[9](false) or dp[9-1](false)
dp[8] = dp[8](false) or dp[8-1](false)
dp[7] = dp[7](false) or dp[7-1](false)
dp[6] = dp[6](false) or dp[6-1](false)
dp[5] = dp[5](false) or dp[5-1](false)
dp[4] = dp[4](false) or dp[4-1](false)
dp[3] = dp[3](false) or dp[3-1](false)
dp[2] = dp[2](false) or dp[2-1](false)
dp[1] = dp[1](true) or dp[1-1](true)

nums[1] = 5 일때
dp[11] = dp[11](false) or dp[11-5](false)
dp[10] = dp[10](false) or dp[10-5](false)
dp[9] = dp[9](false) or dp[9-5](false)
dp[8] = dp[8](false) or dp[8-5](false)
dp[7] = dp[7](false) or dp[7-5](false)
dp[6] = dp[6](true) or dp[6-5](true)
dp[5] = dp[5](true) or dp[5-5](true)

nums[2] = 11 일때
dp[11] = dp[11](true) or dp[11-11](true) --> 11인 경우를 찾았음

nums[3] = 5 일때
dp[11] = dp[11](true) or dp[11-5](true)
dp[10] = dp[10](true) or dp[10-5](true)
dp[9] = dp[9](false) or dp[9-5](false)
dp[8] = dp[8](false) or dp[8-5](false)
dp[7] = dp[7](false) or dp[7-5](false)
dp[6] = dp[6](true) or dp[6-5](true)
dp[5] = dp[5](true) or dp[5-5](true)
*/
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 홀수면 2개의 subset 합이 같을 수 없다.
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j > 0; j-- {
			if j < nums[i] {
				continue
			}
			dp[j] = dp[j] || dp[j-nums[i]]
			// fmt.Printf("dp[%v] = dp[%v](%v) or dp[%v-%v](%v)\n", j, j, dp[j], j, nums[i], dp[j-nums[i]])
		}
	}
	return dp[sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
