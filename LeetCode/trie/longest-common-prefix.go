type TrieNode struct {
	frequency int
	children  map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			frequency: 0,
			children:  make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	currentNode := this.root
	for _, character := range word {
		if _, ok := currentNode.children[character]; !ok {
			currentNode.children[character] = &TrieNode{
				frequency: 1,
				children:  make(map[rune]*TrieNode),
			}
		} else {
			currentNode.children[character].frequency++
		}
		currentNode = currentNode.children[character]
	}
}

func (this *Trie) GetLCP(word string, stringCount int) string {
	returnIndex := 0
	currentNode := this.root
	for _, character := range word {
		if currentNode.children[character].frequency < stringCount {
			break
		}
		returnIndex++
		currentNode = currentNode.children[character]
	}
	return word[0:returnIndex]
}

func longestCommonPrefix(strs []string) string {
	trie := Constructor()
	for _, str := range strs {
		trie.Insert(str)
	}

	// Pick any string as the rest should have a
	// PREFIX in common with the selected one.
	return trie.GetLCP(strs[0], len(strs))
}
