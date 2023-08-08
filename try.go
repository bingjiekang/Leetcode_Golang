package main

type WordsFrequency struct {
	str map[string]int
}

func Constructor(book []string) WordsFrequency {
	var temp WordsFrequency
	for _, v := range book {
		temp.str[v]++
	}
	return temp
}

func (this *WordsFrequency) Get(word string) int {
	_, ok := this.str[word]
	if ok {
		return this.str[word]
	}
	return 0
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
