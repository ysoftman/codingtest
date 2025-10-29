/*
https://leetcode.com/problems/edit-distance/
72. Edit Distance
Hard

Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character

Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')


Constraints:
0 <= word1.length, word2.length <= 500
word1 and word2 consist of lowercase English letters.
*/

package main

import "fmt"

/*
levenshtein-distance
    h o r s e
  0 1 2 3 4 5 -> (init)
r 1 1 2 2 3 4
o 2 2 1 2 3 4
s 3 3 2 2 2 3 -> 3(distance)
  |
  v
(init)
*/

func printLevenshteinDistance(dp [][]int, word1, word2 string) {
	for i := 0; i < len(dp); i++ {
		if i > 0 {
			fmt.Print(string(word2[i-1]), " ")
		}
		if i == 0 {
			fmt.Print("    ")
			for j := 0; j < len(word1); j++ {
				fmt.Print(string(word1[j]), " ")
			}
			fmt.Println("")
			fmt.Print("  ")
		}
		for j := 0; j < len(dp[i]); j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word1)+1)
	}
	// initialize first col 0,1,2,3,...n
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	// initialize first row 0,1,2,3,...n
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			cost := 1
			if word2[i-1] == word1[j-1] {
				cost = 0
			}
			// fmt.Println(string(word2[i-1]), "-", string(word1[j-1]), "=>", cost)
			dp[i][j] = min(min(dp[i-1][j]+1, // remove
				dp[i][j-1]+1), // insert
				dp[i-1][j-1]+cost) // replace
		}
	}
	printLevenshteinDistance(dp, word1, word2)
	return dp[len(word2)][len(word1)]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
