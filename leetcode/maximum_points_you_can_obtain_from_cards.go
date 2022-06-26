/*
https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/
1423. Maximum Points You Can Obtain from Cards
Medium
There are several cards arranged in a row, and each card has an associated number of points. The points are given in the integer array cardPoints.
In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.
Your score is the sum of the points of the cards you have taken.
Given the integer array cardPoints and the integer k, return the maximum score you can obtain.


Example 1:
Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.

Example 2:
Input: cardPoints = [2,2,2], k = 2
Output: 4
Explanation: Regardless of which two cards you take, your score will always be 4.

Example 3:
Input: cardPoints = [9,7,7,9,7,7,9], k = 7
Output: 55
Explanation: You have to take all the cards. Your score is the sum of points of all cards.

Constraints:
1 <= cardPoints.length <= 105
1 <= cardPoints[i] <= 104
1 <= k <= cardPoints.length
*/
package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	result := 0

	beginSum := []int{0}
	endSum := []int{0}
	// idx=0 또는 idx=len(cardPoints)-1 가 포함되어야 한다.
	// 앞에서 시작했을때 모든 합의 경우를 찾는다.
	for i := 0; i < len(cardPoints); i++ {
		beginSum = append(beginSum, beginSum[len(beginSum)-1]+cardPoints[i])
	}
	// 뒤쪽에서 시작했을때 모든 합의 경우를 찾는다.
	for i := len(cardPoints) - 1; i >= 0; i-- {
		endSum = append(endSum, endSum[len(endSum)-1]+cardPoints[i])
	}
	// fmt.Println("beginSum:", beginSum)
	// fmt.Println("endSum:", endSum)
	// k 길이에 대하서 앞쪽,뒤쪽 조합으로 sum 을 만들 수 있다.
	for i := 0; i <= k; i++ {
		if result < beginSum[i]+endSum[k-i] {
			result = beginSum[i] + endSum[k-i]
		}
	}
	return result
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
	fmt.Println(maxScore([]int{96, 90, 41, 82, 39, 74, 64, 50, 30}, 8))
}
