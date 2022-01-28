package lc211

type WordDictionary struct {
	Root *Trie
	Size int
}

type Trie struct {
	Children map[byte]*Trie
	Word     string
}

func NewTrie() *Trie {
	return &Trie{Children: make(map[byte]*Trie)}
}

func SearchTrieRecur(t *Trie, word string, idx int) bool {
	// breaking condition
	if idx >= len(word) {
		return false
	}
	// breaking condition
	char := word[idx]
	if len(word) == idx+1 {
		if (word == t.Word) || (char == '.' && t.Word != "") {
			return true
		}
		if _, ok := t.Children[char]; ok {
			if t.Word != "" {
				return true
			}
		}

		return false
	}

	if char == '.' {
		for _, v := range t.Children {
			if SearchTrieRecur(v, word, idx+1) {
				return true
			}
		}
	}
	next, ok := t.Children[char]
	if !ok {
		return false
	}

	return SearchTrieRecur(next, word, idx+1)
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &Trie{
			Children: make(map[byte]*Trie),
		},
		Size: 0,
	}
}

func (this *WordDictionary) AddWord(word string) {
	curNode := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		val, ok := curNode.Children[char]
		if !ok {
			val = NewTrie()
			curNode.Children[char] = val
		}
		if len(word)-1 == i {
			if curNode.Word == word {
				return
			}
			curNode.Word = word
			this.Size++
		}
		curNode = val
	}
}

func (this *WordDictionary) Search(word string) bool {
	if this.Size == 0 {
		return false
	}
	if word == "" {
		return false
	}
	curNode := this.Root
	return SearchTrieRecur(curNode, word, 0)
}

/**
* Your WordDictionary object will be instantiated and called as such:
* obj := Constructor();
* obj.AddWord(word);
* param_2 := obj.Search(word);
 */
