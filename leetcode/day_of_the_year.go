/*
1154. Day of the Year
Easy
Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

Example 1:
Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.

Example 2:
Input: date = "2019-02-10"
Output: 41

Constraints:
date.length == 10
date[4] == date[7] == '-', and all other date[i]'s are digits
date represents a calendar date between Jan 1st, 1900 and Dec 31th, 2019.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	sp := strings.Split(date, "-")
	year, _ := strconv.Atoi(sp[0])
	month, _ := strconv.Atoi(sp[1])
	days, _ := strconv.Atoi(sp[2])
	if month > 1 {
		for i := 0; i < month-1; i++ {
			days += monthDays[i]
		}
	}
	if month <= 2 {
		return days
	}
	// 윤달(2월 29일) 체크
	// 4년주기중에서 100년주기를 제외 or 400년주기
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		days++
	}
	return days
}

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
	fmt.Println(dayOfYear("2019-02-10"))
	fmt.Println(dayOfYear("2003-03-01"))
	fmt.Println(dayOfYear("2004-03-01"))
	fmt.Println(dayOfYear("2012-01-02"))
	fmt.Println(dayOfYear("2004-03-01"))
	fmt.Println(dayOfYear("1900-05-02"))
	fmt.Println(dayOfYear("2000-12-04"))
}
