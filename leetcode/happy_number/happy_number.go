/*
https://leetcode.com/problems/happy-number/
202. Happy Number
Easy
Write an algorithm to determine if a number n is happy.
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:
Input: n = 19
Output: true
Explanation:
1**2 + 9**2 = 82
8**2 + 2**2 = 68
6**2 + 8**2 = 100
1**2 + 0**2 + 0**2 = 1

Example 2:
Input: n = 2
Output: false
*/
package main

import "fmt"

func isHappy(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}

	for n != 1 {
		sum := 0
		for n > 0 {
			remainder := n % 10
			sum += remainder * remainder
			n = n / 10
		}
		// check infinite loop
		if sum == 4 {
			return false
		}
		n = sum
		// fmt.Println(n)
	}
	return true
}
func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
