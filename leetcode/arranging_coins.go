/*
https://leetcode.com/problems/arranging-coins/
441. Arranging Coins
Easy
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.
Given the integer n, return the number of complete rows of the staircase you will build.

Example 1:
Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.

Example 2:
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.

Constraints:
1 <= n <= 231 - 1
*/
package main

import "fmt"

// O(N)
func arrangeCoins2(n int) int {
	prerow := 1
	rowcnt := 0
	for n-prerow >= 0 {
		rowcnt++
		n -= prerow
		prerow++
	}
	return rowcnt
}

/*
n => num of complete rows(k)
1 => 1
2 => 1
3 => 2
4 => 2
5 => 2
6 => 3

row 6가 라면 다음과 같은 n 개의 코인이 있어야 한다.
1+2+3+4+5+6 = 21

complete rows 개수를 k,
필요한 코인 개수를 n,
라 할때 다음을 만족해야 한다.

(k*(k+1))/2 <= n

ex) k=3, (3(3+1))/2 = 12/2 = 6 => n (필요한 코인개수)
*/
// O(logN)
func arrangeCoins(n int) int {
	left, right := 0, n
	// binary search 로 범위를 좁혀가며 k 를 찾는다.
	for left <= right {
		mid := (right-left)/2 + left
		cur := mid * (mid + 1) / 2
		if cur == n {
			return mid
		} else if cur < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 마직막까지 진행된 right 가 만들 수 있는 k(num of complete rows)
	return right
}

func main() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(21))
	fmt.Println(arrangeCoins(150))
	fmt.Println(arrangeCoins(236))
}
