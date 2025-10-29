/*
https://leetcode.com/problems/bitwise-and-of-numbers-range/
201. Bitwise AND of Numbers Range
Medium
Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.

Example 1:
Input: left = 5, right = 7
Output: 4

Example 2:
Input: left = 0, right = 0
Output: 0

Example 3:
Input: left = 1, right = 2147483647
Output: 0

Constraints:
0 <= left <= right <= 231 - 1
*/
package main

import "fmt"

// time limit exceeded
func rangeBitwiseAnd2(left int, right int) int {
	r := left
	for i := left + 1; i <= right; i++ {
		r &= i
	}
	return r
}

/*
5 101
6 110
7 111

5 7 사의 값은 0 이 있 and 로 어차피 0이 된다.
5,7 만 오른쪽으로 1 shift => 5, 7 둘다 값이 0이 될때까지
5(101) > 2(10) > 1(1)
7(111) > 3(11) > 1(1)

3번하면 0으로 같아진다. 0 왼쪽으로 값이 있엇도 값이 같으면 된다.

원래 5값을 다시 원위치로 3번 오른쪽 쉬프트하면 MSB(most significant bit) 로분 3자리는 0으로 채워진ㄷ.
x > 3번 오른쪽으로 shift > x000
*/
func rangeBitwiseAnd(left int, right int) int {
	cnt := 0
	for left != right {
		left >>= 1
		right >>= 1
		cnt++
	}
	// fmt.Printf("left:%v, right:%v\n", left, right)
	return left << cnt
}

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
	fmt.Println(rangeBitwiseAnd(20000, 2147483647))
}
