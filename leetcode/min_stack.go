/*
https://leetcode.com/problems/min-stack/
155. Min Stack
Easy
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
Implement the MinStack class:
MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.

Example 1:
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
Output
[null,null,null,null,-3,null,0,-2]
Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:
-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
*/
package main

import "fmt"

type MinStack struct {
	s    []int
	mins []int // min 값을 top 으로 유지하는 스택
}

func Constructor() MinStack {
	return MinStack{
		s:    make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else {
		if this.GetMin() >= val {
			this.mins = append(this.mins, val)
		}
	}
	this.s = append(this.s, val)
}

func (this *MinStack) Pop() {
	if len(this.s) > 0 {
		// pop 하는 값이 최소 값이마년 this.min 에서도 빼준다.
		if this.GetMin() == this.s[len(this.s)-1] {
			this.mins = this.mins[:len(this.mins)-1]
		}
		this.s = this.s[:len(this.s)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.s) == 0 {
		return 0
	}
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		return 0
	}
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(3)
	obj.Push(4)
	obj.Push(1)
	// s 에 [2,3,4,1] 이면 mins 에는 [2,1]
	fmt.Println(obj.s)
	fmt.Println(obj.mins)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	fmt.Println("-----")
	// ["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
	// [[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	// s = [2,0,3,0], mins = [2,0,0]
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())

}
