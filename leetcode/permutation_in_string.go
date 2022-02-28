/*
https://leetcode.com/problems/permutation-in-string/
567. Permutation in String
Medium

5342

147

Add to List

Share
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.



Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/
package main

import "fmt"

func permute(str []byte, i, j int, result *[][]byte) {
	if i == j {
		temp := make([]byte, len(str))
		copy(temp, str)
		*result = append(*result, temp)
		return
	}
	for idx := i; idx <= j; idx++ {
		str[i], str[idx] = str[idx], str[i]
		permute(str, i+1, j, result)
		str[i], str[idx] = str[idx], str[i]
	}
}

// too long time, failed in submissions
func checkInclusion2(s1 string, s2 string) bool {
	result := make([][]byte, 0)
	permute([]byte(s1), 0, len(s1)-1, &result)
	for _, v := range result {
		for i := 0; i <= len(s2)-len(v); i++ {
			// fmt.Println(string(v), "==" , s2[i:i+len(v)])
			if string(v) == s2[i:i+len(v)] {
				return true
			}
		}
	}
	return false

}

func copyMap(src, dst map[byte]int) {
	for k, v := range dst {
		src[k] = v
	}
}
func checkAllZero(temp map[byte]int) bool {
	for _, v := range temp {
		if v != 0 {
			return false
		}
	}
	return true
}

// don't permute, just compare with hashmap of s1
func checkInclusion(s1 string, s2 string) bool {
	s1map := make(map[byte]int, len(s1))
	i := 0
	for _, v := range s1 {
		s1map[byte(v)]++
		i++
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		temp := make(map[byte]int, len(s1))
		copyMap(temp, s1map)

		for j := i; j < i+len(s1); j++ {
			if _, found := temp[byte(s2[j])]; found {
				temp[byte(s2[j])]--
			}
		}
		if checkAllZero(temp) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("prosperity", "properties"))
}
