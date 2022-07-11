/*
https://leetcode.com/problems/binary-tree-right-side-view/
199. Binary Tree Right Side View
Medium
Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
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
오른쪽에서 binary tree 를 봤을때 보이는 노드 찾기
현재 level 에서 오른쪽 노드는 nil 인데
next(child) level 이 아직 남아 있을 수 도 있어
무조건 오른쪽 노드만 탐색하면 안되고 각 level에 가장 오른쪽에 있는 노드만 추가한다.
힌트) 각 level에 노드값 1개씩만 보인다(추가)

                오른쪽에서 봤을때
     1     <--- 1이 보인다.
   2   3   <--- 3이 보인다.
 4    <--- 4가 보인다.
5  6   <--- 6이 보인다.
*/

// dfs(pre-order) 방법
func recursiveRightNode(root *TreeNode, level int, r *[]int) {
	if root == nil {
		return
	}
	// 각 레벨당 1개씩이니, 현재 레벨에 보이는것 추가
	if len(*r) < level {
		*r = append(*r, root.Val)
	}
	// 오른쪽 부터 탐색
	recursiveRightNode(root.Right, level+1, r)
	recursiveRightNode(root.Left, level+1, r)

}
func rightSideView(root *TreeNode) []int {
	result := []int{}
	recursiveRightNode(root, 1, &result)
	return result
}

// bfs 방법
func rightSideView2(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		// 현재 큐의 가장 오른쪽
		right := q[len(q)-1]
		result = append(result, right.Val)

		tempq := []*TreeNode{}
		for len(q) > 0 {
			top := q[0]
			q = q[1:]
			if top.Left != nil {
				tempq = append(tempq, top.Left)
			}
			if top.Right != nil {
				tempq = append(tempq, top.Right)
			}
		}
		q = append(q, tempq...)
		// fmt.Println("---", q)
	}
	return result
}
func main() {
	root := makeArrayToBinaryTreeNode([]string{"1", "2", "3", "null", "5", "null", "4"})
	fmt.Println(rightSideView(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "null", "3"})
	fmt.Println(rightSideView(root))
	root = makeArrayToBinaryTreeNode([]string{})
	fmt.Println(rightSideView(root))
	root = makeArrayToBinaryTreeNode([]string{"1", "2", "3", "4", "null", "null", "null", "5", "6"})
	fmt.Println(rightSideView(root))
}
