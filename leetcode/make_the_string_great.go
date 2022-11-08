/*
https://leetcode.com/problems/make-the-string-great/
1544. Make The String Great
Easy

Given a string s of lower and upper case English letters.

A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
0 <= i <= s.length - 2
s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.

Return the string after making it good. The answer is guaranteed to be unique under the given constraints.

Notice that an empty string is also good.

Example 1:
Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".

Example 2:
Input: s = "abBAcC"
Output: ""
Explanation: We have many possible scenarios, and all lead to the same answer. For example:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""

Example 3:
Input: s = "s"
Output: "s"

Constraints:
1 <= s.length <= 100
s contains only lower and upper case English letters.
*/
package main

import (
	"fmt"
)

func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

// 연속하는 2 영문자가 같은데 대,소문자로 다를때 삭제
// using stack
func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}
	diff := 'a' - 'A'
	// fmt.Println(diff, reflect.TypeOf(diff))
	stack := string(s[0])
	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack += string(s[i])
			continue
		}
		top := stack[len(stack)-1]
		// stack top(마직에 추가된 글자와) 현재 s[i] 가 같은 영문자인데 대,소문자로 다른 경우 둘다 삭제
		// fmt.Println(string(top), string(s[i]), top, s[i], int32(top)-int32(s[i]))
		if abs(int32(top)-int32(s[i])) == diff {
			// pop stack
			stack = stack[:len(stack)-1]
			continue
		}
		stack += string(s[i])
	}
	return stack
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
	fmt.Println(makeGood("pP"))
	fmt.Println(makeGood("Pp"))
}
