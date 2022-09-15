/*
https://leetcode.com/problems/bulls-and-cows/
299. Bulls and Cows
Medium
You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

The number of "bulls", which are digits in the guess that are in the correct position.
The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.

The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.

Example 1:
Input: secret = "1807", guess = "7810"
Output: "1A3B"
Explanation: Bulls are connected with a '|' and cows are underlined:
"1807"
  |
"7810"

Example 2:
Input: secret = "1123", guess = "0111"
Output: "1A1B"
Explanation: Bulls are connected with a '|' and cows are underlined:
"1123"        "1123"
  |      or     |
"0111"        "0111"
Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.

Constraints:
1 <= secret.length, guess.length <= 1000
secret.length == guess.length
secret and guess consist of digits only.
*/
package main

import "fmt"

// 한국에서는 야구게임(스트라이트:숫자+자리, 볼:숫자, 아웃:없음)이라고 하는데...
func getHint(secret string, guess string) string {
	secretRune := []rune(secret)
	guessRune := []rune(guess)
	bulls := 0
	for i := 0; i < len(guessRune); i++ {
		if guessRune[i] == secretRune[i] {
			bulls++
			guessRune[i] = '_'
			secretRune[i] = '_'
		}
	}

	m := make(map[rune]int)
	for i := 0; i < len(secretRune); i++ {
		m[secretRune[i]]++
	}

	cows := 0
	for i := 0; i < len(guessRune); i++ {
		if guessRune[i] == '_' {
			continue
		}
		if val, exist := m[guessRune[i]]; exist {
			if val > 0 {
				cows++
				m[guessRune[i]]--
			}
		}
	}

	return fmt.Sprintf("%vA%vB", bulls, cows)
}

func main() {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))
	secret = "1123"
	guess = "0111"
	fmt.Println(getHint(secret, guess))
}
