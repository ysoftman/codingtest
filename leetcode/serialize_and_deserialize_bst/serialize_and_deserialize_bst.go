/*
https://leetcode.com/problems/serialize-and-deserialize-bst/
449. Serialize and Deserialize BST
Medium

Serialization is converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary search tree. There is no restriction on how your serialization/deserialization algorithm should work. You need to ensure that a binary search tree can be serialized to a string, and this string can be deserialized to the original tree structure.

The encoded string should be as compact as possible.

Example 1:
Input: root = [2,1,3]
Output: [2,1,3]

Example 2:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 104].
0 <= Node.val <= 104
The input tree is guaranteed to be a binary search tree.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ysoftman/ysoftmancommon"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func _to_string(root *ysoftmancommon.TreeNode, r *[]int) {
	if root == nil {
		return
	}
	*r = append(*r, root.Val)
	_to_string(root.Left, r)
	_to_string(root.Right, r)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *ysoftmancommon.TreeNode) string {
	nodeValues := []int{}
	// preoder 노드 탐색
	_to_string(root, &nodeValues)

	// 노드 값을 공백으로 구분한 스트링으로 리턴
	r := ""
	for i := 0; i < len(nodeValues); i++ {
		r += fmt.Sprintf("%d ", nodeValues[i])
	}
	return strings.TrimSpace(r)
}

func _to_BinaryTreeNode(q *[]string, min, max int) *ysoftmancommon.TreeNode {
	if len(*q) == 0 {
		return nil
	}
	// 큐 맨 앞값 파악
	frontIntVal, _ := strconv.Atoi((*q)[0])
	// 현재 노드 하위가 아닌 경우
	if frontIntVal < min || frontIntVal > max {
		return nil
	}

	// 큐 프론트 빼기
	*q = (*q)[1:]
	// 새 노드
	newNode := &ysoftmancommon.TreeNode{
		Val: frontIntVal,
	}
	newNode.Left = _to_BinaryTreeNode(q, min, frontIntVal)
	newNode.Right = _to_BinaryTreeNode(q, frontIntVal, max)
	return newNode
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *ysoftmancommon.TreeNode {
	if len(data) == 0 {
		return nil
	}
	q := strings.Split(data, " ")
	if len(q) == 0 {
		return nil
	}

	// 큐 맨 앞값 파악
	min := -1 << 31
	max := 1<<31 - 1
	return _to_BinaryTreeNode(&q, min, max)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

func main() {
	root := &ysoftmancommon.TreeNode{
		Val: 2,
		Left: &ysoftmancommon.TreeNode{
			Val: 1,
		},
		Right: &ysoftmancommon.TreeNode{
			Val: 3,
		},
	}
	ysoftmancommon.PrintTreeNodeByDFS(root)
	fmt.Println()

	ser := Constructor()
	treeString := ser.serialize(root)
	fmt.Println(treeString)
	ans := ser.deserialize(treeString)
	// ysoftmancommon.PrintTreeNodeByBFS(ans)
	ysoftmancommon.PrintTreeNodeByDFS(ans)
	fmt.Println()
}
