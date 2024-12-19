/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// Floyd's Cycle Finding Algorithm aka Tortoise and Hare Algorithm.
	// Inserting all nodes into a map also works. But memory usage will be higher.

	if head == nil || head.Next == nil {
		return false
	}

	for slow, fast := head, head; slow != nil && fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}