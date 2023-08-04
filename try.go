package main

import (
	"fmt"
)

// "strconv"
// "time"
// "math"

func main() {
	st := []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}
	fmt.Println(removeComments(st))

}

// 两大类：在块内和不在块内，用bloak进行标记
func removeComments(source []string) (sult []string) {
	// 定义bloak用来记录是否在标注内
	var bloak bool = false
	// 当检查到// 时直接结束这个
	temp := []byte{}
	for _, v := range source {
		// var temp []byte = make([]byte, 0)

		length := len(v)
		// 当检查到/* bloak进行标记，直到*/进行解除
		for i := 0; i < length; i++ {
			if i < length-1 && v[i] == '/' && v[i+1] == '/' && bloak == false {
				break
			} else if i < length-1 && v[i] == '/' && v[i+1] == '*' {
				i++
				bloak = true
			} else if i < length-1 && bloak && v[i] == '*' && v[i+1] == '/' {
				i++
				bloak = false
			} else {
				if !bloak {
					temp = append(temp, v[i])
				}
			}
		}
		ltmp := string(temp)
		if ltmp != "" && !bloak {
			sult = append(sult, ltmp)
			temp = []byte{}
		}

	}
	return
}
