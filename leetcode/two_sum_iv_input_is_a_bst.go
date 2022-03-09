/*
https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
653. Two Sum IV - Input is a BST
Easy
Given the root of a Binary Search Tree and a target number k, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:
Input: root = [5,3,6,2,4,null,7], k = 9
Output: true

Example 2:
Input: root = [5,3,6,2,4,null,7], k = 28
Output: false
*/
package main

import (
	"fmt"
	"strconv"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsTwoSum(node *TreeNode, hashmap map[int]int, k int) bool {
	if node == nil {
		return false
	}
	if _, exist := hashmap[k-node.Val]; exist {
		return true
	} else {
		hashmap[node.Val] = 1
	}

	if dfsTwoSum(node.Left, hashmap, k) {
		return true
	}
	if dfsTwoSum(node.Right, hashmap, k) {
		return true
	}
	return false
}
func findTarget(root *TreeNode, k int) bool {
	hashmap := make(map[int]int)
	return dfsTwoSum(root, hashmap, k)
}

func makeArrayToBinaryTree(arr []string) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	temp, _ := strconv.Atoi(arr[0])
	root := &TreeNode{
		Val: temp,
	}
	node := root

	queue := make([]*TreeNode, 0)
	queue = append(queue, node)
	for i := 1; i < len(arr); i++ {
		if len(queue) == 0 {
			continue
		}
		front := queue[0]
		queue = queue[1:]

		if arr[i] != "null" {
			left, _ := strconv.Atoi(arr[i])
			front.Left = &TreeNode{
				Val: left,
			}
			queue = append(queue, front.Left)
		}
		if i+1 >= len(arr) {
			continue
		}
		i++
		if arr[i] != "null" {
			right, _ := strconv.Atoi(arr[i])
			front.Right = &TreeNode{
				Val: right,
			}
			queue = append(queue, front.Right)
		}
	}
	return root
}
func printTreeByBFS(root *TreeNode) {
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if top == nil {
			fmt.Printf("nil ")
			continue
		}
		fmt.Printf("%v ", top.Val)
		q = append(q, top.Left)
		q = append(q, top.Right)
	}
	fmt.Println()
}
func main() {
	node := makeArrayToBinaryTree([]string{"5", "3", "6", "2", "4", "null", "7"})
	printTreeByBFS(node)
	fmt.Println(findTarget(node, 9))
	node = makeArrayToBinaryTree([]string{"5", "3", "6", "2", "4", "null", "7"})
	printTreeByBFS(node)
	fmt.Println(findTarget(node, 28))
}
