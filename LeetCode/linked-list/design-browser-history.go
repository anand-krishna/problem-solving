type DoublyLinkedNode struct {
	url        string
	next, prev *DoublyLinkedNode
}

// Move reference as required. Can also be implemented with two
// stacks/linked-lists. Maybe faster. Need to check.
type BrowserHistory struct {
	homepage  string
	reference *DoublyLinkedNode
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		homepage: homepage,
		reference: &DoublyLinkedNode{
			url:  homepage,
			next: nil,
			prev: nil,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	// This is just adding a new node at the end.
	// Use the browser to see for yourself...

	// These few lines are for clearing the forward visit
	// history. Increase the runtime.
	oldNextOfReference := this.reference.next
	if oldNextOfReference != nil {
		oldNextOfReference.prev = nil
	}

	prevNode, currentNode := oldNextOfReference, oldNextOfReference
	for currentNode != nil {
		prevNode = currentNode
		currentNode = currentNode.next

		prevNode.next = nil

		if currentNode != nil {
			currentNode.prev = nil
		}
	}

	this.reference.next = &DoublyLinkedNode{
		url:  url,
		next: nil,
		prev: this.reference,
	}
	this.reference = this.reference.next
}

func (this *BrowserHistory) Back(steps int) string {
	currentNode := this.reference
	for ; steps > 0 && currentNode.prev != nil; currentNode = currentNode.prev {
		steps = steps - 1
	}
	this.reference = currentNode
	return this.reference.url
}

func (this *BrowserHistory) Forward(steps int) string {
	currentNode := this.reference
	for ; steps > 0 && currentNode.next != nil; currentNode = currentNode.next {
		steps = steps - 1
	}
	this.reference = currentNode
	return this.reference.url
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */