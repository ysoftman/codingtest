/*
https://leetcode.com/problems/factorial-trailing-zeroes/
172. Factorial Trailing Zeroes
Medium

Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

Example 1:
Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:
Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.

Example 3:
Input: n = 0
Output: 0

Constraints:
0 <= n <= 104

Follow up: Could you write a solution that works in logarithmic time complexity?
*/
package main

import "fmt"

// func factorial(n int) int {
//     if n <= 1 {
//         return n
//     }
//     return n * factorial(n-1)
// }
// // n 이 크면 factorial 결과 오버플로우 error 발생해서 안됨
// func trailingZeroes2(n int) int {
//     if n == 0 {
//         return 0
//     }
//     num := factorial(n)
//     strNum := fmt.Sprintf("%d", num)
//     fmt.Println("strNum:", strNum)
//     cnt := 0
//     for i:=len(strNum)-1;i>=0;i-- {
//         if strNum[i] == '0' {
//             cnt++
//         } else {
//             break
//         }
//     }
//     return cnt
// }

/*
0이 되는 경우 5의배수가 있어야 한다.
2*5 = 10
4*5 = 20
6*5 = 30
8*5 = 40
2*10 = 20
2*15 = 30 ...

ex) 5
5 / 5 = 1
5 / 25 = 0.xxx 1보다 작으면 끝
1

ex) 30
30 / 5 = 6
30 / 25 = 1
6+1 = 7

ex) 100
100 / 5 = 20
100 / 25 = 4
20+4 = 24

ex) 250
250 / 5 = 50
250 / 25 = 10
250 / 125 = 2
50+10+2 = 62
*/
func trailingZeroes(n int) int {
	cnt := 0
	i := 5
	for n/i > 0 {
		cnt += n / i
		// fmt.Printf("%v / %v = %v\n", n, i, n/i)
		i *= 5
	}
	return cnt
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
	fmt.Println(trailingZeroes(30))
	fmt.Println(trailingZeroes(100))
	fmt.Println(trailingZeroes(250))
}
