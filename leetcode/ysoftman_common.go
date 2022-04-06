/*
local 에서 테스트를 위해 공통으로 필요한 함수들을 모아 둠
*/
package main

import (
	"fmt"
	"strconv"
)

func printMatrix(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%2v", matrix[i][j])
			if j != n-1 {
				fmt.Printf(",")
			}
		}
		fmt.Println()
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(head *ListNode) {
	fmt.Printf("[")
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Printf(",")
		}
		head = head.Next
	}
	fmt.Printf("]")
	fmt.Println()
}

func makeLinkedList(nums []int) (root *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	root = head
	for i, v := range nums {
		head.Val = v
		if i < len(nums)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return root
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeArrayToBinaryTreeNode(arr []string) *TreeNode {
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
func printTreeNodeByBFS(root *TreeNode) {
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
func printTreeNodeByDFS(root *TreeNode) {
	if root == nil {
		fmt.Printf("nil ")
		return
	}
	fmt.Printf("%v ", root.Val)
	printTreeNodeByDFS(root.Left)
	printTreeNodeByDFS(root.Right)
}
