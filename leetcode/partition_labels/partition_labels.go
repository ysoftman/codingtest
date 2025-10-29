/*
https://leetcode.com/problems/partition-labels/
763. Partition Labels
Medium
You are given a string s. We want to partition the string into as many parts as possible so that each letter appears in at most one part.
Note that the partition is done so that after concatenating all the parts in order, the resultant string should be s.
Return a list of integers representing the size of these parts.

Example 1:
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits s into less parts.

Example 2:
Input: s = "eccbbbbdec"
Output: [10]

Constraints:
1 <= s.length <= 500
s consists of lowercase English letters.
*/
package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func partitionLabels(s string) []int {
	m := make(map[byte]int, 0)
	// 각 문자의 최대 위치를 기록
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}
	// fmt.Println("--", m)

	result := []int{}
	pre := -1
	maxi := 0
	for i := 0; i < len(s); i++ {
		// 현재 문자의 최대 위치 파악
		maxi = max(maxi, m[s[i]])
		// 현재 문자의 최대 위치까지 왔다면 나눌 수 있다.
		if maxi == i {
			result = append(result, maxi-pre)
			pre = maxi
		}
	}
	return result
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}
