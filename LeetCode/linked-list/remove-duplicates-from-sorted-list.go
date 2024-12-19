/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev, current := head, head

	for current != nil {
		for prev != nil && current != nil && prev.Val == current.Val {
			toDelete := current
			current = current.Next
			toDelete.Next = nil
		}
		prev.Next = current
		prev = current
	}

	return head
}