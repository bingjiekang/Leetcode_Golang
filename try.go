package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	// var m int
	// dt := map[string][]string{}
	// tgname := []string{}
	// fmt.Scanf("%d", &m)

	// for i := 0; i < m; i++ {
	// 	var name, add string

	// 	fmt.Scanln(&name, &add)
	// 	_, ok := dt[name]
	// 	if !ok {
	// 		dt[name] = append(dt[name], add)
	// 		tgname = append(tgname, name)

	// 	} else {
	// 		var tag bool = true
	// 		for _, v := range dt[name] {
	// 			if add == v {
	// 				tag = false
	// 				break
	// 			}
	// 		}
	// 		if tag {
	// 			dt[name] = append(dt[name], add)

	// 		}
	// 	}
	// }
	// sort.Strings(tgname)
	// for _, v := range tgname {
	// 	fmt.Println(v, dt[v][0], len(dt[v])-1)
	// }

	test()
}

func test() {

	dict := map[string]string{"order_id": "20220201030210321",
		"amount":       "42",
		"notify_url":   "http://example.com/notify",
		"redirect_url": "http://example.com/redirect"}

	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var str string
	for i := 0; i < len(keys); i++ {
		x := keys[i]
		value := dict[x]
		str += x + "=" + value + "&"
	}
	param := str[0:len(str)-1] + "epusdt_password_xasddawqe"

	// fmt.Println(param)
	A := MD5(param)

	dict["signature"] = A
	fmt.Println(dict)

}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// type infomg struct{
//     address []string
//     length int
// }
