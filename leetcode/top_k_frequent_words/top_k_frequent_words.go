/*
https://leetcode.com/problems/top-k-frequent-words/
692. Top K Frequent Words
Medium

Given an array of strings words and an integer k, return the k most frequent strings.

Return the answer sorted by the frequency from highest to lowest. Sort the words with the same frequency by their lexicographical order.

Example 1:
Input: words = ["i","love","leetcode","i","love","coding"], k = 2
Output: ["i","love"]
Explanation: "i" and "love" are the two most frequent words.
Note that "i" comes before "love" due to a lower alphabetical order.

Example 2:
Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
Output: ["the","is","sunny","day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.


Constraints:
1 <= words.length <= 500
1 <= words[i].length <= 10
words[i] consists of lowercase English letters.
k is in the range [1, The number of unique words[i]]

Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?
*/

package main

import (
	"fmt"
	"sort"
)

type data struct {
	word string
	cnt  int
}

// func priority(a, b interface{}) int {
// 	// cnt 순으로 우선순위 결정
// 	pa := a.(data).cnt
// 	pb := b.(data).cnt
// 	// 같은 빈도의 단어라면 사전우선순위
// 	if pa == pb {
// 		paWord := a.(data).word
// 		pbWord := b.(data).word
// 		return utils.StringComparator(paWord, pbWord)
// 	}
// 	// - 하면 내림차순
// 	return -utils.IntComparator(pa, pb)
// }

// // gods 패키지 사용
// // https://godoc.org/github.com/emirpasic/gods
// // go install github.com/emirpasic/gods@latest
// // 우선 순위 큐 사용하기
// func topKFrequent2(words []string, k int) []string {
// 	// leetcode 에서 안되는것 같음!!!
// 	q := priorityqueue.NewWith(priority)
// 	defer q.Clear()
// 	m := make(map[string]int)
// 	// 단어별 빈도 카운트
// 	for i := 0; i < len(words); i++ {
// 		m[words[i]]++
// 	}
// 	// priorityqueue 에 넣어 높은 빈도의 단어가 트리의 위로 되도록 한다.
// 	for k, v := range m {
// 		d := data{
// 			word: k,
// 			cnt:  v,
// 		}
// 		q.Enqueue(d)
// 	}

// 	r := []string{}
// 	for i := 0; i < k; i++ {
// 		val, exist := q.Peek()
// 		if !exist {
// 			break
// 		}
// 		// fmt.Println(val)
// 		r = append(r, val.(data).word)

// 		q.Dequeue()
// 	}
// 	return r
// }

//////////

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	// 단어별 빈도 카운트
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	// 정렬을 위해 data 리스트로 옮기기
	datas := []data{}
	for k, v := range m {
		d := data{
			word: k,
			cnt:  v,
		}
		datas = append(datas, d)
	}
	sort.Slice(datas, func(a, b int) bool {
		if datas[a].cnt == datas[b].cnt {
			return datas[a].word < datas[b].word
		}
		return datas[a].cnt > datas[b].cnt
	})
	r := []string{}
	for i := 0; i < len(datas); i++ {
		if i >= k {
			break
		}
		r = append(r, datas[i].word)
	}
	return r
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
