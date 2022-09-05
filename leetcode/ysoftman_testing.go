package main

import "fmt"

// go run ./ysoftman_test.go ./ysoftman_common.go
func main() {
	fmt.Println("--- printLinkedList ---")
	func() {
		node := makeLinkedList([]int{4, 5, 1, 9})
		printLinkedList(node)
	}()
	fmt.Println("")

	fmt.Println("--- printTreeNodeByDFS ---")
	func() {
		root := []string{"3", "9", "20", "null", "null", "15", "7"}
		printTreeNodeByDFS(makeArrayToBinaryTreeNode(root))
	}()
	fmt.Println("\n")

	fmt.Println("--- printTreeNodeByBFS ---")
	func() {
		root := []string{"3", "9", "20", "null", "null", "15", "7"}
		printTreeNodeByBFS(makeArrayToBinaryTreeNode(root))
	}()
	fmt.Println("")

	fmt.Println("--- printNodeByBFS ---")
	func() {
		root := []string{"1", "null", "3", "2", "4", "null", "5", "6"}
		printNodeByBFS(makeArrayToNode(root))

		root = []string{"1", "null", "2", "3", "4", "5", "null", "null", "6", "7", "null", "8", "null", "9", "10", "null", "null", "11", "null", "12", "null", "13", "null", "null", "14"}
		printNodeByBFS(makeArrayToNode(root))
	}()
	fmt.Println("")
}
