/*
https://leetcode.com/problems/implement-queue-using-stacks/
232. Implement Queue using Stacks
Easy
Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).
Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.
Notes:

You must use only standard operations of a stack, which means only push to top, peek/pop from top, size, and is empty operations are valid.
Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

Example 1:
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
*/
package main

import "fmt"

/*
push 1
stack1:
stack2:1
stack1:1
stack2:

push 2
stack1:
stack2:1
stack1:2,1
stack2:

push 3
stack1:2,1
stack2:1,2
stack1:3,2,1
stack2:
*/
type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: []int{},
		stack2: []int{},
	}
}
func (this *MyQueue) Push(x int) {
	// stack1 -> stack2
	for len(this.stack1) > 0 {
		this.stack2 = append(this.stack2, this.Pop())
	}
	this.stack1 = append(this.stack1, x)
	// stack2 -> stack1
	for len(this.stack2) > 0 {
		this.stack1 = append(this.stack1, this.PopStack2())
	}
}
func (this *MyQueue) PopStack2() int {
	if len(this.stack2) == 0 {
		return 0
	}
	pop := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return pop
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	pop := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	return pop
}
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.stack1[len(this.stack1)-1]
}
func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func main() {
	myQueue := MyQueue{}
	myQueue.Push(1) // queue is: [1]
	fmt.Printf("Push(1), myQueue=%v\n", myQueue)
	myQueue.Push(2) // queue is: [1, 2] (leftmost is front of the queue)
	fmt.Printf("Push(2), myQueue=%v\n", myQueue)
	fmt.Printf("Peek(): %v, myQueue=%v\n", myQueue.Peek(), myQueue) // return 1
	myQueue.Pop()                                                   // return 1, queue is [2]
	fmt.Printf("Pop(), myQueue=%v\n", myQueue)
	fmt.Printf("Empty(): %v, myQueue=%v\n", myQueue.Empty(), myQueue) // return false
}
