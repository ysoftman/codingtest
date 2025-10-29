/*
https://leetcode.com/problems/base-7/
504. Base 7
Easy
Given an integer num, return a string of its base 7 representation.

Example 1:
Input: num = 100
Output: "202"

Example 2:
Input: num = -7
Output: "-10"

Constraints:
-107 <= num <= 107
*/
package main

import "fmt"

// iteration
func convertToBase7_(num int) string {
	r := ""
	quotient := num
	if quotient < 0 {
		quotient *= -1
	}
	for quotient >= 7 {
		r += fmt.Sprintf("%d", quotient%7)
		quotient /= 7
	}
	r += fmt.Sprintf("%d", quotient)
	// fmt.Println(r)
	r2 := ""
	for i := len(r) - 1; i >= 0; i-- {
		r2 += string(r[i])
	}
	if num < 0 {
		r2 = "-" + r2
	}
	return string(r2)
}

// simple recursive
func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(-num)
	}
	if num < 7 {
		return fmt.Sprintf("%d", num)
	}
	return convertToBase7(num/7) + fmt.Sprintf("%d", num%7)
}

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}
