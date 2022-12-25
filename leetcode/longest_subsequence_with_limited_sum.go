/*
https://leetcode.com/problems/longest-subsequence-with-limited-sum/
2389. Longest Subsequence With Limited Sum
Easy

You are given an integer array nums of length n, and an integer array queries of length m.

Return an array answer of length m where answer[i] is the maximum size of a subsequence that you can take from nums such that the sum of its elements is less than or equal to queries[i].

A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

Example 1:
Input: nums = [4,5,2,1], queries = [3,10,21]
Output: [2,3,4]
Explanation: We answer the queries as follows:
- The subsequence [2,1] has a sum less than or equal to 3. It can be proven that 2 is the maximum size of such a subsequence, so answer[0] = 2.
- The subsequence [4,5,1] has a sum less than or equal to 10. It can be proven that 3 is the maximum size of such a subsequence, so answer[1] = 3.
- The subsequence [4,5,2,1] has a sum less than or equal to 21. It can be proven that 4 is the maximum size of such a subsequence, so answer[2] = 4.

Example 2:
Input: nums = [2,3,4,5], queries = [1]
Output: [0]
Explanation: The empty subsequence is the only subsequence that has a sum less than or equal to 1, so answer[0] = 0.

Constraints:
n == nums.length
m == queries.length
1 <= n, m <= 1000
1 <= nums[i], queries[i] <= 106
*/
package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)

	accumulated := make([]int, len(nums))
	accumulated[0] = nums[0]
	// nums i번째 요소까지 누적합을 기록해둔다.
	for i := 1; i < len(nums); i++ {
		accumulated[i] += accumulated[i-1] + nums[i]
	}
	r := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(accumulated); j++ {
			// 쿼리보다 작은 누적합의 위치(누적된 elemenet 개수) 기록
			if accumulated[j] > queries[i] {
				r[i] = j
				break
			}
		}
		// 모든 원소를 다 합해도 쿼리 숫자 보다 작은 경우
		if accumulated[len(accumulated)-1] <= queries[i] {
			r[i] = len(accumulated)
		}
	}
	return r
}

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
	fmt.Println(answerQueries([]int{1_000_000}, []int{1_000_000}))
}
