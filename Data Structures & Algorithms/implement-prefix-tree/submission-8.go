type TrieNode struct {
	child map[rune]*TrieNode
	endOfWord bool 
}
type PrefixTree struct {
	root *TrieNode 
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: nil,
	}
}	

func (this *PrefixTree) Insert(word string) {
	if this.root == nil {
		this.root = &TrieNode{child: make(map[rune]*TrieNode)}
	}

	node := this.root
	for _, w := range word {
		if node.child[w] == nil {
			node.child[w] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		node = node.child[w]
	}
	node.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	if this.root == nil {
		return false
	}

	node := this.root
	for _, w := range word{
		if node.child[w] == nil {
			return false
		}
		node = node.child[w]
	}

	return node.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	if this.root == nil {
		return false
	}

	node := this.root
	for _, p := range prefix{
		if node.child[p] == nil {
			return false
		}
		node = node.child[p]
	}

	return true
}
