/*
https://leetcode.com/problems/counting-bits/
338. Counting Bits
Easy
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:
Input: n = 2
Output: [0,1,1]
Explanation:
0 --> 0
1 --> 1
2 --> 10

Example 2:
Input: n = 5
Output: [0,1,1,2,1,2]
Explanation:
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

Constraints:
0 <= n <= 105

Follow up:
It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?
*/

package main

import "fmt"

/*
i
0 --> 0     0
1 --> 1     1

2 --> 10    1     [i-2]+1
3 --> 11    2     [i-2]+1

4 --> 100   1     [i-2]
5 --> 101   2     [i-2]
6 --> 110   2     [i-2]+1
7 --> 111   3     [i-2]+1

8 --> 1000  1     [i-4]
9 --> 1001  2     [i-4]
10 -> 1010  2     [i-4]
11 -> 1011  3     [i-4]
12 -> 1100  2     [i-4]+1
13 -> 1101  3     [i-4]+1
14 -> 1110  3     [i-4]+1
15 -> 1111  4     [i-4]+1

16 -> 10000 1     [i-8]

현재 i가 2,4,8,16 처럼 10, 100, 100, 1000 될때 i/2로 처음 반은 i/2, 나중 반은 i/2 인덱스 +1 을 해 나간다.
*/
func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	if n == 0 {
		return dp
	}
	dp[1] = 1
	digit := 2
	digitCnt := 0
	for i := 2; i <= n; i++ {
		// check 2,4,8,16,...
		if i&(i-1) == 0 {
			digit = i
			digitCnt = 0
		}
		digitCnt++
		dp[i] = dp[i-(digit/2)]
		if digitCnt > digit/2 {
			dp[i]++
		}
	}
	return dp
}

func main() {
	fmt.Println(countBits(0))
	fmt.Println(countBits(1))
	fmt.Println(countBits(2))
	fmt.Println(countBits(3))
	fmt.Println(countBits(4))
	fmt.Println(countBits(5))
	fmt.Println(countBits(10))
	fmt.Println(countBits(20))
	fmt.Println(countBits(100))
}
