package main

import (
	"fmt"
	"io"
	"sort"
)

// 用来读取有含有空格字符串 []byte
// reader := bufio.NewReader(os.Stdin)
// tmp, _, _ := reader.ReadLine()

// 试题概述:小红书的搜索引擎根据搜索词分词后得到很多关键词,统计关键词出现次数大于3次的(出现次数多到出现次数低),如果出现次数相同则按字典序升序排序

// 解题思路:读取[xsfas fasdfsd fasdf...]关键词后,利用map遍历存储出现关键词的次数,使用sort函数排序,对出现次数进行排序,如果次数相同则根据字典序排序,输入次数大于等于3次的关键词

func main() {
	// 定义一个切片用来存储输入的全部关键词
	var tmp []string = make([]string, 0)
	// 用来获得 每一次的输入信息
	var st string
	// 定义一个map用来存储关键词出现的次数
	var hmap map[string]int = make(map[string]int, 0)
	// 循环遍历等待输入信息
	for {
		_, err := fmt.Scan(&st)
		if err != nil {
			// 如果输入结束则停止输入
			if err == io.EOF {
				break
			}
		} else { // 将关键词加入到切片中存储
			tmp = append(tmp, st)
		}
	}
	// fmt.Println(tmp)
	// map存储关键词出现的次数
	for _, v := range tmp {
		hmap[v]++
	}
	// 定义一个结构体类型切片,用来将对应map中的key和value,对应到结构体mp中的name和num
	var temp []mp = make([]mp, 0)
	// 遍历存储关键词和对应次数的map
	for k, v := range hmap {
		// 定义一个临时结构体存储key和value,方便后续加入
		tp := mp{k, v}
		// 将每一个关键词和次数对应的结构体加入到temp切片中
		temp = append(temp, tp)
	}
	// 对temp进行排序
	sort.Sort(Ed(temp))
	// 遍历排序后的切片,当出现的次数大于3时输出
	for _, v := range temp {
		if v.num >= 3 {
			fmt.Println(v.name)
		} else {
			break
		}
	}

}

// 下面是用来实现sort.Sort()函数的接口,当结构体实现了Len()int,Less(n,m int)bool,Swap(n,m int) 后就可以进行排序

// name记录map中的key即关键词,num记录map中的value即次数
type mp struct {
	name string
	num  int
}

// 定义一个切片类型结构体,Ed实现了Len/Less/Swap方法,就可以使用sort.Sort调用这个Ed
type Ed []mp

// 返回这个切片结构体的长度
func (this Ed) Len() int {
	return len(this)
}

// Less用来实现升序或降序
func (this Ed) Less(n, m int) bool {
	// 根据关键词出现的次数实现降序排列
	if this[n].num > this[m].num {
		return true
	} else if this[n].num == this[m].num { // 如果出现的次数相同,则根据出现的关键词进行升序排列
		if this[n].name < this[m].name { // 字符串比较,谁字典序在前谁先输出
			return true
		} else {
			return false
		}
	} else {
		return false // 否则返回false
	}
}

// 用来交换对应切片里的name/num即关键词和出现的次数
func (this Ed) Swap(n, m int) {
	this[n].name, this[m].name = this[m].name, this[n].name
	this[n].num, this[m].num = this[m].num, this[n].num
}
