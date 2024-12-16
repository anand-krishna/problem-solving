type TrieNode struct {
	isWord   bool
	children map[rune]*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			isWord:   false,
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	return searchForWord([]rune(word), 0, this.root)
}

func searchForWord(word []rune, index int, currentNode *TrieNode) bool {
	for i := index; i < len(word); i++ {
		if word[i] == '.' {
			// Wildcard. Iterate over all children and return the first one that matches.
			for _, childNode := range currentNode.children {
				if searchForWord(word, i+1, childNode) {
					return true
				}
			}
			return false // Could not find anything.
		} else {
			if _, ok := currentNode.children[word[i]]; !ok {
				return false
			}
			currentNode = currentNode.children[word[i]]
		}
	}
	return currentNode.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */