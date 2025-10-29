/*
https://leetcode.com/problems/lexicographical-numbers/
386. Lexicographical Numbers
Medium

Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.
You must write an algorithm that runs in O(n) time and uses O(1) extra space.

Example 1:
Input: n = 13
Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]

Example 2:
Input: n = 2
Output: [1,2]

Constraints:
1 <= n <= 5 * 104
*/

package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	r := make([]int, n)
	num := 1
	for i := 1; i <= n; i++ {
		r[i-1] = num
		// 2 -> 20
		// 20 -> 200
		if num*10 <= n {
			num *= 10
			continue
		}
		if num >= n {
			num /= 10
		}
		num++
		// 20 -> 2
		for num%10 == 0 {
			num /= 10
		}

	}
	return r
}

func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(lexicalOrder(i))
	}
	fmt.Println(lexicalOrder(230))
}
