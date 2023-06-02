/*
1108. Defanging an IP Address
Easy
Given a valid (IPv4) IP address,
return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

Example 1:
Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"

Example 2:
Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"

Constraints:
The given address is a valid IPv4 address.
*/

package main

import "fmt"

func defangIPaddr(address string) string {
	out := make([]byte, len(address)+6)
	j := 0
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			out[j] = '['
			j++
			out[j] = '.'
			j++
			out[j] = ']'
		} else {
			out[j] = address[i]
		}
		j++
	}
	return string(out)
}
func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}
