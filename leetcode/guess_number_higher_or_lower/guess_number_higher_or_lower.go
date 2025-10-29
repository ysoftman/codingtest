/*
https://leetcode.com/problems/guess-number-higher-or-lower/
374. Guess Number Higher or Lower
Easy
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Example 1:
Input: n = 10, pick = 6
Output: 6

Example 2:
Input: n = 1, pick = 1
Output: 1

Example 3:
Input: n = 2, pick = 1
Output: 1

Constraints:

1 <= n <= 231 - 1
1 <= pick <= n
*/
package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pickedNumber = 0

func guess(num int) int {
	if num > pickedNumber {
		return -1
	} else if num < pickedNumber {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	left, right := 0, n
	mid := 0
	for left <= right {
		mid = (right-left)/2 + left
		gnum := guess(mid)
		if gnum == -1 {
			right = right - 1
		} else if gnum == 1 {
			left = mid + 1
		} else {
			break
		}
	}
	return mid
}

func main() {
	pickedNumber = 6
	fmt.Println(guessNumber(10))
	pickedNumber = 1
	fmt.Println(guessNumber(1))
	pickedNumber = 2
	fmt.Println(guessNumber(2))
}
