/*
https://leetcode.com/problems/combinations/
77. Combinations
Medium
Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

Example 2:
Input: n = 1, k = 1
Output: [[1]]
*/

/*
k = 4, n = 2
1,2 1,3 1,4
2,3 2,4
3,4

k = 4, n = 3
1,2,3 1,2,4 1,3,4
2,3,4
*/
package main

import "fmt"

func makeCombination(nums []int, k int, a int, temp []int, result *[][]int) {
	if a == len(nums) {
		return
	}
	temp = append(temp, nums[a])
	if len(temp) == k {
		// 슬라이스 temp 크기(길이)가 변경되지 않으면 메모리가 공유되기 때문에 새로운 값으로 복사해서 결과에 추가해야 한다.
		r := make([]int, len(temp))
		copy(r, temp)
		*result = append(*result, r)
		return
	}
	for i := a + 1; i < len(nums); i++ {
		makeCombination(nums, k, i, temp, result)
	}
}
func combine(n int, k int) [][]int {
	nums := []int{}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		temp := []int{}
		makeCombination(nums, k, i, temp, &result)
	}
	return result
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(4, 3))
	fmt.Println(combine(5, 3))
	fmt.Println(combine(5, 4))
}
