type TrieNode struct {
	isWord   bool
	children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			isWord:   false,
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	currentNode := this.root
	for _, character := range word {
		if _, ok := currentNode.children[character]; !ok {
			currentNode.children[character] = &TrieNode{
				isWord:   false,
				children: make(map[rune]*TrieNode),
			}
		}
		currentNode = currentNode.children[character]
	}
	currentNode.isWord = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this.root
	for _, character := range word {
		if _, ok := currentNode.children[character]; !ok {
			return false
		}
		currentNode = currentNode.children[character]
	}
	return currentNode.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this.root
	for _, character := range prefix {
		if _, ok := currentNode.children[character]; !ok {
			return false
		}
		currentNode = currentNode.children[character]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
