/*
https://leetcode.com/problems/elimination-game/
390. Elimination Game
Medium

You have a list arr of all integers in the range [1, n] sorted in a strictly increasing order. Apply the following algorithm on arr:

Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
Repeat the previous step again, but this time from right to left, remove the rightmost number and every other number from the remaining numbers.
Keep repeating the steps again, alternating left to right and right to left, until a single number remains.
Given the integer n, return the last number that remains in arr.


Example 1:
Input: n = 9
Output: 6
Explanation:
arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]
arr = [2, 4, 6, 8]
arr = [2, 6]
arr = [6]

Example 2:
Input: n = 1
Output: 1

Constraints:
1 <= n <= 109
*/

package main

import "fmt"

/*
1~n 까지 숫자가 빠진 숫자 없이 정렬된 배열이 있을때
step1, left -> right,  왼쪽 첫번째 시작해서 오른쪽 끝까지 1,3,5,7,9 삭제
step2, right -> left,  오른쪽 끝에서 왼쪽으로 8, 4 삭제
원소1개가 남을때까지 1,2 반복
*/
func lastRemaining(n int) int {
	left := true
	remain := n
	step := 1
	head := 1
	for remain > 1 {
		// left -> right 스텝이 증가할때마다 head 의 위치가 오른쪽으로 이동(증가)한다.
		if left || remain%2 == 1 {
			head += step
		}
		// 스텝마다 남은 원소는 1/2로 줄어든다.
		remain /= 2
		step *= 2
		left = !left
	}
	return head
}

func main() {
	fmt.Println(lastRemaining(9))
	fmt.Println(lastRemaining(1))
}
