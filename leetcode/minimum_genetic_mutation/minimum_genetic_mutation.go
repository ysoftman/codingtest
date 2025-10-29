/*
https://leetcode.com/problems/minimum-genetic-mutation/
433. Minimum Genetic Mutation
Medium

A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.

Suppose we need to investigate a mutation from a gene string start to a gene string end where one mutation is defined as one single character changed in the gene string.

For example, "AACCGGTT" --> "AACCGGTA" is one mutation.
There is also a gene bank bank that records all the valid gene mutations. A gene must be in bank to make it a valid gene string.

Given the two gene strings start and end and the gene bank bank, return the minimum number of mutations needed to mutate from start to end. If there is no such a mutation, return -1.

Note that the starting point is assumed to be valid, so it might not be included in the bank.

Example 1:
Input: start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
Output: 1

Example 2:
Input: start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
Output: 2

Example 3:
Input: start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
Output: 3

Constraints:
start.length == 8
end.length == 8
0 <= bank.length <= 10
bank[i].length == 8
start, end, and bank[i] consist of only the characters ['A', 'C', 'G', 'T'].
*/
package main

import "fmt"

// BFS, time complexity : O(B), B=length of bank
func minMutation(start string, end string, bank []string) int {
	q := []string{}
	seen := make(map[string]bool)
	q = append(q, start)
	seen[start] = true
	cnt := 0
	for len(q) > 0 {
		// q 는 계속 변경되니, 현재 q 범위를 알고 있어야 한다.
		currQSize := len(q)
		for i := 0; i < currQSize; i++ {
			front := q[0]
			q = q[1:]
			for front == end {
				return cnt
			}
			for _, c := range "ACGT" {
				for j := 0; j < len(front); j++ {
					temp := []rune(front)
					temp[j] = rune(c)
					neighbor := string(temp)
					// fmt.Println(neighbor)
					if _, exist := seen[neighbor]; exist {
						continue
					}
					bFind := false
					for b := 0; b < len(bank); b++ {
						if bank[b] == neighbor {
							bFind = true
							break
						}
					}
					if !bFind {
						continue
					}
					q = append(q, neighbor)
					seen[neighbor] = true
					// fmt.Println(q)
				}
			}
		}
		cnt++
	}
	return -1
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
}
