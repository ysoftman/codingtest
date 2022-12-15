/*
https://leetcode.com/problems/flatten-nested-list-iterator/
341. Flatten Nested List Iterator
Medium

You are given a nested list of integers nestedList. Each element is either an integer or a list whose elements may also be integers or other lists. Implement an iterator to flatten it.

Implement the NestedIterator class:

NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with the nested list nestedList.
int next() Returns the next integer in the nested list.
boolean hasNext() Returns true if there are still some integers in the nested list and false otherwise.
Your code will be tested with the following pseudocode:

initialize iterator with nestedList
res = []
while iterator.hasNext()

	append iterator.next() to the end of res

return res
If res matches the expected flattened list, then your code will be judged as correct.

Example 1:
Input: nestedList = [[1,1],2,[1,1]]
Output: [1,1,2,1,1]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,1,2,1,1].

Example 2:
Input: nestedList = [1,[4,[6]]]
Output: [1,4,6]
Explanation: By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,4,6].

Constraints:
1 <= nestedList.length <= 500
The values of the integers in the nested list is in the range [-106, 106].
*/
package main

import "fmt"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
//
type NestedInteger struct {
	NestedList []*NestedInteger
	Value      int
}

// [[1,1],2,[1,1]]
func MakeNestedInteger() NestedInteger {
	ni := NestedInteger{}
	ele1 := NestedInteger{
		Value: 1,
	}
	ele2 := NestedInteger{
		Value: 1,
	}
	group1 := NestedInteger{}
	group1.NestedList = append(group1.NestedList, &ele1)
	group1.NestedList = append(group1.NestedList, &ele2)
	ni.NestedList = append(ni.NestedList, &group1)
	ni.NestedList = append(ni.NestedList, &NestedInteger{Value: 2})
	ele3 := NestedInteger{
		Value: 1,
	}
	ele4 := NestedInteger{
		Value: 1,
	}
	group2 := NestedInteger{}
	group2.NestedList = append(group2.NestedList, &ele3)
	group2.NestedList = append(group2.NestedList, &ele4)
	ni.NestedList = append(ni.NestedList, &group2)

	root := NestedInteger{}
	root.NestedList = append(root.NestedList, &ni)
	return root
}
func (this NestedInteger) IsInteger() bool {
	if this.NestedList == nil {
		return true
	}
	return false
}
func (this NestedInteger) GetInteger() int {
	return this.Value
}
func (this NestedInteger) GetList() []*NestedInteger {
	return this.NestedList
}

func main() {
	root := MakeNestedInteger()
	iter := Constructor(root.NestedList)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}

/////////////////////////////////////

type NestedIterator struct {
	list   []int
	curidx int
}

func dfs(nestedList []*NestedInteger, r *[]int) {
	if len(nestedList) == 0 {
		return
	}
	for _, v := range nestedList {
		if v.IsInteger() {
			*r = append(*r, v.GetInteger())
		} else {
			dfs(v.GetList(), r)
		}
	}
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := NestedIterator{}
	dfs(nestedList, &ni.list)
	// fmt.Println("list:", ni.list)
	return &ni
}

func (this *NestedIterator) Next() int {
	n := 0
	if this.curidx < len(this.list) {
		n = this.list[this.curidx]
		this.curidx++
	}
	return n
}

func (this *NestedIterator) HasNext() bool {
	if this.curidx < len(this.list) {
		return true
	}
	return false
}
