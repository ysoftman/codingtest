/*
https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
1680. Concatenation of Consecutive Binary Numbers
Medium

Given an integer n, return the decimal value of the binary string formed by concatenating the binary representations of 1 to n in order, modulo 109 + 7.

Example 1:
Input: n = 1
Output: 1
Explanation: "1" in binary corresponds to the decimal value 1.

Example 2:
Input: n = 3
Output: 27
Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
After concatenating them, we have "11011", which corresponds to the decimal value 27.

Example 3:
Input: n = 12
Output: 505379714
Explanation: The concatenation results in "1101110010111011110001001101010111100".
The decimal value of that is 118505380540.
After modulo 109 + 7, the result is 505379714.

Constraints:
1 <= n <= 105
*/
package main

import (
	"fmt"
)

// brute-force
// n 이 크면 binstr 길이가 int64 범위 이상으로 너무 길어서 << 연산이 유효하지 않음
// func concatenatedBinary2(n int) int {
// 	binstr := ""
// 	for i := 1; i <= n; i++ {
// 		binstr += fmt.Sprintf("%v", strconv.FormatInt(int64(i), 2))
// 	}
// 	// fmt.Println(binstr)
// 	sum := 0
// 	for i := 0; i < len(binstr); i++ {
// 		if binstr[i] == '0' {
// 			continue
// 		}
// 		sum += 1 << (len(binstr) - 1 - i)
// 	}
// 	return sum % (1_000_000_000 + 7)
// }

/*
<< 를 이동 범위를 줄이기 위해서 다음과 같은 알고리즘을 사용해야 한다.
length=0, i=0;i<=n;i++ 일때
i값(2,10진수)
   0 0
   1 1 0&1==0 과 같이 0이면 length++ = 1,  0<<1 | 1 = 1 = 1
  10 2 1&2==0 과 같이 0이면 length++ = 2,  100(1<<2) | 10 = 110 = 6
 011 3 11000(110<<2) | 11 = 11011 = 27
...
*/
func concatenatedBinary(n int) int {
	sum := 0
	length := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			length++
		}
		sum = (sum << length) | i
		sum %= (1_000_000_000 + 7)
		// fmt.Println(sum)
	}
	return sum
}
func main() {
	fmt.Println(concatenatedBinary(1))
	fmt.Println(concatenatedBinary(3))
	fmt.Println(concatenatedBinary(12))
	fmt.Println(concatenatedBinary(42))
}
