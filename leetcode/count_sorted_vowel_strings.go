/*
https://leetcode.com/problems/count-sorted-vowel-strings/
1641. Count Sorted Vowel Strings
Medium
Given an integer n, return the number of strings of length n that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.

A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.

Example 1:
Input: n = 1
Output: 5
Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].

Example 2:
Input: n = 2
Output: 15
Explanation: The 15 sorted strings that consist of vowels only are
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

Example 3:
Input: n = 33
Output: 66045

Constraints:
1 <= n <= 50
*/
package main

import "fmt"

// accepted but slow...
func recursiveMakeVowelString(n int, temp string, r *[]string) {
	if n == 0 {
		*r = append(*r, temp)
		return
	}
	v := []string{"a", "e", "i", "o", "u"}
	for j := 0; j < len(v); j++ {
		// ex) ...ea -> s[i](e)가 s[i+1](a) 보다 크면 안된다.
		if len(temp) > 0 && string(temp[len(temp)-1]) > v[j] {
			continue
		}
		recursiveMakeVowelString(n-1, temp+v[j], r)
	}
}

func countVowelStrings2(n int) int {
	r := []string{}
	recursiveMakeVowelString(n, "", &r)
	// fmt.Println(r)
	return len(r)
}

// combination
func countVowelStrings3(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}

// dynamic programming
/**
이전 문자 => 이전문자보다 큰 문자만 추가
a => aa, ae, ai, ao, au => (a+=e+i+o+u)
e => ee, ei, eo, eu => (e+=i+o+u)
i => ii, io, iu => (i+=o+u)
o => oo, ou => (o+=u)
u => uu => (u+=0)
*/
func countVowelStrings(n int) int {
	// a,e,i,o,u
	dp := make([]int, 5)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	// n = 1 일때는 위에 이미 저장되어 있음(1+1+1+1+1)
	for n > 1 {
		for i := 0; i < 5; i++ {
			for j := i + 1; j < 5; j++ {
				dp[i] += dp[j]
			}
		}
		n--
	}
	sum := 0
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
	}
	return sum
}

func main() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(3))
	fmt.Println(countVowelStrings(4))
	fmt.Println(countVowelStrings(5))
	fmt.Println(countVowelStrings(33))
}
