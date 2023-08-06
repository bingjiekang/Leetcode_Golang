package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var tmp []string = make([]string, 0)
	var st string
	var hmap map[string]int = make(map[string]int, 0)

	for {
		_, err := fmt.Scan(&st)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			tmp = append(tmp, st)
		}
	}
	// fmt.Println(tmp)
	for _, v := range tmp {
		hmap[v]++
	}

	var temp []mp = make([]mp, 0)
	for k, v := range hmap {
		tp := mp{k, v}
		temp = append(temp, tp)
	}
	sort.Sort(Ed(temp))
	for _, v := range temp {
		if v.num >= 3 {
			fmt.Println(v.name)
		} else {
			break
		}
	}

}

type mp struct {
	name string
	num  int
}

type Ed []mp

func (this Ed) Len() int {
	return len(this)
}

func (this Ed) Less(n, m int) bool {
	if this[n].num > this[m].num {
		return true
	} else if this[n].num == this[m].num {
		if this[n].name < this[m].name {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (this Ed) Swap(n, m int) {
	this[n].name, this[m].name = this[m].name, this[n].name
	this[n].num, this[m].num = this[m].num, this[n].num
}
