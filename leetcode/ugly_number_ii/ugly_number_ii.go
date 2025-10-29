/*
https://leetcode.com/problems/ugly-number-ii/
264. Ugly Number II
Medium
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.

Example 1:
Input: n = 10
Output: 12
Explanation: [1, 2, 3, 4, 5, 6, 8, 9, 10, 12] is the sequence of the first 10 ugly numbers.

Example 2:
Input: n = 1
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.

Constraints:
1 <= n <= 1690
*/
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// prime factors(소인수)가 2,3,5
// 1 -> 소수가 아니다. 소인수가 없어 2,3,5 모두 된다.
// 2 -> 2
// 3 -> 3
// 4 -> 2*2
// 5 -> 5
// 6 -> 2*3
// 7 -> 7 --> X
// 8 -> 2*2*2
// 9 -> 3*3
// 10 -> 2*5
// 11 -> 11 -> X
// 12 -> 2*2*3
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	p2 := 0
	p3 := 0
	p5 := 0
	i := 1
	for i < n {
		// i번째의 ugly num
		nums[i] = min(nums[p2]*2, min(nums[p3]*3, nums[p5]*5))

		// 2,3,5 곱이되는 이전 위치를 참조한다.(dynamic programming)
		// 다음 ugly num 을 위해 2,3,5 곱이되는 위치를 업데이트한다.
		if nums[i] == nums[p2]*2 {
			p2++
		}
		if nums[i] == nums[p3]*3 {
			p3++
		}
		if nums[i] == nums[p5]*5 {
			p5++
		}
		i++
	}
	// fmt.Println("nums:", nums)
	return nums[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1))
	fmt.Println(nthUglyNumber(120))
}
