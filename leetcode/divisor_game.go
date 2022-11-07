/*
https://leetcode.com/problems/divisor-game/
1025. Divisor Game
Easy

Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number n on the chalkboard. On each player's turn, that player makes a move consisting of:

Choosing any x with 0 < x < n and n % x == 0.
Replacing the number n on the chalkboard with n - x.
Also, if a player cannot make a move, they lose the game.

Return true if and only if Alice wins the game, assuming both players play optimally.

Example 1:
Input: n = 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.

Example 2:
Input: n = 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.

Constraints:
1 <= n <= 1000
*/
package main

import "fmt"

/*
n=10 (1~9)
alice: 10%1 ==0, n=10-1=9
bob:   9%3 ==0, n=9-3=6
alice: 6%2 ==0, n=6-2=4
box: X --> alice wins
*/
// mathmatical solution...
func divisorGame_mathmatical(n int) bool {
	return n%2 == 0
}

// brute-force --> Time Limit Exceeded
func divisorGame_brute_force_time_limit_exceeded(n int) bool {
	for x := 1; x < n; x++ {
		// n%x == 0 -> 현재턴 alice
		// divisorGame(n-x) -> 다음턴 bob 결과, 실패하면 alice 가 이기는것임
		if n%x == 0 && !divisorGame(n-x) {
			return true
		}
	}
	return false
}

// dynamic programming and recursive
func divisorRecursive(dp map[int]int, n int) bool {
	if dp[n] != 0 {
		return dp[n] == 1
	}
	for x := 1; x <= n/2; x++ {
		if dp[n] == 1 {
			break
		}
		// 놓을자리가 없으면 alice 패배
		if n%x != 0 {
			dp[n] = 2
			continue
		}
		// bob 차례에서 bob 이 이기면 alice 가 패배
		if divisorRecursive(dp, n-1) {
			dp[n] = 2
			continue
		}
		// 이외는 alice 승리
		dp[n] = 1
	}
	return dp[n] == 1
}

func divisorGame(n int) bool {
	// alice 가 n 일때에서의 승리,패배 결과를 dp 로 저장
	dp := make(map[int]int, 1001)
	return divisorRecursive(dp, n)
}

func main() {
	fmt.Println(divisorGame(2), divisorGame_mathmatical(2))
	fmt.Println(divisorGame(3), divisorGame_mathmatical(3))
	fmt.Println(divisorGame(10), divisorGame_mathmatical(10))
	fmt.Println(divisorGame(100), divisorGame_mathmatical(100))
}
