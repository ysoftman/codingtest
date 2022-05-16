/*
https://leetcode.com/problems/binary-watch/
401. Binary Watch
Easy
A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59). Each LED represents a zero or one, with the least significant bit on the right.

For example, the below binary watch reads "4:51".

Given an integer turnedOn which represents the number of LEDs that are currently on, return all possible times the watch could represent. You may return the answer in any order.

The hour must not contain a leading zero.

For example, "01:00" is not valid. It should be "1:00".
The minute must be consist of two digits and may contain a leading zero.

For example, "10:2" is not valid. It should be "10:02".

Example 1:
Input: turnedOn = 1
Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

Example 2:
Input: turnedOn = 9
Output: []


Constraints:

0 <= turnedOn <= 10
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findTime(r *[]string, turnedOn int, bintimestr []byte, startIdx int) {
	if turnedOn == 0 {
		hourtime, _ := strconv.ParseInt(string(bintimestr[:4]), 2, 32)
		if hourtime >= 12 {
			return
		}
		mintime, _ := strconv.ParseInt(string(bintimestr[4:]), 2, 32)
		if mintime >= 60 {
			return
		}
		curtime := fmt.Sprintf("%d:%02d", hourtime, mintime)
		*r = append(*r, curtime)
		return
	}
	for i := startIdx; i < len(bintimestr); i++ {
		bintimestr[i] = '1'
		findTime(r, turnedOn-1, bintimestr, i+1)
		bintimestr[i] = '0'
	}
}

// combination 으로 처리
func readBinaryWatch2(turnedOn int) []string {
	r := []string{}
	bintimestr := []byte("0000000000")
	findTime(&r, turnedOn, bintimestr, 0)
	// for debugging
	// sort.Slice(r, func(a, b int) bool {
	// 	aa, _ := strconv.ParseInt(strings.Split(string(r[a]), ":")[0], 10, 32)
	// 	bb, _ := strconv.ParseInt(strings.Split(string(r[b]), ":")[0], 10, 32)
	// 	cc, _ := strconv.ParseInt(strings.Split(string(r[a]), ":")[1], 10, 32)
	// 	dd, _ := strconv.ParseInt(strings.Split(string(r[b]), ":")[1], 10, 32)
	// 	if aa < bb {
	// 		return true
	// 	}
	// 	if aa == bb && cc < dd {
	// 		return true
	// 	}
	// 	return false
	// })
	return r
}

// 단순하게 모든 시간의 경우에 대해서 turnedOn 체크
func readBinaryWatch(turnedOn int) []string {
	r := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			temp := strconv.FormatInt(int64(i), 2) + strconv.FormatInt(int64(j), 2)
			if strings.Count(temp, "1") == turnedOn {
				r = append(r, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return r
}

func main() {
	fmt.Println(readBinaryWatch(1))
	fmt.Println(readBinaryWatch(2))
	fmt.Println(readBinaryWatch(3))
	fmt.Println(readBinaryWatch(4))
	fmt.Println(readBinaryWatch(5))
}
