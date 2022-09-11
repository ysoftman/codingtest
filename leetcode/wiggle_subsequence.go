/*
https://leetcode.com/problems/wiggle-subsequence/
376. Wiggle Subsequence
Medium
A wiggle sequence is a sequence where the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with one element and a sequence with two non-equal elements are trivially wiggle sequences.

For example, [1, 7, 4, 9, 2, 5] is a wiggle sequence because the differences (6, -3, 5, -7, 3) alternate between positive and negative.
In contrast, [1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences. The first is not because its first two differences are positive, and the second is not because its last difference is zero.
A subsequence is obtained by deleting some elements (possibly zero) from the original sequence, leaving the remaining elements in their original order.

Given an integer array nums, return the length of the longest wiggle subsequence of nums.

Example 1:
Input: nums = [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence with differences (6, -3, 5, -7, 3).

Example 2:
Input: nums = [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length.
One is [1, 17, 10, 13, 10, 16, 8] with differences (16, -7, 3, -3, 6, -8).

Example 3:
Input: nums = [1,2,3,4,5,6,7,8,9]
Output: 2

Constraints:
1 <= nums.length <= 1000
0 <= nums[i] <= 1000

Follow up: Could you solve this in O(n) time?
*/
package main

import "fmt"

// dp(dynamic programming)
// time complexity: O(n)
// space complexity: O(1)
func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	// 이전보다 큰 경우, 이전보다 작은 경우로 결과(최대 wiggle 수배열? 길이) 값을 저장 (dynamic programming)
	// up := make([]int, len(nums))
	// down := make([]int, len(nums))
	// 어차피 이전 결과 값을 이어 받게 되니, up, down 하나를 계속 업데이트하며 된다.
	up := 1
	down := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			// 이전수 보다 커지는 경우, down 을 다시 사용하되 up 에 +1
			up = down + 1
		} else if nums[i-1] > nums[i] {
			// 이전수 보다 작아지는 경우, up 을 다시 사용하되 down 에 +1
			down = up + 1
		} else {
			// 이전수와 같은 값이면 up, down 각각 이전 기록값을 이어나간다.
			up = up
			down = down
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
