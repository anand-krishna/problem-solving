/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// Find length.
	length := 0
	for currentNode := head; currentNode != nil; currentNode = currentNode.Next {
		length++
	}

	currentNode := head
	// Iterate and return.
	for i := 0; i < length/2; i++ {
		currentNode = currentNode.Next
	}

	return currentNode
}