/*
https://leetcode.com/problems/fraction-to-recurring-decimal/
166. Fraction to Recurring Decimal
Medium
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return any of them.

It is guaranteed that the length of the answer string is less than 104 for all the given inputs.

Example 1:
Input: numerator = 1, denominator = 2
Output: "0.5"

Example 2:
Input: numerator = 2, denominator = 1
Output: "2"

Example 3:
Input: numerator = 4, denominator = 333
Output: "0.(012)"
Constraints:
-231 <= numerator, denominator <= 231 - 1
denominator != 0
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(4, 333))
	fmt.Println(fractionToDecimal(1231, 342))
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// fraction : 분수
// numerator : 분자
// denominator : 분모
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	// 분자, 분모 부호가 다르면 - 인경우
	res := []string{}
	if (numerator > 0) && (denominator < 0) {
		res = append(res, "-")
	} else if (numerator < 0) && (denominator > 0) {
		res = append(res, "-")
	}

	// 부호 삭제
	num := abs(numerator)
	den := abs(denominator)

	res = append(res, fmt.Sprintf("%v", num/den))
	// fraction part 가 없는 경우
	if num%den == 0 {
		return strings.Join(res, "")
	}

	// fractional(소수점 아래 파악)
	res = append(res, ".")
	m := make(map[int]int)
	num %= den
	m[num] = len(res)
	for num != 0 {
		num *= 10
		res = append(res, fmt.Sprintf("%v", num/den))
		num %= den
		// 순환되는지 파악
		// insert (
		// append )
		if _, exist := m[num]; exist {
			index := m[num]
			res = append(res, "")
			copy(res[index+1:], res[index:])
			res[index] = "("
			res = append(res, ")")
			break
		} else {
			m[num] = len(res)
		}
	}
	// fmt.Println("----", res)
	return strings.Join(res, "")
}
