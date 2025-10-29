/*
https://leetcode.com/problems/magical-string/
481. Magical String
Medium

A magical string s consists of only '1' and '2' and obeys the following rules:

The string s is magical because concatenating the number of contiguous occurrences of characters '1' and '2' generates the string s itself.
The first few elements of s is s = "1221121221221121122……". If we group the consecutive 1's and 2's in s, it will be "1 22 11 2 1 22 1 22 11 2 11 22 ......" and the occurrences of 1's or 2's in each group are "1 2 2 1 1 2 1 2 2 1 2 2 ......". You can see that the occurrence sequence is s itself.

Given an integer n, return the number of 1's in the first n number in the magical string s.

Example 1:
Input: n = 6
Output: 3
Explanation: The first 6 elements of magical string s is "122112" and it contains three 1's, so return 3.

Example 2:
Input: n = 1
Output: 1

Constraints:
1 <= n <= 105
*/

package main

import "fmt"

/*
최초 magic string = 1221121221221121122... 연속수가 있을때 1,2로 각각 그룹을 지으면
1 22 11 2 1 22 1 22 11 2 11 22... 이 상태에서 각 그룹내의 개수를 세보면
1개 2개 2개 1개 1개 2개 ...
1 2 2 1 1 2 1 2 2 1 2 2... 로 원래 연속되는 수가 된다.
magic string 이 n까지 있을때 1의 개수를 리턴
n = 6 이면 1 2 2 1 1 2 로 1이 3번 나온다.
*/
func magicalString(n int) int {
	if n <= 3 {
		// n=3번째까지는 122 로 1은 1번 나온다.
		return 1
	}

	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 2

	LastIndex := 2      // 현재 생성된 수의 마지막 인덱스
	newNumberIndex := 3 // 새로운 수가 들어갈 인덱스
	num := 1
	result := 1

	for newNumberIndex < n {
		// 0~lastindex의 수(값)를 참고해서 newNumberIndex 에 새로운 수를 넣는다.
		// arr[LastIndex] 는 1 또는 2
		for i := 0; i < arr[LastIndex]; i++ {
			arr[newNumberIndex] = num
			// 생성된 새로운 수가 1이라면 카운트 한다.
			if num == 1 && newNumberIndex < n {
				result++
			}
			newNumberIndex++
		}

		// num 1 일 경우 --> 01(1) xor 11(3) = 10(2)
		// num 2 일 경우 --> 10(2) xor 11(3) = 01(1)
		// num = num ^ 3 로 해도 된다.
		// num = num ^ 3
		if num == 1 {
			num = 2
		} else if num == 2 {
			num = 1
		}

		LastIndex++
	}
	return result
}

func main() {
	fmt.Println(magicalString(6))
	fmt.Println(magicalString(1))
	fmt.Println(magicalString(10))
}
