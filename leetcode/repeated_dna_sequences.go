/*
https://leetcode.com/problems/repeated-dna-sequences/
The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

For example, "ACGAATTCCG" is a DNA sequence.
When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.

Example 1:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]

Example 2:
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]

Constraints:
1 <= s.length <= 105
s[i] is either 'A', 'C', 'G', or 'T'.
*/
package main

import "fmt"

// using hashmap
func findRepeatedDnaSequences(s string) []string {
	result := []string{}
	seen := make(map[string]int, 0)
	// 10개 길이 substring으로 오른쪽으로 이동하면 hashmap 에 카운트해놓는다.
	for i := 0; i+9 < len(s); i++ {
		tenstr := s[i : i+10]
		seen[tenstr]++
	}
	for k, v := range seen {
		if v >= 2 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
