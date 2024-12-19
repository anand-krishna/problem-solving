const TableSize int = 10

func hashFunc(key int, modValue int) int {
	return key % modValue
}

// Hold key and value to handle operations to this particular key.
type DoublyLinkedNode struct {
	Key        int
	val        int
	prev, next *DoublyLinkedNode
}

type MyHashMap struct {
	tableSize int
	table     []*DoublyLinkedNode
}

func Constructor() MyHashMap {
	return MyHashMap{
		tableSize: TableSize,
		table:     make([]*DoublyLinkedNode, TableSize, TableSize),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	tableIndex := hashFunc(key, this.tableSize)

	nodeToAdd := &DoublyLinkedNode{
		Key:  key,
		val:  value,
		prev: nil,
		next: nil,
	}

	if this.table[tableIndex] == nil {
		this.table[tableIndex] = nodeToAdd
	} else {
		wasElementAdded := false
		currentNode := this.table[tableIndex]
		for ; currentNode != nil; currentNode = currentNode.next {
			if currentNode.Key == key {
				currentNode.val = value
				wasElementAdded = true
			}
		}
		if !wasElementAdded {
			this.table[tableIndex].prev = nodeToAdd
			nodeToAdd.next = this.table[tableIndex]
			this.table[tableIndex] = nodeToAdd
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	tableIndex := hashFunc(key, this.tableSize)

	for currentNode := this.table[tableIndex]; currentNode != nil; currentNode = currentNode.next {
		// fmt.Printf("%v\n", currentNode)
		if currentNode.Key == key {
			return currentNode.val
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	tableIndex := hashFunc(key, this.tableSize)

	for currentNode := this.table[tableIndex]; currentNode != nil; currentNode = currentNode.next {
		if currentNode.Key == key {
			// Case 1. Key at beginnig.
			// Case 2. Last key.
			// Case 3. Key in the middle.
			prevNode, nextNode := currentNode.prev, currentNode.next
			if this.table[tableIndex] == currentNode {
				this.table[tableIndex] = nextNode
			} else {
				if prevNode != nil {
					prevNode.next = nextNode
				}
				if nextNode != nil {
					nextNode.prev = prevNode
				}
				currentNode.prev = nil
				currentNode.next = nil
			}
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */