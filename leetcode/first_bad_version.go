/*
https://leetcode.com/problems/first-bad-version/
278. First Bad Version
Easy
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.
Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

Example 1:
Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.

Example 2:
Input: n = 1, bad = 1
Output: 1
*/
package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 테스트를 위한 badVersion 설정
var badVersion int = 0

// 테스트를 위한 isBadVersion() 함수 구현
func isBadVersion(n int) bool {
	if n == badVersion {
		return true
	}
	return false
}
func firstBadVersion(n int) int {
	left, right := 1, n
	badVersion, goodVersion := 0, 0
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			badVersion = mid
			right = mid - 1
		} else {
			goodVersion = mid
			left = mid + 1
		}
		if badVersion-1 == goodVersion {
			return badVersion
		}
	}
	return -1
}

func main() {
	badVersion = 4
	fmt.Println(firstBadVersion(5))
	badVersion = 1
	fmt.Println(firstBadVersion(1))
}
