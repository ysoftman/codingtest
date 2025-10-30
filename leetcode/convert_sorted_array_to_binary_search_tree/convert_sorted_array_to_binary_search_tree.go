/*
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
108. Convert Sorted Array to Binary Search Tree
Easy
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

Example 1:
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:
Input: nums = [1,3]
Output: [3,1]
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.


Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in a strictly increasing order.
*/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addNumToBST(nums []int, left, right int) *TreeNode {
	if left >= right {
		return nil
	}
	mid := (right-left)/2 + left
	root := &TreeNode{Val: nums[mid]}
	root.Left = addNumToBST(nums, left, mid)
	root.Right = addNumToBST(nums, mid+1, right)
	return root
}
func sortedArrayToBST(nums []int) *TreeNode {
	left := 0
	right := len(nums)
	return addNumToBST(nums, left, right)
}

func main() {
	ysoftmancommon.PrintTreeNodeByBFS(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	ysoftmancommon.PrintTreeNodeByBFS(sortedArrayToBST([]int{1, 3}))
	ysoftmancommon.PrintTreeNodeByBFS(sortedArrayToBST([]int{-10, -5, -4, -3, 0, 5, 9}))
}
