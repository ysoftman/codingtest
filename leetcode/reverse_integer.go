/*
// https://leetcode.com/problems/reverse-integer/
7. Reverse Integer
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21
*/

package main

import "fmt"

func reverse(x int) int {
	minus := 1
	if x < 0 {
		minus = -1
		x *= minus
	}
	temp := []int{}
	for x > 9 {
		temp = append(temp, (x % 10))
		x /= 10
	}
	temp = append(temp, x)

	fmt.Println("temp:", temp)

	var result int
	start := false

	for i, v := range temp {
		a := 1
		if start == false && v == 0 {
			continue
		} else {
			start = true
			for j := 0; j < (len(temp)-i)-1; j++ {
				a *= 10
			}
		}
		result += v * a
	}
	result = result * minus
	fmt.Println("result:", result)

	if result > 1<<31-1 || result < -1<<31 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(123))
}
