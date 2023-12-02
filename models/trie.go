package models

type Trie struct {
	isWord bool
	children [26]*Trie
}

func ConstructTrie() Trie {
	return Trie{
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string)  {
	hc := this
	for i:=0; i<len(word); i++ {
		inx := word[i]-'a'
		if hc.children[inx] == nil {
			hc.children[inx] = &Trie{
				children: [26]*Trie{},
			}
		}
		hc = hc.children[inx]
	}
	hc.isWord = true
}

func (this *Trie) Search(word string) bool {
	hc := this
	for i:=0; i<len(word); i++ {
		inx := word[i]-'a'
		if hc.children[inx] == nil {
			return false
		}
		hc = hc.children[inx]
	}
	return hc.isWord
}


func (this *Trie) StartsWith(prefix string) bool {
	hc := this
	for i:=0; i<len(prefix); i++ {
		inx := prefix[i]-'a'
		if hc.children[inx] == nil {
			return false
		}
		hc = hc.children[inx]
	}
	return true
}