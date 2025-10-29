/*
https://leetcode.com/problems/super-pow/
372. Super Pow
Medium
Your task is to calculate ab mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array.

Example 1:
Input: a = 2, b = [3]
Output: 8

Example 2:
Input: a = 2, b = [1,0]
Output: 1024

Example 3:
Input: a = 1, b = [4,3,3,8,5,2]
Output: 1

Constraints:
1 <= a <= 231 - 1
1 <= b.length <= 2000
0 <= b[i] <= 9
b does not contain leading zeros.
*/
package main

import "fmt"

/*
a^b % k = (a%k)(b%k)%k
ex)
a^1234567%k = (a^1234560%k) * (a^7%k)%k = (a^123456%k)^10%k * (a^7%k)%k
*/
func powmod(a, k int) int {
	a %= 1337
	result := 1
	for i := 0; i < k; i++ {
		result = (result * a) % 1337
	}
	return result
}

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	n := b[len(b)-1]
	b = b[:len(b)-1]
	return powmod(superPow(a, b), 10) * powmod(a, n) % 1337
}

func main() {
	a := 2
	b := []int{3}
	fmt.Println(superPow(a, b))
	a = 2
	b = []int{1, 0}
	fmt.Println(superPow(a, b))
	a = 1
	b = []int{4, 3, 3, 8, 5, 2}
	fmt.Println(superPow(a, b))
}
