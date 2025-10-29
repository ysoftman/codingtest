/*
https://leetcode.com/problems/binary-search-tree-iterator/
173. Binary Search Tree Iterator
Medium
Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The root of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.
boolean hasNext() Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
int next() Moves the pointer to the right, then returns the number at the pointer.
Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

You may assume that next() calls will always be valid. That is, there will be at least a next number in the in-order traversal when next() is called.

Example 1:
Input
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // return 3
bSTIterator.next();    // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 20
bSTIterator.hasNext(); // return False

Constraints:
The number of nodes in the tree is in the range [1, 105].
0 <= Node.val <= 106
At most 105 calls will be made to hasNext, and next.

Follow up:
Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?
*/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	list   []int
	curIdx int
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{
		list:   make([]int, 0),
		curIdx: 0,
	}
	stack := make([]*TreeNode, 0)

	// make inorder traversal result to bstiterator
	stack = append(stack, root)
	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root.Left)
			root = root.Left
		}

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			bst.list = append(bst.list, top.Val)
		}
		if top != nil && top.Right != nil {
			stack = append(stack, top.Right)
			root = top.Right
		}
	}
	// fmt.Println("bst.list:", bst.list)
	return bst
}

func (this *BSTIterator) Next() int {
	if this.HasNext() {
		this.curIdx++
		return this.list[this.curIdx-1]
	}
	return 0
}

func (this *BSTIterator) HasNext() bool {
	if this.curIdx < len(this.list) {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	root := makeArrayToBinaryTreeNode([]string{"7", "3", "15", "null", "null", "9", "20"})
	obj := Constructor(root)
	fmt.Println(obj.Next())    // return 3
	fmt.Println(obj.Next())    // return 7
	fmt.Println(obj.HasNext()) // return True
	fmt.Println(obj.Next())    // return 9
	fmt.Println(obj.HasNext()) // return True
	fmt.Println(obj.Next())    // return 15
	fmt.Println(obj.HasNext()) // return True
	fmt.Println(obj.Next())    // return 20
	fmt.Println(obj.HasNext()) // return False
}
