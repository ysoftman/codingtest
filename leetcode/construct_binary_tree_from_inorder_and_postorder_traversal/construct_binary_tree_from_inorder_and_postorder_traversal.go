/*
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
106. Construct Binary Tree from Inorder and Postorder Traversal
Medium
Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

Example 1:
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]

Example 2:
Input: inorder = [-1], postorder = [-1]
Output: [-1]

Constraints:
1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder and postorder consist of unique values.
Each value of postorder also appears in inorder.
inorder is guaranteed to be the inorder traversal of the tree.
postorder is guaranteed to be the postorder traversal of the tree.
*/

// go run ./construct_binary_tree_from_inorder_and_postorder_traversal.go
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

/*
inorder = [9,3,15,20,7] => inorder 의 첫번째 원소는 left
postorder = [9,15,7,20,3] => postorder 의 마지막 원소는 root

	     3
	 /      \
	9    [15,20,7]   inorder 에서 root, left 제외한 나머지, 15 가 left
	     [15,7,20]   postorder 에서 root, left 제외한 나머지, 20 이 root
	          20
	      /       \
	    15         [7]
	               [7]
*/
func makeTree(im map[int]int, inorder, postorder []int, inorderLeftIdx, inorderRightIdx int, postorderRightIdx *int) *ysoftmancommon.TreeNode {
	if inorderLeftIdx > inorderRightIdx {
		return nil
	}
	// postorder 마지막 원소로 root 생성
	root := &ysoftmancommon.TreeNode{Val: postorder[*postorderRightIdx]}
	// inorder 에서 root 값은 제외 하기 위해 index를 파악한다.
	inorderIdx := im[root.Val]
	// root 값을 제외한 inorder, postorder 에서 left, right 값을 찾는다.
	// postorder 마지막 원소로 root 가져왔으니 이제 마지막 원소는 제외
	// postorderRightIdx 값은 모든 콜스택에서 공유되어야 한다.
	*postorderRightIdx--

	// right 노드부터 파악해야 된다.
	root.Right = makeTree(im, inorder, postorder, inorderIdx+1, inorderRightIdx, postorderRightIdx)
	root.Left = makeTree(im, inorder, postorder, inorderLeftIdx, inorderIdx-1, postorderRightIdx)

	return root
}

func buildTree(inorder []int, postorder []int) *ysoftmancommon.TreeNode {
	im := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		im[inorder[i]] = i
	}
	postorderRightIdx := len(postorder) - 1
	return makeTree(im, inorder, postorder, 0, len(inorder)-1, &postorderRightIdx)
}

func main() {
	ysoftmancommon.PrintTreeNodeByBFS(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	ysoftmancommon.PrintTreeNodeByBFS(buildTree([]int{-1}, []int{-1}))
}
