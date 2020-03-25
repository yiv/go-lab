package main

import "fmt"

func main() {
	words := "一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十一二三四五六七八九十"
	fmt.Println(len(words))
	fmt.Println(len([]rune(words)))
}
