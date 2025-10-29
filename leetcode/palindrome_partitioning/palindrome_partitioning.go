/*
https://leetcode.com/problems/palindrome-partitioning/
131. Palindrome Partitioning
Medium
Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

A palindrome string is a string that reads the same backward as forward.

Example 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:
Input: s = "a"
Output: [["a"]]

Constraints:
1 <= s.length <= 16
s contains only lowercase English letters.
*/
package main

import "fmt"

func is_palindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
            aab
    a         aa          aab(x)
  a  ab(x)   b
b
dfs 로 palindrome 을 찾는다.
dfs 시 index 로 현재 탐색중인 범위를 지정한다.
*/
func dfs_partition(s string, idx int, r []string, results *[][]string) {
	if idx == len(s) {
		result := make([]string, len(r))
		copy(result, r)
		*results = append(*results, result)
		return
	}
	for i := idx; i < len(s); i++ {
		if is_palindrome(s[idx : i+1]) {
			// palindrome 을 저장하고 i 이후의 스트림에 대해서 다시 검사
			dfs_partition(s, i+1, append(r, s[idx:i+1]), results)
		}
	}
}

func partition(s string) [][]string {
	result := [][]string{}
	dfs_partition(s, 0, []string{}, &result)
	return result
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("ababab"))
}
