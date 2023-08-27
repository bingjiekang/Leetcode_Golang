package main

import (
	"fmt"
	"sort"
)

func main() {
	var m int
	dt := map[string][]string{}
	tgname := []string{}
	fmt.Scanf("%d", &m)

	for i := 0; i < m; i++ {
		var name, add string

		fmt.Scanln(&name, &add)
		_, ok := dt[name]
		if !ok {
			dt[name] = append(dt[name], add)
			tgname = append(tgname, name)

		} else {
			var tag bool = true
			for _, v := range dt[name] {
				if add == v {
					tag = false
					break
				}
			}
			if tag {
				dt[name] = append(dt[name], add)

			}
		}
	}
	sort.Strings(tgname)
	for _, v := range tgname {
		fmt.Println(v, dt[v][0], len(dt[v])-1)
	}

}

// type infomg struct{
//     address []string
//     length int
// }
