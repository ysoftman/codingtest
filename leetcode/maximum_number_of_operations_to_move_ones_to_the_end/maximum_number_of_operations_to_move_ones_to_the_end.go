/*
https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end
3228. Maximum Number of Operations to Move Ones to the End
Medium
Hint
You are given a binary string s.

You can perform the following operation on the string any number of times:

Choose any index i from the string where i + 1 < s.length such that s[i] == '1' and s[i + 1] == '0'.
Move the character s[i] to the right until it reaches the end of the string or another '1'. For example, for s = "010010", if we choose i = 1, the resulting string will be s = "000110".
Return the maximum number of operations that you can perform.

Example 1:
Input: s = "1001101"
Output: 4
Explanation:
We can perform the following operations:
Choose index i = 0. The resulting string is s = "0011101".
Choose index i = 4. The resulting string is s = "0011011".
Choose index i = 3. The resulting string is s = "0010111".
Choose index i = 2. The resulting string is s = "0001111".

Example 2:
Input: s = "00111"
Output: 0

Constraints:
1 <= s.length <= 105
s[i] is either '0' or '1'.
*/
package main

import "fmt"

/*
maximum number of operation 계산은 다음과 같은 패턴이 된다.
앞에서 부터 0이 나타나기 까지 1 의 개수 카운트
1001110010 -> 1
0011110010 -> 4=(1+3)
0000111110 -> 5=(4+1)
1+4+5 = 10
*/
func maxOperations(s string) int {
	cnt := []int{0}
	oneCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCnt++
		} else if s[i] == '0' {
			if oneCnt > 0 {
				cnt = append(cnt, cnt[len(cnt)-1]+oneCnt)
				oneCnt = 0
			}
		}
	}

	sumCnt := 0
	for i := range cnt {
		sumCnt += cnt[i]
	}
	return sumCnt
}

func main() {
	fmt.Println(maxOperations("1001101"))
	fmt.Println(maxOperations("00111"))
	fmt.Println(maxOperations("1001110010"))
	fmt.Println(maxOperations("1001101101011101010110101010101011010101010110001010101"))
}
