/*
https://leetcode.com/problems/binary-prefix-divisible-by-5/description
1018. Binary Prefix Divisible By 5
Easy
You are given a binary array nums (0-indexed).

We define xi as the number whose binary representation is the subarray nums[0..i] (from most-significant-bit to least-significant-bit).

For example, if nums = [1,0,1], then x0 = 1, x1 = 2, and x2 = 5.
Return an array of booleans answer where answer[i] is true if xi is divisible by 5.

Example 1:
Input: nums = [0,1,1]
Output: [true,false,false]
Explanation: The input numbers in binary are 0, 01, 011; which are 0, 1, and 3 in base-10.
Only the first number is divisible by 5, so answer[0] is true.

Example 2:
Input: nums = [1,1,1]
Output: [false,false,false]

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
*/
package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	r := []bool{}
	temp := 0
	for i := range nums {
		// 큰 수에서 overflow 발생
		// temp <<= 1
		// 기존 수를 %5로 유지하고 새 자리수를 더할때는 2배후 더한면 overflow 없이 5로 나눠지는 판별 가능하다.
		temp %= 5
		temp *= 2
		temp += nums[i]
		if temp%5 == 0 {
			r = append(r, true)
		} else {
			r = append(r, false)
		}
	}
	// fmt.Println(temp)
	return r
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}))
}
