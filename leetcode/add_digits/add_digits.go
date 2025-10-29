/*
https://leetcode.com/problems/add-digits/
258. Add Digits
Easy

Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

Example 1:
Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.

Example 2:
Input: num = 0
Output: 0

Constraints:
0 <= num <= 231 - 1

Follow up: Could you do it without any loop/recursion in O(1) runtime?
*/
package main

import "fmt"

func addDigits2(num int) int {
	if num < 10 {
		return num
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return addDigits2(sum)
}

// time complexity: O(1)
// mathematical approach
// 어차피 결과는 0~9 값인데
// 9로 나누어 떨어지는 경우0, 그외 9 modulo 값이 된다고 한다~
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(0))
}
