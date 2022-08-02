/*
https://leetcode.com/problems/find-unique-binary-string/
1980. Find Unique Binary String
Medium
Given an array of strings nums containing n unique binary strings each of length n, return a binary string of length n that does not appear in nums. If there are multiple answers, you may return any of them.

Example 1:
Input: nums = ["01","10"]
Output: "11"
Explanation: "11" does not appear in nums. "00" would also be correct.

Example 2:
Input: nums = ["00","01"]
Output: "11"
Explanation: "11" does not appear in nums. "10" would also be correct.

Example 3:
Input: nums = ["111","011","001"]
Output: "101"
Explanation: "101" does not appear in nums. "000", "010", "100", and "110" would also be correct.

Constraints:
n == nums.length
1 <= n <= 16
nums[i].length == n
nums[i] is either '0' or '1'.
All the strings of nums are unique.
*/

package main

import "fmt"

func replaceStringChar(str string, idx int, newChar string) string {
	if idx >= len(str) {
		return ""
	}
	temp := str[:idx]
	temp += newChar
	if idx+1 < len(str) {
		temp += str[idx+1:]
	}
	return string(temp)
}
func findNotExistBinStr(binstr string, startIdx int, m map[string]bool) string {
	if _, exist := m[binstr]; !exist {
		return binstr
	}

	temp := replaceStringChar(binstr, startIdx, "0")
	if len(temp) > 0 {
		if ret := findNotExistBinStr(temp, startIdx+1, m); ret != "" {
			return ret
		}
		if ret := findNotExistBinStr(replaceStringChar(binstr, startIdx, "1"), startIdx+1, m); ret != "" {
			return ret
		}
	}
	return ""
}

// 모든 경우의 수를 만들어 비교하는 방법
func findDifferentBinaryString2(nums []string) string {
	m := map[string]bool{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	temp := ""
	for i := 0; i < len(nums[0]); i++ {
		temp += "0"
	}
	return findNotExistBinStr(temp, 0, m)
}

/*
n 개의 원소들이 1~n 자리를 중 하나에 대해 반대 값을 취하는 방법
n == nums.length 로 nums 각 원소의 길이 만큼 원소가 존재
n = 2, 01 00 으로 2개의 원소
n = 3, 010 000 110 으로 3개의 원소

010 -> not 첫번째 비트 -> 1
000 -> not 두번째 비트 -> 1
010 -> not 세번째 비트 -> 1

요렇게 하면 원소에 없는 값을 만들 수 있다.
*/
func findDifferentBinaryString(nums []string) string {
	r := make([]byte, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i][i] == '0' {
			r[i] = '1'
		} else {
			r[i] = '0'
		}
	}
	return string(r)
}

func main() {
	fmt.Println(findDifferentBinaryString([]string{"01", "10"}))
	fmt.Println(findDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(findDifferentBinaryString([]string{"111", "011", "001"}))
	fmt.Println(findDifferentBinaryString([]string{"0"}))
}
