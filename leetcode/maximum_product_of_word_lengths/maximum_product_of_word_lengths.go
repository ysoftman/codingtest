/*
https://leetcode.com/problems/maximum-product-of-word-lengths/
318. Maximum Product of Word Lengths
Medium
Given a string array words, return the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. If no such two words exist, return 0.

Example 1:
Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
Output: 16
Explanation: The two words can be "abcw", "xtfn".

Example 2:
Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
Output: 4
Explanation: The two words can be "ab", "cd".

Example 3:
Input: words = ["a","aa","aaa","aaaa"]
Output: 0
Explanation: No such pair of words.

Constraints:
2 <= words.length <= 1000
1 <= words[i].length <= 1000
words[i] consists only of lowercase English letters.
*/

package main

import "fmt"

func shareCommonLetter(w1, w2 string) bool {
	for i := 0; i < len(w1); i++ {
		for j := 0; j < len(w2); j++ {
			if w1[i] == w2[j] {
				return true
			}
		}
	}
	return false
}

// brute-force, time complexity: O(n*n*j*k)
func maxProduct2(words []string) int {
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if shareCommonLetter(words[i], words[j]) == false {
				if max < len(words[i])*len(words[j]) {
					max = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return max
}

// time complexity: O((n*n)+(n*k))
func maxProduct(words []string) int {
	bitwords := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			// a 1
			// b 10
			// c 100
			// ...
			// z 10000(26-length)
			bitwords[i] |= 1 << ((byte(words[i][j]) - 'a') + 1)
		}
	}
	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bitwords[i]&bitwords[j] == 0 {
				if max < len(words[i])*len(words[j]) {
					max = len(words[i]) * len(words[j])
				}
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	fmt.Println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}
