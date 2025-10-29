/*
https://leetcode.com/problems/2-keys-keyboard
650. 2 Keys Keyboard
Medium
There is only one character 'A' on the screen of a notepad. You can perform one of two operations on this notepad for each step:

Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
Paste: You can paste the characters which are copied last time.
Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.

Example 1:

Input: n = 3
Output: 3
Explanation: Initially, we have one character 'A'.
In step 1, we use Copy All operation.
In step 2, we use Paste operation to get 'AA'.
In step 3, we use Paste operation to get 'AAA'.
Example 2:

Input: n = 1
Output: 0

Constraints:

1 <= n <= 1000
*/

/*
n = 1 -> 0, copy 후 paste 하면 AA 로 2개로 A 1개를 얻을 수 없다.
n = 2 -> 2, copy -> paste 2개의 operation
n = 3 -> 3, copy -> paste(AA) -> paste(AAA) 3개의 operation
n = 4 -> 4, copy -> paste(AA) -> paste(AAA) -> paste(AAAA)

	또는
	copy -> paste(AA) -> copy -> paste(AAAA)

n = 5 -> 5, copy -> paste(AA) -> paste(AAA) -> paste(AAAA) -> paste(AAAAA)
n = 6 -> 5, copy -> paste(AA) -> paste(AAA) -> copy -> paste(AAAAAA)
n = 7 -> 7, copy -> paste(AA) -> paste(AAA) -> paste(AAAA) -> paste(AAAAA) -> paste(AAAAAA) -> paste(AAAAAAA)
n = 8 -> 6, copy -> paste(AA) -> paste(AAA) -> paste(AAAA) -> copy -> paste(AAAAAAAA)

	또는
	copy -> paste(AA) -> copy -> paste(AAAA) -> paste(AAAAAA) -> paste(AAAAAAAA)

n = 9 -> 6, copy -> paste(AA) -> paste(AAA) -> copy -> paste(AAAAAA) -> paste(AAAAAAAAA)
n = 10 -> 7, copy -> paste(AA) -> paste(AAA) -> paste(AAAA) -> paste(AAAAA) -> copy -> paste(AAAAAAAAAA)
n = 11 -> 11

n 을 소인수 분해(참고로 n 이 소수인 경우는 n 스텝이 된다.)
n=6 -> 2x3 -> 2+3 = 5
n=9 -> 3x3 -> (copy -> paste(AA) -> paste(AAA))  (copy -> paste(AA) -> paste(AAA)) -> 3 + 3 = 6
*/
package main

import "fmt"

// 소인수 분해 O(root n)
func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	steps := 0
	factor := 2

	for n > 1 {
		// 소인수 분해 될때마다 step 카운트
		for n%factor == 0 {
			steps += factor
			n /= factor
		}
		factor++
	}
	return steps
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, minSteps(i))
	}
}
