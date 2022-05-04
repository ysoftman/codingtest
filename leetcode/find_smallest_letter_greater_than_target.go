/*
https://leetcode.com/problems/find-smallest-letter-greater-than-target/
744. Find Smallest Letter Greater Than Target
Easy
Given a characters array letters that is sorted in non-decreasing order and a character target, return the smallest character in the array that is larger than target.
Note that the letters wrap around.
For example, if target == 'z' and letters == ['a', 'b'], the answer is 'a'.

Example 1:
Input: letters = ["c","f","j"], target = "a"
Output: "c"

Example 2:
Input: letters = ["c","f","j"], target = "c"
Output: "f"

Example 3:
Input: letters = ["c","f","j"], target = "d"
Output: "f"

Constraints:
2 <= letters.length <= 104
letters[i] is a lowercase English letter.
letters is sorted in non-decreasing order.
letters contains at least two different characters.
target is a lowercase English letter.
*/

package main

import "fmt"

func nextGreatestLetter2(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}

// using binary search O(logN)
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	find := 0
	for left <= right {
		mid := (right-left)/2 + left
		if letters[mid] > target {
			find = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return letters[find]
}

func main() {
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'a', 'b'}, 'z'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
}
