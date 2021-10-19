package main

func main() {

}

type WordDictionary struct {
	words []string
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	this.words = append(this.words, word)
}

func (this *WordDictionary) Search(word string) bool {
	matched := map[string]bool{}
	for _, w := range this.words {
		if len(w) != len(word) {
			continue
		}
		matched[w] = true
	}
	for i := range word {
		if len(matched) == 0 {
			return false
		}
		for k := range matched {
			if k[i] != word[i] && word[i] != '.' {
				delete(matched, k)
			}
		}
	}
	return len(matched) > 0
}
