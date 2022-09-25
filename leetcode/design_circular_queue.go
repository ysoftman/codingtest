/*
https://leetcode.com/problems/design-circular-queue/
622. Design Circular Queue
Medium

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Implementation the MyCircularQueue class:

MyCircularQueue(k) Initializes the object with the size of the queue to be k.
int Front() Gets the front item from the queue. If the queue is empty, return -1.
int Rear() Gets the last item from the queue. If the queue is empty, return -1.
boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
boolean isEmpty() Checks whether the circular queue is empty or not.
boolean isFull() Checks whether the circular queue is full or not.
You must solve the problem without using the built-in queue data structure in your programming language.

Example 1:
Input
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Output
[null, true, true, true, false, 3, true, true, true, 4]
Explanation
MyCircularQueue myCircularQueue = new MyCircularQueue(3);
myCircularQueue.enQueue(1); // return True
myCircularQueue.enQueue(2); // return True
myCircularQueue.enQueue(3); // return True
myCircularQueue.enQueue(4); // return False
myCircularQueue.Rear();     // return 3
myCircularQueue.isFull();   // return True
myCircularQueue.deQueue();  // return True
myCircularQueue.enQueue(4); // return True
myCircularQueue.Rear();     // return 4

Constraints:
1 <= k <= 1000
0 <= value <= 1000
At most 3000 calls will be made to enQueue, deQueue, Front, Rear, isEmpty, and isFull.
*/
package main

import "fmt"

type MyCircularQueue struct {
	q     []int
	front int
	rear  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		// front 위치에는 값을 넣지 않아야 하기 때문에 k+1 공간이 필요한다.
		// front 위치를 항상 비워둠으로써 다음 2가지 상태를 파악할 수 있다.
		// front == rear ==> empty
		// front == rear+1 ==> full
		q:     make([]int, k+1),
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.rear == len(this.q)-1 {
		this.rear = 0
	} else {
		this.rear++
	}
	this.q[this.rear] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.front++
	if this.front >= len(this.q) {
		this.front = 0
	}
	this.q[this.front] = -1 // delete front value
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	// front 위치는 항상 비워둔다.
	// front+1 위치가 실제 값이 있는 위치다.
	if this.front+1 >= len(this.q) {
		return this.q[0]
	}
	return this.q[this.front+1]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.front == this.rear {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	// rear 다음칸이 front 이면 꽉찬것임
	if this.rear+1 >= len(this.q) {
		if this.front == 0 {
			return true
		}
	} else {
		if this.rear+1 == this.front {
			return true
		}
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
func main() {
	myCircularQueue := Constructor(3)
	fmt.Println(myCircularQueue.EnQueue(1)) // return True
	fmt.Println(myCircularQueue.EnQueue(2)) // return True
	fmt.Println(myCircularQueue.EnQueue(3)) // return True
	fmt.Println(myCircularQueue.EnQueue(4)) // return False
	fmt.Println(myCircularQueue.Rear())     // return 3
	fmt.Println(myCircularQueue.IsFull())   // return True
	fmt.Println(myCircularQueue.DeQueue())  // return True
	fmt.Println(myCircularQueue.EnQueue(4)) // return True
	fmt.Println(myCircularQueue.Rear())     // return 4

	fmt.Println("-----")
	// ["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
	// [[6],[6],[],[],[],[5],[],[],[],[],[],[]]
	myCircularQueue = Constructor(6)
	fmt.Println(myCircularQueue.EnQueue(6))
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.EnQueue(5))
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.Front())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.DeQueue())
}
