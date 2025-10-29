/*
https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/
1209. Remove All Adjacent Duplicates in String II
Medium
You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, causing the left and the right side of the deleted substring to concatenate together.
We repeatedly make k duplicate removals on s until we no longer can.
Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.

Example 1:
Input: s = "abcd", k = 2
Output: "abcd"
Explanation: There's nothing to delete.

Example 2:
Input: s = "deeedbbcccbdaa", k = 3
Output: "aa"
Explanation:
First delete "eee" and "ccc", get "ddbbbdaa"
Then delete "bbb", get "dddaa"
Finally delete "ddd", get "aa"

Example 3:
Input: s = "pbbcggttciiippooaais", k = 2
Output: "ps"

Constraints:
1 <= s.length <= 105
2 <= k <= 104
s only contains lower case English letters.
*/
package main

import "fmt"

// time limit exceeded
func removeDuplicates2(s string, k int) string {
	removeList := []int{}
	changed := true
	for changed {
		changed = false
		pre := s[0]
		start := 0
		same := 0
		for i := 1; i < len(s); i++ {
			if s[i] == pre {
				same++
			} else {
				start = i
				pre = s[i]
				same = 0
			}
			if same == k-1 {
				removeList = append(removeList, start, i)
				i++
				if i == len(s) {
					break
				}
				start = i
				pre = s[i]
				same = 0
				changed = true
			}
		}
		// fmt.Println(removeList)
		temp := ""
		for i := 0; i < len(s); i++ {
			skip := false
			for j := 0; j < len(removeList); j += 2 {
				if removeList[j] <= i && i <= removeList[j+1] {
					skip = true
				}
			}
			if !skip {
				temp += string(s[i])
			}
		}
		// fmt.Println(temp)
		s = temp
		removeList = []int{}
	}
	return s
}

// using stack + data(char, cnt)
func removeDuplicates(s string, k int) string {
	r := ""

	type data struct {
		s   byte
		cnt int
	}
	stack := []data{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := &stack[len(stack)-1]
			if top.s == s[i] {
				top.cnt++
				if top.cnt == k {
					// pop(remove)
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, data{s[i], 1})
			}
		} else {
			stack = append(stack, data{s[i], 1})
		}
	}
	for i := 0; i < len(stack); i++ {
		for j := 0; j < stack[i].cnt; j++ {
			r += string(stack[i].s)
		}
	}
	return r
}

func main() {
	fmt.Println(removeDuplicates("abcd", 2))
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
	fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2))
	fmt.Println(removeDuplicates("fdsahfasdiyfiqwueyrewuoiyrwiqyrrfaaasiudhfihfifhffhfihfidididhhkqhqwoefhfhhfhffhheheieiewifhdskkakkakakjdfidshfiadshfihqwiehfqweifhqewkfhkaldshfkdahskdkkkkkfhdifhaqoqoerhqlhrqlkhahashshalfhadslkfhadslfhaklhfddkdhdhdiqqqqqkhdkkhkhkqhqkhqwiehqifjiafhiadshfidsahfioasdhfoahfehfoqhhfdsuihfadsohfsofhahhahqiwoeyrqwoiyroiwqyqyqyoyovvzxvnvzxnvzxnzvnxvhoiowiehiqhiqh", 5))
}
