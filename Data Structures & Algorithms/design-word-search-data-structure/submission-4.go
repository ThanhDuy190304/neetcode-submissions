type TrieNode struct{
	child map[rune]*TrieNode
	endOfWord bool
}
type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: nil,
	}
}

func (this *WordDictionary) AddWord(word string)  {
	if this.root == nil{
		this.root = &TrieNode{child: make(map[rune]*TrieNode)}
	}
	node := this.root
	for _, w := range word{
		if node.child[w] == nil {
			node.child[w] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		node = node.child[w] 
	}

	node.endOfWord = true
}

type frame struct {
	index     int
	node      *TrieNode
	children  []*TrieNode
	nextChild int
}


func (this *WordDictionary) Search(word string) bool {
	stack := []frame{{index: 0, node: this.root}}

	for len(stack) > 0 {
		top := &stack[len(stack)-1] 

		if top.node == nil {
			stack = stack[:len(stack)-1]
			continue
		}
		if top.index == len(word) {
			if top.node.endOfWord {
				return true
			}
			stack = stack[:len(stack)-1]
			continue
		}

		c := rune(word[top.index])
		if c != '.' {
			nextNode := top.node.child[c] 
			stack = stack[:len(stack)-1]
			stack = append(stack, frame{index: top.index + 1, node: nextNode})
		} else {
			if top.children == nil {
				top.children = make([]*TrieNode, 0, len(top.node.child))
				for _, child := range top.node.child {
					top.children = append(top.children, child)
				}
			}
			if top.nextChild >= len(top.children) {
				stack = stack[:len(stack)-1]
				continue
			}
			child := top.children[top.nextChild]
			top.nextChild++
			stack = append(stack, frame{index: top.index + 1, node: child})
		}
	}
	return false
}
