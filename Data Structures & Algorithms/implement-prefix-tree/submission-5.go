type TrieNode struct {
	child map[string]*TrieNode
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

func (this *PrefixTree) createLeft() TrieNode{
	return TrieNode{
		endOfWord: true,
	}
}

func (this *PrefixTree) Insert(word string) {
	if this.root == nil {
		this.root = &TrieNode{child: make(map[string]*TrieNode)}
	}

	node := this.root
	for _, w := range word {
		ch := string(w)
		if node.child[ch] == nil {
			node.child[ch] = &TrieNode{child: make(map[string]*TrieNode)}
		}
		node = node.child[ch]
	}
	node.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	if this.root == nil {
		return false
	}

	node := this.root
	for _, w := range word{
		ch := string(w)
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}

	return node.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	if this.root == nil {
		return false
	}

	node := this.root
	for _, p := range prefix{
		ch := string(p)
		if node.child[ch] == nil {
			return false
		}
		node = node.child[ch]
	}

	return true
}
