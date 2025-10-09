/*
https://leetcode.com/problems/groups-of-special-equivalent-strings
893. Groups of Special-Equivalent Strings
Medium
You are given an array of strings of the same length words.

In one move, you can swap any two even indexed characters or any two odd indexed characters of a string words[i].

Two strings words[i] and words[j] are special-equivalent if after any number of moves, words[i] == words[j].

For example, words[i] = "zzxy" and words[j] = "xyzz" are special-equivalent because we may make the moves "zzxy" -> "xzzy" -> "xyzz".
A group of special-equivalent strings from words is a non-empty subset of words such that:

Every pair of strings in the group are special equivalent, and
The group is the largest size possible (i.e., there is not a string words[i] not in the group such that words[i] is special-equivalent to every string in the group).
Return the number of groups of special-equivalent strings from words.

Example 1:
Input: words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
Output: 3
Explanation:
One group is ["abcd", "cdab", "cbad"], since they are all pairwise special equivalent, and none of the other strings is all pairwise special equivalent to these.
The other two groups are ["xyzz", "zzxy"] and ["zzyx"].
Note that in particular, "zzxy" is not special equivalent to "zzyx".

Example 2:
Input: words = ["abc","acb","bac","bca","cab","cba"]
Output: 3

Constraints:
1 <= words.length <= 1000
1 <= words[i].length <= 20
words[i] consist of lowercase English letters.
All the strings are of the same length.
*/
package main

import (
	"fmt"
	"sort"
)

/*
words = ["abc","acb","bac","bca","cab","cba"]
group1 => ["abc", cba -> "abc"]
group2 => ["acb", bca -> "acb"]
group3 => ["bac", cab -> "bac"]
*/
func numSpecialEquivGroups(words []string) int {
	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		// words[i](string)은 한글자를 수정해서 저장할 수 없어서 rune 타입으로 사용
		word := []rune(words[i])
		evenRunes := []rune{}
		oddRunes := []rune{}

		for j := 0; j < len(word); j++ {
			if j%2 == 0 {
				evenRunes = append(evenRunes, word[j])
			} else {
				oddRunes = append(oddRunes, word[j])
			}
		}

		// ascending sort english letters
		sort.Slice(evenRunes, func(a, b int) bool {
			return evenRunes[a] > evenRunes[b]
		})
		sort.Slice(oddRunes, func(a, b int) bool {
			return oddRunes[a] > oddRunes[b]
		})

		m[string(evenRunes)+string(oddRunes)]++
	}
	fmt.Println(m)
	return len(m)
}

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	fmt.Println(numSpecialEquivGroups([]string{"couxuxaubw", "zsptcwcghr", "kkntvvhbcc", "nkhtcvvckb", "crcwhspgzt"}))
}
