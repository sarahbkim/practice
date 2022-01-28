package main

type Trie struct {
	chars  map[rune]*Trie
	isWord bool
}

func NewTrie() Trie {
	return Trie{
		chars:  make(map[rune]*Trie, 26),
		isWord: false,
	}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	var curr = t
	for i, c := range word {
		if curr.chars[c] == nil {
			curr.chars[c] = &Trie{chars: make(map[rune]*Trie, 26), isWord: false}
		}
		if i == len(word)-1 {
			curr.isWord = true
		}
		curr = curr.chars[c]
	}
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	var curr = t
	for i, c := range word {
		if curr.chars[c] == nil {
			return false
		}
		if i == len(word)-1 && curr.isWord {
			return true
		}
		curr = curr.chars[c]
	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	var curr = t
	for _, c := range prefix {
		if curr.chars[c] == nil {
			return false
		}
		curr = curr.chars[c]
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
