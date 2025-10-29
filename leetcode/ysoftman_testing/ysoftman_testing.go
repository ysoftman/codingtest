package main

import (
	"fmt"

	ysoftmancommon "github.com/ysoftman/ysoftmancommon"
)

// go run ./ysoftman_testing.go
func main() {
	fmt.Println("--- printLinkedList ---")
	func() {
		node := ysoftmancommon.MakeLinkedList([]int{4, 5, 1, 9})
		ysoftmancommon.PrintLinkedList(node)
	}()
	fmt.Println("")

	fmt.Println("--- printTreeNodeByDFS ---")
	func() {
		root := []string{"3", "9", "20", "null", "null", "15", "7"}
		ysoftmancommon.PrintTreeNodeByDFS(ysoftmancommon.MakeArrayToBinaryTreeNode(root))
	}()
	fmt.Println("")
	fmt.Println("")

	fmt.Println("--- printTreeNodeByBFS ---")
	func() {
		root := []string{"3", "9", "20", "null", "null", "15", "7"}
		ysoftmancommon.PrintTreeNodeByBFS(ysoftmancommon.MakeArrayToBinaryTreeNode(root))
	}()
	fmt.Println("")

	fmt.Println("--- printNodeByBFS ---")
	func() {
		root := []string{"1", "null", "3", "2", "4", "null", "5", "6"}
		ysoftmancommon.PrintNodeByBFS(ysoftmancommon.MakeArrayToNode(root))

		root = []string{"1", "null", "2", "3", "4", "5", "null", "null", "6", "7", "null", "8", "null", "9", "10", "null", "null", "11", "null", "12", "null", "13", "null", "null", "14"}
		ysoftmancommon.PrintNodeByBFS(ysoftmancommon.MakeArrayToNode(root))
	}()
	fmt.Println("")
}
