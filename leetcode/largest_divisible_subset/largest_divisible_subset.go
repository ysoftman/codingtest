/*
https://leetcode.com/problems/largest-divisible-subset/
368. Largest Divisible Subset
Medium

Given a set of distinct positive integers nums, return the largest subset answer such that every pair (answer[i], answer[j]) of elements in this subset satisfies:

answer[i] % answer[j] == 0, or
answer[j] % answer[i] == 0
If there are multiple solutions, return any of them.

Example 1:
Input: nums = [1,2,3]
Output: [1,2]
Explanation: [1,3] is also accepted.

Example 2:
Input: nums = [1,2,4,8]
Output: [1,2,4,8]

Constraints:
1 <= nums.length <= 1000
1 <= nums[i] <= 2 * 109
All the integers in nums are unique.
*/
package main

import (
	"fmt"
	"sort"
)

/*
ex) 2,4,5,8,16
dp[0]=[2] = 2 = [2]
dp[1]=[4] = 4 % 2 = 0 = [2] + [4] = [2,4]
dp[2]=[5] = 5 % 2 != 0 = [5]
dp[3]=[8] = maxlen(8 % 2 = 0 , 8 % 4 = 0) = [2,4] + [8] = [2,4,8]
dp[4]=[16] = maxlen(16 % 2 = 0 , 16 % 4 = 0 , 16 % 8 = 0) = [2,4,8] + [16] = [2,4,8,16]
*/
func largestDivisibleSubset(nums []int) []int {
	// 앞의 수가 작아서 작은 수로 뒤의 큰수를 나머지 연산을 하기 위해서 정렬한다.
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = []int{nums[i]}
	}
	maxIdx := 0
	maxlen := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// if nums[i]%nums[j] == 0 {
			// dp[j]+1 이 기존 dp[i] 원소개수보다 많아야 한다.
			// fmt.Println(nums[i], "%", nums[j], "=", nums[i]%nums[j], i, j, dp[i], dp[j])
			if nums[i]%nums[j] == 0 && len(dp[i]) < len(dp[j])+1 {
				// append 의 src 는 dp[j] 를 참조하기 때문에서 새로운 temp 를 만들어 복사후 temp 를 append 해야 한다.
				// temp := make([]int, len(dp[j]))
				// copy(temp, dp[j])
				// dp[i] = append(temp, []int{nums[i]}...)
				// 또는 append 시 새로운 빈공간에 dp[j] 원소를 복사
				dp[i] = append([]int{}, dp[j]...)
				dp[i] = append(dp[i], []int{nums[i]}...)
				if len(dp[i]) > maxlen {
					maxIdx = i
					maxlen = len(dp[i])
				}
			}
			// fmt.Println(dp)
		}
	}

	return dp[maxIdx]
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))
	fmt.Println(largestDivisibleSubset([]int{2, 4, 5, 8, 16}))
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8, 16}))
}
