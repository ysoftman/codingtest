/*
https://leetcode.com/problems/generate-parentheses/
22. Generate Parentheses
Medium
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Example 2:
Input: n = 1
Output: ["()"]

Constraints:
1 <= n <= 8
*/

package main

import "fmt"

/*
            (
      /           \
     (             )
    / \           /
   (   )         (
      / \      /    \
     (   )    (      )
         /     \    /
        (       )  (
*/

func dfsGenerateParenthesis(openN, closeN int, temp string, result *[]string) {
	if openN == 0 && closeN == 0 {
		*result = append(*result, temp)
	}

	if openN > 0 {
		dfsGenerateParenthesis(openN-1, closeN, temp+"(", result)
	}
	if closeN > 0 && openN < closeN {
		dfsGenerateParenthesis(openN, closeN-1, temp+")", result)
	}
}

func generateParenthesis(n int) []string {
	r := []string{}
	temp := ""
	dfsGenerateParenthesis(n, n, temp, &r)
	return r
}
func main() {
	for i := 1; i <= 8; i++ {
		fmt.Println(generateParenthesis(i))
	}
}
