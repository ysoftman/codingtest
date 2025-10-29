/*
https://leetcode.com/problems/h-index-ii/
275. H-Index II
Medium

Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper and citations is sorted in an ascending order, return compute the researcher's h-index.

According to the definition of h-index on Wikipedia: A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations each.
If there are several possible values for h, the maximum one is taken as the h-index.

You must write an algorithm that runs in logarithmic time.

Example 1:
Input: citations = [0,1,3,5,6]
Output: 3
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had received 0, 1, 3, 5, 6 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:
Input: citations = [1,2,100]
Output: 2

Constraints:
n == citations.length
1 <= n <= 105
0 <= citations[i] <= 1000
citations is sorted in ascending order.
*/

package main

import "fmt"

// h-index 는 연구자(researcher)의 영향력(인용, citation) 지표다.
// n 개의 논문(paper) 중 h번 이상 인용된 논문이 h개 있고,
// 나머지 논문이 h번 이하 인용된경우
// h 최대값이 연구자의 h-index 이다.

// time complexity:  O(logN)
// logarithmic time 에 결과를 도출하기 위해 binary search
func hIndex(citations []int) int {
	n := len(citations)
	l := 0
	// r := n-1 --> [0] 처럼 원소가 1개인인데, 인용회수가 0인 경우 n-r 리턴시 0이 아닌 1(잘못된값)로 리턴됨
	r := n
	for l < r {
		mid := ((r - l) / 2) + l
		// citations[mid](현재 인용 회수)가 n-mid(전체 논문수-현재 논문 수)를 넘는 지점이지 체크
		if citations[mid] < n-mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return n - r
}

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{1, 2, 100}))
	fmt.Println(hIndex([]int{0}))
	fmt.Println(hIndex([]int{1}))
}
