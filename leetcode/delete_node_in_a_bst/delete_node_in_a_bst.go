/*
https://leetcode.com/problems/delete-node-in-a-bst/
450. Delete Node in a BST
Medium

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.

Example 1:
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.

Example 3:
Input: root = [], key = 0
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 104].
-105 <= Node.val <= 105
Each node has a unique value.
root is a valid binary search tree.
-105 <= key <= 105

Follow up: Could you solve it with time complexity O(height of tree)?
*/
package main

import "github.com/ysoftman/ysoftmancommon"

/**
* Definition for a binary tree node.
* type TreeNode struct {
	*     Val int
	*     Left *TreeNode
	*     Right *TreeNode
	* }
*/

func findMinNode(root *ysoftmancommon.TreeNode) *ysoftmancommon.TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func deleteNode(root *ysoftmancommon.TreeNode, key int) *ysoftmancommon.TreeNode {
	if root == nil {
		return nil
	}
	// left 부분에서 찾아야 하는 경우 필요한 경우
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val { // right 부분에서 찾아야 하는 경우
		root.Right = deleteNode(root.Right, key)
	} else { // key == root.Val 로 찾은 경우
		// left, right 중 하나만 존재하면  존재하는 노드 선택
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// left, right 둘다 있다면 하위 가장 작은값 노드로 대체한다.
		minNode := findMinNode(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}

func main() {
	root := ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "7"})
	key := 3
	deleteNode(root, key)
	ysoftmancommon.PrintTreeNodeByBFS(root)
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{"5", "3", "6", "2", "4", "null", "7"})
	key = 0
	deleteNode(root, key)
	ysoftmancommon.PrintTreeNodeByBFS(root)
	root = ysoftmancommon.MakeArrayToBinaryTreeNode([]string{})
	key = 0
	deleteNode(root, key)
	ysoftmancommon.PrintTreeNodeByBFS(root)
}
