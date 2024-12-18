/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultList, currentResultNode *ListNode = nil, nil
	currentNodeL1, currentNodeL2 := l1, l2
	// Iterate and sum values in each node.
	for ; currentNodeL1 != nil && currentNodeL2 != nil; currentNodeL1, currentNodeL2 = currentNodeL1.Next, currentNodeL2.Next {
		nodeToBeAdded := &ListNode{
			Val:  currentNodeL1.Val + currentNodeL2.Val,
			Next: nil,
		}
		if resultList == nil {
			resultList = nodeToBeAdded
			currentResultNode = nodeToBeAdded
		} else {
			currentResultNode.Next = nodeToBeAdded
			currentResultNode = currentResultNode.Next
		}
	}

	// In case len(l1 > l2) and we could not cover the end of the list.
	for ; currentNodeL1 != nil; currentNodeL1 = currentNodeL1.Next {
		nodeToBeAdded := &ListNode{
			Val:  currentNodeL1.Val,
			Next: nil,
		}
		currentResultNode.Next = nodeToBeAdded
		currentResultNode = currentResultNode.Next
	}

	// In case len(l2 > l1) and we could not cover the end of the list.
	for ; currentNodeL2 != nil; currentNodeL2 = currentNodeL2.Next {
		nodeToBeAdded := &ListNode{
			Val:  currentNodeL2.Val,
			Next: nil,
		}
		currentResultNode.Next = nodeToBeAdded
		currentResultNode = currentResultNode.Next
	}

	// Handle carry by forwarding it to the next node until you reach the end.
	valueToCarry := 0
	for currentNode := resultList; currentNode != nil; currentNode = currentNode.Next {
		if valueToCarry > 0 {
			currentNode.Val += valueToCarry
			valueToCarry = 0
		}
		if currentNode.Val > 9 {
			valueToCarry = currentNode.Val / 10
			currentNode.Val = currentNode.Val % 10
		}
	}

	// Pending carry will be placed in a new node at the end.
	if valueToCarry > 0 {
		currentResultNode.Next = &ListNode{
			Val:  valueToCarry,
			Next: nil,
		}
	}

	return resultList
}