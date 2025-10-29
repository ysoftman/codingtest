/*
https://leetcode.com/problems/delete-and-earn/
740. Delete and Earn
Medium
Share
You are given an integer array nums. You want to maximize the number of points you get by performing the following operation any number of times:
Pick any nums[i] and delete it to earn nums[i] points. Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.
Return the maximum number of points you can earn by applying the above operation some number of times.

Example 1:
Input: nums = [3,4,2]
Output: 6
Explanation: You can perform the following operations:
- Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
- Delete 2 to earn 2 points. nums = [].
You earn a total of 6 points.

Example 2:
Input: nums = [2,2,3,3,3,4]
Output: 9
Explanation: You can perform the following operations:
- Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
- Delete a 3 again to earn 3 points. nums = [3].
- Delete a 3 once more to earn 3 points. nums = [].
You earn a total of 9 points.

Constraints:
1 <= nums.length <= 2 * 104
1 <= nums[i] <= 104
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deleteAndEarn(nums []int) int {
	m := make(map[int]int, 0)
	maxPoint := 0

	for i := 0; i < len(nums); i++ {
		/*
		   각 숫자를 지우면 얻을 수 있는 최대 포인트를 기록
		   [2,2,3,3,3,4] => 3을 선택한 경우, 3-2, 3+1 은 제거
		   [2,3,3] => 3포인트 += 3
		   [3] => 3포인트 += 3
		   [] => 3포인트 += 3
		*/
		m[nums[i]] += nums[i]
		// 가장 큰 포인트를 기록
		maxPoint = max(maxPoint, m[nums[i]])
	}
	// fmt.Println(m, maxPoint)
	// bottom-up dynamic-programming
	dp := make([]int, maxPoint+1)
	dp[0] = 0
	dp[1] = m[1]
	for i := 2; i <= maxPoint; i++ {
		/*
			[2,2,3,3,3,4] => map[2:2 3:3 4:4] 최대포인트:9
			dp[1], 과 dp[2]+(m[3]=4) 중 최대값 => 9
			maxPoint => 3포인트 9 인 경우
			dp[1]=0, 과 dp[0]=0+(m[2]=4) 중 최대값 => 4
			dp[2]=4, 과 dp[1]=0+(m[3]=4) 중 최대값 => 9
			dp[3]=9, 과 dp[2]=4+(m[4]=4) 중 최대값 => 9
			...
		*/
		// fmt.Println(i-1, dp[i-1], i-2, dp[i-2], m[i], dp[i-2]+m[i])
		dp[i] = max(dp[i-1], dp[i-2]+m[i])
	}
	return dp[maxPoint]
}

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4}))
	fmt.Println(deleteAndEarn([]int{1, 8, 5, 9, 6, 9, 4, 1, 7, 3, 3, 6, 3, 3, 8, 2, 6, 3, 2, 2, 1, 2, 9, 8, 7, 1, 1, 10, 6, 7, 3, 9, 6, 10, 5, 4, 10, 1, 6, 7, 4, 7, 4, 1, 9, 5, 1, 5, 7, 5}))
}
