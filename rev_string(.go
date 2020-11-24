package main

import (
	"fmt"
	"strings"
)

// 字符串位置互换
func main() {
	s := "I am a student."
	fmt.Println(revString(s))
}

func revString(s string) string {
	ss := strings.Split(s, " ")
	for i, sss := range ss {
		ss[i] = revWord(sss)
	}
	return strings.Join(ss, " ")
}

func revWord(s string) string {
	b := []byte(s)
	l := len(b) - 1
	for i := 0; i < l; i++ {
		b[i], b[l] = b[l], b[i]
		l--
	}
	return string(b)
}
