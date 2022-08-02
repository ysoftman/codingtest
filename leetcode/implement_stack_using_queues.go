/*
https://leetcode.com/problems/implement-stack-using-queues
225. Implement Stack using Queues
Easy
Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means that only push to back, peek/pop from front, size and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.


Example 1:

Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // return 2
myStack.pop(); // return 2
myStack.empty(); // return False
*/

package main

import "fmt"

/*
type MyStack struct {
    data []int
}
func Constructor() MyStack {
    return MyStack {
        data : make([]int, 0),
    }
}
func (this *MyStack) Push(x int)  {
    this.data = append(this.data, x)
}
func (this *MyStack) Pop() int {
    ret := this.data[len(this.data)-1]

    // this.data = append(this.data[:len(this.data)-1])
    this.data = this.data[:len(this.data)-1]

    return ret
}
func (this *MyStack) Top() int {
    if len(this.data) <= 0 {
        return 0
    }
    return this.data[len(this.data)-1]
}
func (this *MyStack) Empty() bool {
    if len(this.data) <= 0 {
        return true
    }
    return false
}
*/

/*
push 1
queue1:
queue2:1
queue1:1
queue2:

push 2
queue1:
queue2:1
queue1:2,1
queue2:

push 3
queue1:
queue2:2,1
queue1:3,2,1
queue2:
*/
type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}
func (this *MyStack) Push(x int) {
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.Pop())
	}
	this.queue1 = append(this.queue1, x)
	for len(this.queue2) > 0 {
		this.queue1 = append(this.queue1, this.PopQueue2())
	}
}
func (this *MyStack) PopQueue2() int {
	if len(this.queue2) == 0 {
		return 0
	}
	ret := this.queue2[0]
	this.queue2 = this.queue2[1:]
	return ret
}
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	ret := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return ret
}
func (this *MyStack) Top() int {
	if len(this.queue1) == 0 {
		return 0
	}
	return this.queue1[0]
}
func (this *MyStack) Empty() bool {
	if len(this.queue1) <= 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	fmt.Printf("%v obj.Push(1)\n", obj)
	obj.Push(2)
	fmt.Printf("%v obj.Push(2)\n", obj)
	obj.Push(3)
	fmt.Printf("%v obj.Push(3)\n", obj)
	obj.Push(4)
	fmt.Printf("%v obj.Push(4)\n", obj)
	obj.Push(5)
	fmt.Printf("%v obj.Push(5)\n", obj)
	fmt.Printf("%v Empty(): %v\n", obj, obj.Empty())
	fmt.Printf("%v Top(): %v\n", obj, obj.Top())
	fmt.Printf("%v Top(): %v\n", obj, obj.Top())
	fmt.Printf("%v Pop(): %v\n", obj, obj.Pop())
	fmt.Printf("%v Pop(): %v\n", obj, obj.Pop())
	fmt.Printf("%v Pop(): %v\n", obj, obj.Pop())
	fmt.Printf("%v Pop(): %v\n", obj, obj.Pop())
	fmt.Printf("%v Pop(): %v\n", obj, obj.Pop())
	fmt.Printf("%v Empty(): %v\n", obj, obj.Empty())
}
