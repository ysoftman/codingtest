/*
https://leetcode.com/problems/shuffle-string/
1528. Shuffle String
Easy
You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.

Example 1:
Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
Output: "leetcode"
Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

Example 2:
Input: s = "abc", indices = [0,1,2]
Output: "abc"
Explanation: After shuffling, each character remains in its position.

Constraints:
s.length == indices.length == n
1 <= n <= 100
s consists of only lowercase English letters.
0 <= indices[i] < n
All values of indices are unique.
*/
package main

import "fmt"

// time complexity: O(n)
// space complexity: O(n)
func restoreString2(s string, indices []int) string {
	r := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		r[indices[i]] = s[i]
	}
	fmt.Println("indices:", indices)
	return string(r)
}

// time complexity: O(n)
// space complexity: O(1)
func restoreString(s string, indices []int) string {
	// golang 에서는 string immutable 이라, byte slice 로 변경한다.
	// 다른 언어에서는 이과정이 없어 이부분은 시간 복잡도에서 제외했음
	bstr := make([]byte, len(s))
	copy(bstr, s)
	fmt.Println("bstr:", string(bstr))
	for i := 0; i < len(indices); i++ {
		// 현재 indices[i] 가 제자리를 찾을때까지 스왑 반복
		for indices[i] != i {
			// swap
			bstr[indices[i]], bstr[i] = bstr[i], bstr[indices[i]]
			// 이미 스왑으로 자리를 잡은 문자가 다시 스왑되지 않도록 한번 스왑된면 해당 indices 값도 스왑해준다.
			fmt.Println("indices[i], indices[indices[i]] :", indices[i], indices[indices[i]])
			indices[i], indices[indices[i]] = indices[indices[i]], indices[i]
		}
	}
	// fmt.Println("indices:", indices)
	return string(bstr)
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreString("abc", []int{0, 1, 2}))
	fmt.Println(restoreString("aiohn", []int{3, 1, 4, 2, 0}))
}
