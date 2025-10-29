/*
https://leetcode.com/problems/sort-characters-by-frequency/
451. Sort Characters By Frequency
Medium

Given a string s, sort it in decreasing order based on the frequency of the characters. The frequency of a character is the number of times it appears in the string.
Return the sorted string. If there are multiple answers, return any of them.

Example 1:
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.

Example 2:
Input: s = "cccaaa"
Output: "aaaccc"
Explanation: Both 'c' and 'a' appear three times, so both "cccaaa" and "aaaccc" are valid answers.
Note that "cacaca" is incorrect, as the same characters must be together.

Example 3:
Input: s = "Aabb"
Output: "bbAa"
Explanation: "bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.

Constraints:
1 <= s.length <= 5 * 105
s consists of uppercase and lowercase English letters and digits.
*/
package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	frequencies := []byte{}

	// map key(frequency)만 slice(frequencies) 로 구성
	for k := range m {
		frequencies = append(frequencies, k)
	}
	// frequencies 를 map value 기준으로 정렬
	sort.Slice(frequencies, func(a, b int) bool {
		return m[frequencies[a]] > m[frequencies[b]]
	})

	r := ""
	for i := 0; i < len(frequencies); i++ {
		cnt := m[frequencies[i]]
		for j := 0; j < cnt; j++ {
			r += string(frequencies[i])
		}
	}
	return r
}

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
