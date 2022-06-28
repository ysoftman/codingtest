/*
https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
1647. Minimum Deletions to Make Character Frequencies Unique
Medium
A string s is called good if there are no two different characters in s that have the same frequency.
Given a string s, return the minimum number of characters you need to delete to make s good.
The frequency of a character in a string is the number of times it appears in the string. For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.

Example 1:
Input: s = "aab"
Output: 0
Explanation: s is already good.

Example 2:
Input: s = "aaabbbcc"
Output: 2
Explanation: You can delete two 'b's resulting in the good string "aaabcc".
Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".

Example 3:
Input: s = "ceabaacb"
Output: 2
Explanation: You can delete both 'c's resulting in the good string "eabaab".
Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).

Constraints:
1 <= s.length <= 105
s contains only lowercase English letters.
*/

package main

import "fmt"

/*
case: aab
a frequency:2
b frequency:1
different frequency > good
delete cnt : 0

case: aaabbbcc
a frequency:3
b frequency:3
c frequency:2
a,b가 같은 frequency > delete b > b frequency 2
b,c가 같은 frequency > delete b > b frequency 1 > a:3 b:1 c:2 로 모두 다른 frequency > good
     또는           > delete c > c frequency 1 > a:3 b:2 c:1 로 모두 다른 frequency > good
delete cnt : 2
*/
func minDeletions(s string) int {
	deleteCnt := 0
	usedFreq := make(map[int]bool, 0)
	freqCnt := make(map[byte]int, 0)
	// 문자별 출현 빈도 파악
	for i := 0; i < len(s); i++ {
		freqCnt[s[i]]++
	}
	// fmt.Println(freqCnt)
	for i := range freqCnt {
		for freqCnt[i] > 0 {
			// 현재 빈도 값(freqCnt)이 이미 usedFreq 에 존재한다면 (이미 해당 빈도가 이전에 있었음)
			// 현재 빈도 값(freqCnt)의 문자는 delete 를 한번해야 한다.
			if _, exist := usedFreq[freqCnt[i]]; exist {
				freqCnt[i]--
				deleteCnt++
			} else {
				// 현재 빈도 값이 처음이라면 기록
				usedFreq[freqCnt[i]] = true
				break
			}
		}
	}
	return deleteCnt
}

func main() {
	fmt.Println(minDeletions("aab"))
	fmt.Println(minDeletions("aaabbbcc"))
	fmt.Println(minDeletions("ceabaacb"))
}
