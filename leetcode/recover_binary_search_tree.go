/*
https://leetcode.com/problems/recover-binary-search-tree/
99. Recover Binary Search Tree
Medium
You are given the root of a binary search tree (BST), where the values of exactly two nodes of the tree were swapped by mistake. Recover the tree without changing its structure.

Example 1:
Input: root = [1,3,null,null,2]
Output: [3,1,null,null,2]
Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.

Example 2:
Input: root = [3,1,4,null,null,2]
Output: [2,1,4,null,null,3]
Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.

Constraints:
The number of nodes in the tree is in the range [2, 1000].
-231 <= Node.val <= 231 - 1

Follow up: A solution using O(n) space is pretty straight-forward. Could you devise a constant O(1) space solution?
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
/*
 2
3 1

 3
2 1

 1
2 3

 2
1 3
*/
func recursiveRecoverTree(root, minparent, maxparent *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val < minparent.Val {
		fmt.Println("root.Val, minparent.Val:", root.Val, minparent.Val)
		root.Val, minparent.Val = minparent.Val, root.Val
		return true
	} else if root.Val > maxparent.Val {
		fmt.Println("root.Val, maxparent.Val:", root.Val, maxparent.Val)
		root.Val, maxparent.Val = maxparent.Val, root.Val
		return true
	}
	return recursiveRecoverTree(root.Left, minparent, root) || recursiveRecoverTree(root.Right, root, maxparent)
}
func recoverTree(root *TreeNode) {
	// if true(changed) then check again!
	for recursiveRecoverTree(root, &TreeNode{-1<<31 - 1, nil, nil}, &TreeNode{1<<31 - 1, nil, nil}) {
	}
	fmt.Println("-----")
}
func main() {
	node := makeArrayToBinaryTreeNode([]string{"3", "null", "2", "null", "1"})
	recoverTree(node)
	printTreeNodeByBFS(node)
	node = makeArrayToBinaryTreeNode([]string{"2", "3", "1"})
	recoverTree(node)
	printTreeNodeByBFS(node)
	node = makeArrayToBinaryTreeNode([]string{"2", "null", "1"})
	recoverTree(node)
	printTreeNodeByBFS(node)
	node = makeArrayToBinaryTreeNode([]string{"1", "3", "null", "null", "2"})
	recoverTree(node)
	printTreeNodeByBFS(node)
	node = makeArrayToBinaryTreeNode([]string{"3", "1", "4", "null", "null", "2"})
	recoverTree(node)
	printTreeNodeByBFS(node)
}
