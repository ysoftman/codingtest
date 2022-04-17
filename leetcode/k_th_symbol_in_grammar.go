/*
https://leetcode.com/problems/k-th-symbol-in-grammar/
779. K-th Symbol in Grammar
Medium
We build a table of n rows (1-indexed). We start by writing 0 in the 1st row. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

For example, for n = 3, the 1st row is 0, the 2nd row is 01, and the 3rd row is 0110.
Given two integer n and k, return the kth (1-indexed) symbol in the nth row of a table of n rows.

Example 1:
Input: n = 1, k = 1
Output: 0
Explanation: row 1: 0

Example 2:
Input: n = 2, k = 1
Output: 0
Explanation:
row 1: 0
row 2: 01

Example 3:
Input: n = 2, k = 2
Output: 1
Explanation:
row 1: 0
row 2: 01

Constraints:
1 <= n <= 30
1 <= k <= 2n - 1
*/
package main

import "fmt"

// brute force, time limit exceeded
func kthGrammar2(n int, k int) int {
	s := "0"
	for i := 1; i < n; i++ {
		temp := ""
		for j := 0; j < len(s); j++ {
			if s[j] == '0' {
				temp += "01"
			} else { // 1
				temp += "10"
			}
		}
		s = temp
		// fmt.Println(s)
	}

	return int(s[k-1] - '0')
}

/*
1 0 mid=0
2 01  => 짝수 반대값으로 대칭(mid=1)
3 0110 => 홀수 같은값으로 대칭(mid=2)
4 01101001 => 짝수 반대값으로 대칭(mid=4)
5 0110100110010110 => 홀수 같은값으로 대칭 (mid=8)
6 01101001100101101001011001101001 => 짝수 반대값으로 대칭(mid=16)
...

ex) n=6, k=5
n=6, k=5 => mid = 16, 5<=16 때문에, n(6-1)에서 k(5)를 찾으면된다.
n=5, k=5 => mid = 8, 5<=8 때문에, n(5-1)에서 k(5)를 찾으면된다.
n=4, k=5 => mid = 4, 5>4 때문에, n(4-1)에서 k(5-4)를 찾고 반대값을 취하면 된다.
n=3, k=1 => mid = 2, 1<=2 때문에, n(3-1)에서 k(1)를 찾으면된다.
n=2, k=1 => mid = 1, 1<=1 때문에, n(2-1)에서 k(1)를 찾으면된다.
n=1, k=0 => 0 반대값1번 => 1

ex) n=6, k=4
n=6, k=4 => mid = 16, 4<=16 때문에, n(6-1)에서 k(4)를 찾으면된다.
n=5, k=4 => mid = 8, 4<=8 때문에, n(5-1)에서 k(4)를 찾으면된다.
n=4, k=4 => mid = 4, 4<=4 때문에, n(4-1)에서 k(4)를 찾으면된다.
n=3, k=4 => mid = 2, 4>2 때문에, n(3-1)에서 k(4-2)를 찾고 반대값을 취하면 된다.
n=2, k=2 => mid = 1, 2>1 때문에, n(2-1)에서 k(2-1)를 찾고 반대값을 취하면 된다.
n=1, k=1 => 0 반대값2번 => 0

*/
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	mid := (1 << (n - 1)) / 2
	if k <= mid {
		return kthGrammar(n-1, k)
	}
	if kthGrammar(n-1, k-mid) == 0 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(kthGrammar(1, 1))
	fmt.Println(kthGrammar(2, 1))
	fmt.Println(kthGrammar(2, 2))
	fmt.Println(kthGrammar(6, 5))
	fmt.Println(kthGrammar(6, 4))
	fmt.Println(kthGrammar(30, 434991989))
}
