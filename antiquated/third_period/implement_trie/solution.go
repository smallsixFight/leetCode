package implement_trie

type Trie struct {
	Child map[byte]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Child: make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		if _, ok := node.Child[word[i]]; !ok {
			node.Child[word[i]] = &Trie{Child: make(map[byte]*Trie)}
		}
		if i == len(word)-1 {
			node.Child[word[i]].isEnd = true
		}
		node = node.Child[word[i]]
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		if _, ok := node.Child[word[i]]; !ok {
			break
		}
		if i == len(word)-1 && node.Child[word[i]].isEnd {
			return true
		}
		node = node.Child[word[i]]
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := node.Child[prefix[i]]; !ok {
			return false
		}
		node = node.Child[prefix[i]]
	}
	return true
}
