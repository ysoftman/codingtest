/*
https://leetcode.com/problems/binary-tree-tilt/
563. Binary Tree Tilt
Easy
Given the root of a binary tree, return the sum of every tree node's tilt.

The tilt of a tree node is the absolute difference between the sum of all left subtree node values and all right subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0. The rule is similar if the node does not have a right child.

Example 1:
Input: root = [1,2,3]
Output: 1
Explanation:
Tilt of node 2 : |0-0| = 0 (no children)
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
Sum of every tilt : 0 + 0 + 1 = 1

Example 2:
Input: root = [4,2,9,3,5,null,7]
Output: 15
Explanation:
Tilt of node 3 : |0-0| = 0 (no children)
Tilt of node 5 : |0-0| = 0 (no children)
Tilt of node 7 : |0-0| = 0 (no children)
Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15

Example 3:
Input: root = [21,7,14,1,1,2,2,3,3]
Output: 9

Constraints:
The number of nodes in the tree is in the range [0, 104].
-1000 <= Node.val <= 1000
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

var totalSum int = 0

func diffLR(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 자식노드들을 모두 탐색해서 ()leftsum - rightsum) 해야 하기 때문에
	// postorder dfs 로 탐색
	leftSum := diffLR(root.Left)
	rightSum := diffLR(root.Right)
	diff := leftSum - rightSum
	if diff < 0 {
		diff *= -1
	}
	totalSum += diff
	// 현재 노드에서의 합이 부모노드에서는 left sum 또는 right sum 이 된다.
	return root.Val + leftSum + rightSum
}

func findTilt(root *TreeNode) int {
	totalSum = 0
	diffLR(root)
	return totalSum
}

func main() {
	fmt.Println(findTilt(makeArrayToBinaryTreeNode([]string{"1", "2", "3"})))
	fmt.Println(findTilt(makeArrayToBinaryTreeNode([]string{"4", "2", "9", "3", "5", "null", "7"})))
	fmt.Println(findTilt(makeArrayToBinaryTreeNode([]string{"21", "7", "14", "1", "1", "2", "2", "3", "3"})))
}
