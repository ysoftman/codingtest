/*
https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
Medium
A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.
Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.

Example 1:
Input: n = "32"
Output: 3
Explanation: 10 + 11 + 11 = 32

Example 2:
Input: n = "82734"
Output: 8

Example 3:
Input: n = "27346209830709182346"
Output: 9

Constraints:
1 <= n.length <= 105
n consists of only digits.
n does not contain any leading zeros and represents a positive integer.
*/
package main

import "fmt"

/*
deci-bianry(0과1로 이루어진 10진수)
case: 32
3 => 111 : 1이 3개 있어야 한다.
2 => 110 : 1이 2개 있어야 한다.
결국 최소 3개의 deci-binary 있어야 한다.

caes: 82734
8 => 11111111 : 1이 8개 있어야 한다.
2 => 11000000 : 1이 2개 있어야 한다.
7 => 11111110 : 1이 7개 있어야 한다.
3 => 11100000 : 1이 3개 있어야 한다.
4 => 11110000 : 1이 4개 있어야 한다.
결국 최소 8개의 deci-binary 가 있어야 한다.

결론 0~9 사이의 값 중 하나가 결과가 된다.
주어진 수의 각 자리수 중 가장 큰 수가 결과가 된다.
*/
func minPartitions(n string) int {
	r := 0
	for i := 0; i < len(n); i++ {
		if r < int(n[i]-'0') {
			r = int(n[i] - '0')
		}
	}
	return r
}

func main() {
	fmt.Println(minPartitions("32"))
	fmt.Println(minPartitions("82734"))
	fmt.Println(minPartitions("27346209830709182346"))
}
