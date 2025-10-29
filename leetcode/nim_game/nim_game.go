/*
https://leetcode.com/problems/nim-game/
292. Nim Game
Easy

You are playing the following Nim Game with your friend:

Initially, there is a heap of stones on the table.
You and your friend will alternate taking turns, and you go first.
On each turn, the person whose turn it is will remove 1 to 3 stones from the heap.
The one who removes the last stone is the winner.
Given n, the number of stones in the heap, return true if you can win the game assuming both you and your friend play optimally, otherwise return false.

Example 1:
Input: n = 4
Output: false
Explanation: These are the possible outcomes:
1. You remove 1 stone. Your friend removes 3 stones, including the last stone. Your friend wins.
2. You remove 2 stones. Your friend removes 2 stones, including the last stone. Your friend wins.
3. You remove 3 stones. Your friend removes the last stone. Your friend wins.
In all outcomes, your friend wins.

Example 2:
Input: n = 1
Output: true

Example 3:
Input: n = 2
Output: true

Constraints:
1 <= n <= 231 - 1
*/

package main

import "fmt"

/*
이기는 경우 : true
상대방은 이길수 있는 최적의 방법으로 플레이한다고 가정.
[1] = true
[2] = true
[3] = true
[4] = 어떤것을 선택해도 진다.
[5] = 1선택하면 이긴다.
[6] = 2선택하면 이긴다.
[7] = 3선택하면 이긴다.
[8] = 어떤것을 선택해도 진다.
*/
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canWinNim(4))
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(2))
	fmt.Println(canWinNim(8))
	fmt.Println(canWinNim(22))
	fmt.Println(canWinNim(35))
	fmt.Println(canWinNim(3241))
	fmt.Println(canWinNim(4321))
	fmt.Println(canWinNim(43255))
	fmt.Println(canWinNim(233564))
	fmt.Println(canWinNim(363456))
	fmt.Println(canWinNim(344325236))
}
