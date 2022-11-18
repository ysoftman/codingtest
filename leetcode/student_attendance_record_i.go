/*
https://leetcode.com/problems/student-attendance-record-i/
551. Student Attendance Record I
Easy
You are given a string s representing an attendance record for a student where each character signifies whether the student was absent, late, or present on that day. The record only contains the following three characters:

'A': Absent.
'L': Late.
'P': Present.
The student is eligible for an attendance award if they meet both of the following criteria:

The student was absent ('A') for strictly fewer than 2 days total.
The student was never late ('L') for 3 or more consecutive days.
Return true if the student is eligible for an attendance award, or false otherwise.

Example 1:
Input: s = "PPALLP"
Output: true
Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.

Example 2:
Input: s = "PPALLL"
Output: false
Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.

Constraints:
1 <= s.length <= 1000
s[i] is either 'A', 'L', or 'P'.
*/
package main

import "fmt"

// time complexity: O(N)
func checkRecord(s string) bool {
	cntA := 0
	cntLLL := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			// 결석 카운트
			cntA++
		}
		if i+2 < len(s) {
			// 3일 연속 지각 카운트
			if s[i:i+3] == "LLL" {
				cntLLL++
			}
		}
	}
	// 개근상 여부
	return cntA <= 1 && cntLLL == 0
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
	fmt.Println(checkRecord("AA"))
}
