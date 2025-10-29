/*
https://leetcode.com/problems/maximum-69-number/
1323. Maximum 69 Number
Easy

You are given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

Example 1:
Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.

Example 2:
Input: num = 9996
Output: 9999
Explanation: Changing the last digit 6 to 9 results in the maximum number.

Example 3:
Input: num = 9999
Output: 9999
Explanation: It is better not to apply any change.

Constraints:
1 <= num <= 104
num consists of only 6 and 9 digits.
*/
package main

import (
	"fmt"
	"strconv"
)

func maximum69Number2(num int) int {
	n := []int{}
	for num > 0 {
		n = append(n, num%10)
		num = num / 10
	}
	r := 0
	changed := false
	for i := len(n) - 1; i >= 0; i-- {
		r *= 10
		if !changed && n[i] == 6 {
			r += 9
			changed = true
			continue
		}
		r += n[i]
	}
	return r
}

// concise
func maximum69Number(num int) int {
	strNum := strconv.Itoa(num)
	r := make([]byte, len(strNum))
	copy(r, strNum)
	for i := 0; i < len(r); i++ {
		if r[i] == '6' {
			r[i] = '9'
			break
		}
	}
	// fmt.Println(string(r))
	result, _ := strconv.Atoi(string(r))
	return result
}

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}
