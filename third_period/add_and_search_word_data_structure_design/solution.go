package add_and_search_word_data_structure_design

type WordDictionary struct {
	child map[byte]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		child: make(map[byte]*WordDictionary),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		if _, ok := node.child[word[i]]; !ok {
			node.child[word[i]] = &WordDictionary{child: make(map[byte]*WordDictionary{})}
		}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

}
