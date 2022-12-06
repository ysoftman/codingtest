/*
https://leetcode.com/problems/remove-k-digits/
402. Remove K Digits
Medium

Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.

Example 1:
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

Example 2:
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

Example 3:
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.

Constraints:
1 <= k <= num.length <= 105
num consists of only digits.
num does not have any leading zeros except for the zero itself.
*/
package main

import "fmt"

func removeKdigits(num string, k int) string {
	stack := []rune{}
	for _, n := range num {
		// 앞에 수가 작아야 한다.
		// 마지막으로 추가된 수 가 n 보다 크면 빼기
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > n {
			// pop
			stack = stack[:len(stack)-1]
			k--
		}
		// 0 으로 시작하면 안됨
		if len(stack) == 0 && n == '0' {
			continue
		}
		stack = append(stack, n)
	}

	// 아직 k 번 제거가 필요한 경우 뒤부분 자르기
	// 이미 앞을 작은 수로 채웠기 때문에
	if k > 0 && len(stack) > 0 {
		if k > len(stack) {
			k = len(stack)
		}
		stack = stack[:len(stack)-k]
	}

	if len(stack) > 0 {
		return string(stack)
	}
	return "0"
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10200", 2))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("9", 1))
	fmt.Println(removeKdigits("10001", 4))
	fmt.Println(removeKdigits("1234567890", 9))
}
