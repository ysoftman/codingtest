/*
https://leetcode.com/problems/perfect-number/
507. Perfect Number
Easy
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself. A divisor of an integer x is an integer that can divide x evenly.

Given an integer n, return true if n is a perfect number, otherwise return false.

Example 1:
Input: num = 28
Output: true
Explanation: 28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, and 14 are all divisors of 28.

Example 2:
Input: num = 7
Output: false

Constraints:
1 <= num <= 108
*/
package main

import "fmt"

/*
perfect number 란
1 ... num-1 중에서 num%x==0 인 x 들의 합이 num 와 같아야 한다.
*/
// brute forace : time limit exceeded
func checkPerfectNumber2(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
		if num < sum {
			return false
		}
	}
	return sum == num
}

// time complexity : O(rootN)
func checkPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	// i*i 까지만 iteration
	// 28 = 1 + 2 + 4 + 7 + 14
	// 1+(28/1)+2+(28/2)+4+(28/4)=56
	// 마지막 리턴시 56-28(num) 을 빼준다.
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}
	}
	// fmt.Println("sum:", sum)
	return sum-num == num
}

func main() {
	fmt.Println(checkPerfectNumber(28))
	fmt.Println(checkPerfectNumber(7))
	fmt.Println(checkPerfectNumber(99_999_994))
}
