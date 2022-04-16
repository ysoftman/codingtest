/*
https://leetcode.com/problems/restore-ip-addresses/
93. Restore IP Addresses
Medium

2649

624

Add to List

Share
A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.



Example 1:

Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]
Example 2:

Input: s = "0000"
Output: ["0.0.0.0"]
Example 3:

Input: s = "101023"
Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]


Constraints:

1 <= s.length <= 20
s consists of digits only.
*/
package main

import "fmt"

func resursiveRestoreIpAddresses(s string, i, dotcnt int, temp string, r *[]string) {
	if dotcnt > 4 {
		// return
	}
	if i >= len(s) && dotcnt == 4 {
		*r = append(*r, temp[:len(temp)-1])
		return
	}

	if i < len(s) {
		resursiveRestoreIpAddresses(s, i+1, dotcnt+1, temp+string(s[i])+".", r)
	}
	if i+1 < len(s) && s[i] != '0' {
		resursiveRestoreIpAddresses(s, i+2, dotcnt+1, temp+string(s[i])+string(s[i+1])+".", r)
	}
	if i+2 < len(s) {
		if s[i] == '1' ||
			(s[i] == '2' && s[i+1] < '5') ||
			(s[i] == '2' && s[i+1] == '5' && s[i+2] <= '5') {
			resursiveRestoreIpAddresses(s, i+3, dotcnt+1, temp+string(s[i])+string(s[i+1])+string(s[i+2])+".", r)
		}
	}
}
func restoreIpAddresses(s string) []string {
	r := []string{}
	resursiveRestoreIpAddresses(s, 0, 0, "", &r)
	return r
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
	fmt.Println(restoreIpAddresses("000256"))
	fmt.Println(restoreIpAddresses("172162541"))
}
