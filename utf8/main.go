package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	nickName := "hello你"
	t := truncateUT8String(nickName, 7)
	fmt.Println(string(t))
}

func truncateUT8String(src string, max int) (dst string) {
	src = strings.Trim(src, " ")
	for i := 0; i < max; i++ {
		if len(src) <= 0 {
			return dst
		}
		r, size := utf8.DecodeRuneInString(src)
		src = src[size:]
		rs := make([]byte, size)
		utf8.EncodeRune(rs, r)
		dst += string(rs)
	}
	return
}

//func truncateUT8String(s []byte, l int) (t []byte) {
//	for len(s) > 0 {
//		r, size := utf8.DecodeRune(s)
//		s = s[size:]
//		rs := make([]byte, size)
//		utf8.EncodeRune(rs, r)
//		if (len(t) + size) > l {
//			return t
//		} else if (len(t) + size) == l {
//			return append(t, rs...)
//		} else {
//			t = append(t, rs...)
//		}
//	}
//	return
//}
