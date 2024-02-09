package trie

type WordDictionary struct {
	wordEnd  bool
	children [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (wd *WordDictionary) AddWord(word string) {
	crawler := wd

	for _, ch := range word {
		idx := ch - 'a'
		if crawler.children[idx] == nil {
			crawler.children[idx] = &WordDictionary{}
		}
		crawler = crawler.children[idx]
	}

	crawler.wordEnd = true
}

func searchUtil(root *WordDictionary, word string) bool {
	crawler := root
	for in, ch := range word {
		if ch == '.' {
			for i := 0; i < 26; i++ {
				if crawler.children[i] != nil {
					if searchUtil(crawler.children[i], word[in+1:]) {
						return true
					}
				}
			}
			return false
		} else {
			idx := ch - 'a'
			if crawler.children[idx] == nil {
				return false
			}
			crawler = crawler.children[idx]
		}
	}
	return crawler != nil && crawler.wordEnd
}

func (wd *WordDictionary) Search(word string) bool {
	return searchUtil(wd, word)
}
