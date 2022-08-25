/*
https://leetcode.com/problems/count-numbers-with-unique-digits/
357. Count Numbers with Unique Digits
Medium
Given an integer n, return the count of all numbers with unique digits, x, where 0 <= x < 10n.

Example 1:
Input: n = 2
Output: 91
Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100, excluding 11,22,33,44,55,66,77,88,99

Example 2:
Input: n = 0
Output: 1

Constraints:
0 <= n <= 8
*/
package main

import "fmt"

/*
패턴 파악
0=>1
1=>10  10
2=>91  10+(9*9)
3=>739 10+(9*9)+(9*9*8)
4=>5275 10+(9*9)+(9*9*8)+(9*9*8*7)
5=>32491 10+(9*9)+(9*9*8)+(9*9*8*7)+(9*9*8*7*6)
6=>168571 10+(9*9)+(9*9*8)+(9*9*8*7)+(9*9*8*7*6)+(9*9*8*7*6*5)
7=>712891 10+(9*9)+(9*9*8)+(9*9*8*7)+(9*9*8*7*6)+(9*9*8*7*6*5)+(9*9*8*7*6*5*4)
*/
func countNumbersWithUniqueDigits2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	temp := 9 * 9
	cnt := 10 + temp
	for i := 3; i <= n; i++ {
		temp = temp * (11 - i)
		cnt += temp
	}
	return cnt
}

func recursiveCntUniqDigits(n int, curdigit int, digits []bool) int {
	if n == curdigit {
		return 1
	}
	cnt := 1
	i := 0
	if curdigit == 0 {
		i = 1
	}
	for i <= 9 {
		// 중복으로 사용된 0~9 숫자가 없다면 카운트
		if !digits[i] {
			digits[i] = true
			cnt += recursiveCntUniqDigits(n, curdigit+1, digits)
			digits[i] = false
		}
		i++
	}
	return cnt
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	// 0~9 10개의 숫자가 쓰였는지 여부 파악을 위해
	return recursiveCntUniqDigits(n, 0, make([]bool, 10))
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(0))
	fmt.Println(countNumbersWithUniqueDigits(1))
	fmt.Println(countNumbersWithUniqueDigits(2))
	fmt.Println(countNumbersWithUniqueDigits(3))
	fmt.Println(countNumbersWithUniqueDigits(4))
	fmt.Println(countNumbersWithUniqueDigits(5))
	fmt.Println(countNumbersWithUniqueDigits(7))
	fmt.Println(countNumbersWithUniqueDigits(8))
}
