/*
https://leetcode.com/problems/h-index/
274. H-Index
Medium

Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return compute the researcher's h-index.

According to the definition of h-index on Wikipedia: A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations each.
If there are several possible values for h, the maximum one is taken as the h-index.

Example 1:
Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

Example 2:
Input: citations = [1,3,1]
Output: 1

Constraints:
n == citations.length
1 <= n <= 5000
0 <= citations[i] <= 1000
*/
package main

import "fmt"

// h-index 는 연구자(researcher)의 영향력(인용, citation) 지표다.
// n 개의 논문(paper) 중 h번 이상 인용된 논문이 h개 있고,
// 나머지 논문이 h번 이하 인용된경우
// h 최대값이 연구자의 h-index 이다.
/*
bucket sorting 문제
[3,0,6,1,5] => 5개의 논문의 인용횟수
논문 n(5)개중 n(5)번 이상 인용된 논문수는 6, 5로 [5]=2개
나머지 논문은 그대로 1개씩 카운트 [i]=1
버킷으로 0~(4+1) 표현하면 [1 1 0 1 0 2]
인용회수은 마지막 인덱스 부터 합산 2+0+1... 하면서
합산이 인덱스(i) 보다 크거나 같은 순간이 결과(h-index)
*/
func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= n {
			// n 보다 큰 원소의 개 카운트
			buckets[n]++
		} else {
			buckets[citations[i]]++
		}
	}
	// fmt.Println("buckets:", buckets)

	cnt := 0
	for i := n; i >= 0; i-- {
		cnt += buckets[i]
		if cnt >= i {
			// bucket 큰 원소의 합산이 index(i)보다 커지는 순간이 h-index
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 3, 1}))
}
