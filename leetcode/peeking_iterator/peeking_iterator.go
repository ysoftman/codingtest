/*
https://leetcode.com/problems/peeking-iterator/
284. Peeking Iterator
Medium

Design an iterator that supports the peek operation on an existing iterator in addition to the hasNext and the next operations.

Implement the PeekingIterator class:

PeekingIterator(Iterator<int> nums) Initializes the object with the given integer iterator iterator.
int next() Returns the next element in the array and moves the pointer to the next element.
boolean hasNext() Returns true if there are still elements in the array.
int peek() Returns the next element in the array without moving the pointer.
Note: Each language may have a different implementation of the constructor and Iterator, but they all support the int next() and boolean hasNext() functions.

Example 1:
Input
["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 2, 2, 3, false]
Explanation
PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
peekingIterator.hasNext(); // return False

Constraints:
1 <= nums.length <= 1000
1 <= nums[i] <= 1000
All the calls to next and peek are valid.
At most 1000 calls will be made to next, hasNext, and peek.

Follow up: How would you extend your design to be generic and work with all types, not just integer?
*/
package main

import "fmt"

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter          *Iterator
	peekedElement int
	hasPeeked     bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:          iter,
		peekedElement: 0,
		hasPeeked:     false,
	}
}

func (this *PeekingIterator) hasNext() bool {
	// 이미 hasPeeked 라면 이미 next 로 간것으로 hasNext 하지 않고 다음 element 가 있는것임.
	return this.hasPeeked || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if !this.hasPeeked {
		return this.iter.next()
	}
	// 이미 hasPeeked 라면 이미 next 로 간것으로 peek 값을 리턴
	r := this.peekedElement
	this.hasPeeked = false
	this.peekedElement = 0
	return r
}

func (this *PeekingIterator) peek() int {
	if !this.hasPeeked {
		this.peekedElement = this.iter.next()
		this.hasPeeked = true
	}
	return this.peekedElement
}

func main() {
	iter := Iterator{
		list:   []int{1, 2, 3},
		curIdx: -1,
	}
	p := Constructor(&iter)
	fmt.Println(p.next())
	fmt.Println(p.peek())
	fmt.Println(p.next())
	fmt.Println(p.next())
	fmt.Println(p.hasNext())
}

type Iterator struct {
	list   []int
	curIdx int
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	if this.curIdx+1 < len(this.list) {
		return true
	}
	return false
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	if this.hasNext() {
		this.curIdx++
		return this.list[this.curIdx]
	}
	return -1
}
