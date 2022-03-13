/*
https://leetcode.com/problems/validate-binary-search-tree/
98. Validate Binary Search Tree
Medium
Given the root of a binary tree, determine if it is a valid binary search tree (BST).
A valid BST is defined as follows:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.


Constraints:
The number of nodes in the tree is in the range [1, 104].
-231 <= Node.val <= 231 - 1
*/

// go run ./validate_binary_search_tree.go ./ysoftman_common.go
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
BST : false
  5
 / \
4   6
   / \
  3   7
*/
func checkBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min < root.Val && root.Val < max {
		left := checkBST(root.Left, min, root.Val)
		right := checkBST(root.Right, root.Val, max)
		if left && right {
			return true
		}
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	return checkBST(root, -1<<31-1, 1<<31)
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"2", "1", "3"})
	printTreeNodeByBFS(root)
	fmt.Println(isValidBST(root))
	root = makeArrayToBinaryTreeNode([]string{"5", "1", "4", "null", "null", "3", "6"})
	printTreeNodeByBFS(root)
	fmt.Println(isValidBST(root))
	root = makeArrayToBinaryTreeNode([]string{"5", "4", "6", "null", "null", "3", "7"})
	printTreeNodeByBFS(root)
	fmt.Println(isValidBST(root))
	root = makeArrayToBinaryTreeNode([]string{"2147483647"})
	printTreeNodeByBFS(root)
	fmt.Println(isValidBST(root))
	root = makeArrayToBinaryTreeNode([]string{"2", "2", "2"})
	printTreeNodeByBFS(root)
	fmt.Println(isValidBST(root))
}
