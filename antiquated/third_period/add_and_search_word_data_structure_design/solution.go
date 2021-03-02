package add_and_search_word_data_structure_design

type WordDictionary struct {
	child []*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		child: make([]*WordDictionary, 26),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		a := word[i] - 'a'
		if node.child[a] == nil {
			node.child[a] = &WordDictionary{child: make([]*WordDictionary, 26)}
		}
		if i == len(word)-1 {
			node.child[a].isEnd = true
		}
		node = node.child[a]
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this
	return dfs(node, []byte(word))
}

func dfs(node *WordDictionary, word []byte) bool {
	if len(word) == 0 {
		return node.isEnd
	}
	a := word[0] - 'a'
	if word[0] != '.' && node.child[a] != nil {
		return dfs(node.child[a], word[1:])
	} else if word[0] == '.' {
		for k := range node.child {
			if node.child[k] != nil && dfs(node.child[k], word[1:]) {
				return true
			}
		}
	}
	return false
}
