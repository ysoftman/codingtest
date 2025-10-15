/*
https://leetcode.com/problems/most-common-word
819. Most Common Word
Easy

Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned. It is guaranteed there is at least one word that is not banned, and that the answer is unique.
The words in paragraph are case-insensitive and the answer should be returned in lowercase.
Note that words can not contain punctuation symbols.

Example 1:
Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
Output: "ball"
Explanation:
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"),
and that "hit" isn't the answer even though it occurs more because it is banned.

Example 2:
Input: paragraph = "a.", banned = []
Output: "a"

Constraints:
1 <= paragraph.length <= 1000
paragraph consists of English letters, space ' ', or one of the symbols: "!?',;.".
0 <= banned.length <= 100
1 <= banned[i].length <= 10
banned[i] consists of only lowercase English letters.
*/
package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	parag := []rune(strings.ToLower(paragraph))
	for i := range parag {
		if parag[i] < 'a' || parag[i] > 'z' {
			parag[i] = ' '
		}
	}

	bannedMap := make(map[string]int)
	for _, w := range banned {
		bannedMap[strings.ToLower(w)]++
	}

	result := ""
	maxCnt := 0
	m := make(map[string]int)
	p := strings.Fields(string(parag))
	for _, w := range p {
		if _, exist := bannedMap[w]; exist {
			continue
		}
		m[w]++
		if m[w] > maxCnt {
			maxCnt = m[w]
			result = w
		}
	}
	return result
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("a", []string{""}))
}
