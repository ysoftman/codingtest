/*
https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/
1186. Maximum Subarray Sum with One Deletion
Medium
Given an array of integers, return the maximum sum for a non-empty subarray (contiguous elements) with at most one element deletion. In other words, you want to choose a subarray and optionally delete one element from it so that there is still at least one element left and the sum of the remaining elements is maximum possible.

Note that the subarray needs to be non-empty after deleting one element.

Example 1:
Input: arr = [1,-2,0,3]
Output: 4
Explanation: Because we can choose [1, -2, 0, 3] and drop -2, thus the subarray [1, 0, 3] becomes the maximum value.

Example 2:
Input: arr = [1,-2,-2,3]
Output: 3
Explanation: We just choose [3] and it's the maximum sum.

Example 3:
Input: arr = [-1,-1,-1,-1]
Output: -1
Explanation: The final subarray needs to be non-empty. You can't choose [-1] and delete -1 from it, then get an empty subarray to make the sum equals to 0.


Constraints:
1 <= arr.length <= 105
-104 <= arr[i] <= 104
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time complexity: O(N)
// space complexity: O(2N) => O(N)
func maximumSum2(arr []int) int {
	dp_no_del := make([]int, len(arr))
	dp_with_del := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		dp_no_del[i] = arr[0]
		dp_with_del[i] = arr[0]
	}
	maxSum := arr[0]
	for i := 1; i < len(arr); i++ {
		dp_no_del[i] = max(dp_no_del[i-1]+arr[i], arr[i])
		dp_with_del[i] = max(dp_with_del[i-1]+arr[i], arr[i])
		if i >= 2 {
			dp_with_del[i] = max(dp_with_del[i], dp_no_del[i-2]+arr[i])
		}

		if maxSum < dp_with_del[i] {
			maxSum = dp_with_del[i]
		}
	}
	return maxSum
}

// time complexity: O(N)
// space complexity: O(2) => O(1)
func maximumSum(arr []int) int {
	max_no_del := arr[0]   // 원소 삭제 없이 sum
	max_with_del := arr[0] // 원소 삭제 후 sum
	max_sum := arr[0]
	// 0번째 원소는 위에서 이미 sum 에 포함되어 1부터 시작
	for i := 1; i < len(arr); i++ {
		// 현재 원소를 삭제(제외)를 고려한 최대 sum
		max_with_del = max(max(max_with_del+arr[i], arr[i]), max_no_del)
		// 현재 원소 제외 없이 최대 sum
		max_no_del = max(max_no_del+arr[i], arr[i])

		if max_sum < max_with_del {
			max_sum = max_with_del
		}
	}
	return max_sum
}

func main() {
	fmt.Println(maximumSum([]int{1, -2, 0, 3}))
	fmt.Println(maximumSum([]int{1, -2, -2, 3}))
	fmt.Println(maximumSum([]int{-1, -1, -1, -1}))
	fmt.Println(maximumSum([]int{-50}))
}
