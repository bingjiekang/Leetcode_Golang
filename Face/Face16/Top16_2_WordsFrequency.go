package main

// 设计一个方法，找出任意指定单词在一本书中的出现频率。
// 你的实现应该支持如下操作：
// WordsFrequency(book)构造函数，参数为字符串数组构成的一本书
// get(word)查询指定单词在书中出现的频率
// 示例：
// WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
// wordsFrequency.get("you"); //返回0，"you"没有出现过
// wordsFrequency.get("have"); //返回2，"have"出现2次
// wordsFrequency.get("an"); //返回1
// wordsFrequency.get("apple"); //返回1
// wordsFrequency.get("pen"); //返回1

// 解析:使用map存储对应单词出现的次数,当使用get时提取出即可

type WordsFrequency struct {
	str map[string]int
}

func Constructor(book []string) WordsFrequency {
	var temp WordsFrequency
	temp.str = make(map[string]int, 0)
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
