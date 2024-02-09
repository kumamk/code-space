package trie

type Trie struct {
	wordEnd  bool
	children [26]*Trie
}

func New() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	crawler := trie
	for _, ch := range word {
		idx := ch - 'a'
		if crawler.children[idx] == nil {
			crawler.children[idx] = &Trie{}
		}
		crawler = crawler.children[idx]
	}
	crawler.wordEnd = true
}

func (trie *Trie) Search(word string) bool {
	crawler := trie
	for _, ch := range word {
		idx := ch - 'a'
		if crawler.children[idx] == nil {
			return false
		}
		crawler = crawler.children[idx]
	}
	if crawler != nil && crawler.wordEnd {
		return true
	}
	return false
}

func (trie *Trie) StartsWith(prefix string) bool {
	crawler := trie
	i := 0
	for _, ch := range prefix {
		idx := ch - 'a'
		if crawler.children[idx] == nil {
			return false
		}
		crawler = crawler.children[idx]
		i += 1
	}
	if crawler != nil && i == len(prefix) {
		return true
	}
	return false
}
