/*
https://leetcode.com/problems/missing-number/
268. Missing Number
Easy
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

Example 2:
Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.

Example 3:
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

Constraints:
n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
All the numbers of nums are unique.

Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?
*/
package main

import (
	"fmt"
	"sort"
)

// time complexity: O(NlogN)
func missingNumber2(nums []int) int {
	sort.Ints(nums)
	i := 0
	for i < len(nums) {
		if nums[i] != i {
			return i
		}
		i++
	}
	return i
}

// time complexity: O(N)
func missingNumber(nums []int) int {
	actualSum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		actualSum += nums[i]
		sum += i
	}
	sum += len(nums)
	// fmt.Println("actualSum:", actualSum)
	// fmt.Println("sum:", sum)
	return sum - actualSum
}

func main() {
	fmt.Println(missingNumber2([]int{3, 0, 1}))
	fmt.Println(missingNumber2([]int{0, 1}))
	fmt.Println(missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
