/*
https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
1339. Maximum Product of Splitted Binary Tree
Medium

Given the root of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized.

Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it modulo 109 + 7.

Note that you need to maximize the answer before taking the mod and not after taking it.

Example 1:
Input: root = [1,2,3,4,5,6]
Output: 110
Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)

Example 2:
Input: root = [1,null,2,3,4,null,null,5,6]
Output: 90
Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)

Constraints:
The number of nodes in the tree is in the range [2, 5 * 104].
1 <= Node.val <= 104
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
const modulo = 1000000000 + 7

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getTotalSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + getTotalSum(root.Left) + getTotalSum(root.Right)
}
func sumSubtree(root *TreeNode, totalSum int, r *int) int {
	if root == nil {
		return 0
	}
	leftSum := sumSubtree(root.Left, totalSum, r)
	rightSum := sumSubtree(root.Right, totalSum, r)
	// 전체합에서 현재left/right Subtree 곱중 큰것을 결과로 가져온다.
	*r = max(*r, max((totalSum-leftSum)*leftSum, (totalSum-rightSum)*rightSum))
	// 현재노드 하위 트리의 전체 합을 리턴
	return root.Val + leftSum + rightSum
}
func maxProduct(root *TreeNode) int {
	// 노드 전체 합을 파악해둔다.
	totalSum := getTotalSum(root)

	r := 0
	sumSubtree(root, totalSum, &r)
	return r % modulo
}

func main() {
	root := makeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "5", "6"})
	fmt.Println(maxProduct(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "null", "2", "3", "4", "null", "null", "5", "6"})
	fmt.Println(maxProduct(root))
}
