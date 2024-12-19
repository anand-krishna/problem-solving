/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// Inserting all nodes into a map and count number of times they are visited.
	// First node with numVisits > 1 => Cycle start.

	if head == nil || head.Next == nil {
		return nil
	}

	nodeVisits := make(map[*ListNode]int)

	for currentNode := head; currentNode != nil; currentNode = currentNode.Next {
		if _, ok := nodeVisits[currentNode]; !ok {
			// First time visiting.
			nodeVisits[currentNode] = 1
		} else {
			nodeVisits[currentNode]++
			if nodeVisits[currentNode] > 1 {
				return currentNode
			}
		}
	}
	return nil
}