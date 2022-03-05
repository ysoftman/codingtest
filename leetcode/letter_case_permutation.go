/*
https://leetcode.com/problems/letter-case-permutation/

784. Letter Case Permutation
Medium

Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
Return a list of all possible strings we could create. Return the output in any order.

Example 1:
Input: s = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]

Example 2:
Input: s = "3z4"
Output: ["3z4","3Z4"]
*/
package main

import "fmt"

/*
     ab
  ab    Ab
ab aB  Ab AB
*/
func permutate(s string, i int, result *[]string) {
	if i >= len(s) {
		*result = append(*result, s)
		return
	}
	if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
		permutate(s, i+1, result)
		// string is immutable
		temp := []byte(s)
		if s[i] >= 'a' && s[i] <= 'z' {
			temp[i] -= ('a' - 'A')
		} else {
			temp[i] += ('a' - 'A')
		}
		permutate(string(temp), i+1, result)
	} else {
		permutate(s, i+1, result)
	}

}

func letterCasePermutation(s string) []string {
	result := []string{}
	permutate(s, 0, &result)
	return result
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}
