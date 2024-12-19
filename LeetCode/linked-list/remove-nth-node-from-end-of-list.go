/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLength := 0
	for currentNode := head; currentNode != nil; currentNode = currentNode.Next {
		listLength++
	}

	// Delete only node.
	if listLength == 1 {
		head = nil
		return head
	}

	// Delete head node.
	if listLength-n+1 == 1 {
		head = head.Next
	} else {
		// Traverse and delete.
		prevNode, currentNode := head, head
		for ; currentNode != nil && (listLength-n) > 0; listLength = listLength - 1 {
			prevNode = currentNode
			currentNode = currentNode.Next
		}
		if currentNode != nil {
			prevNode.Next = currentNode.Next
		} else {
			prevNode.Next = nil
		}
	}

	return head
}