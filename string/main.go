package main

import (
	"fmt"
)

func main() {
	getLen()
}

func getLen() {
	str := "abcdefghijklmnopqrstuvwxyz123456abcdefghijklmnopqrstuvwxyz123456"
	fmt.Println("len(str) ", len(str))
	fmt.Println("len(rune) ", len([]rune(str)))
	fmt.Println("3 rune ", string([]rune(str)[0:3]))
}

func distinct() {
	numbers := []int64{15749417216928, 15791695208487, 15749417659653, 15749417803399, 15772581542441, 15772581542441, 15772581542441, 15755297837831, 15749417795121, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417769862, 15755297839936, 15749417659653, 15749417803399, 15772581542441, 15791695208487, 15772581544665, 15772581544665, 15772581544665, 15749417798690, 15772581542441, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417769459, 15749417222933, 15755297837831, 15749417782586, 15749417853013, 15749417659653, 15749417801028, 15749417853013, 15749417853013, 15749417221170, 15749417853013, 15749417853013, 15749417853013, 15749417801028, 15749417853013, 15749417853013, 15772581542441, 15749417770529, 15756307313207, 15749417769459, 15756307316841, 15756307319671, 15756307322893, 15749417659653, 15749417799969, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417801028, 15749417853013, 15749417853013, 15749417798690, 15772581542441, 15791695208487, 15750131361742, 15749417853013, 15749417843948, 15772581547978, 15749417790016, 15772581542441, 15749417843948, 15772581547978, 15749417790016, 15772581542441, 15772581542441, 15749417811480, 15749417853013, 15749417853013, 15749417843948, 15749417853013, 15772581547978, 15749417811480, 15772581547978, 15755297837831, 15749417811480, 15749417853013, 15750131361742, 15749417791267, 15772581542441, 15749417853013, 15749417770026, 15750131361742, 15749417769212, 15749417793896, 15772581542441, 15749417843948, 15772581547978, 15749417793896, 15772581542441, 15749417843948, 15772581547978, 15749417793896, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417810194, 15772581547978, 15749417788831, 15772581542441, 15749417810194, 15749417853013, 15749417853013, 15772581547978, 15749417853013, 15749417853013, 15749417790016, 15772581542441, 15749417853013, 15749417788831, 15772581542441, 15749417843948, 15749417853013, 15772581547978, 15749417788831, 15772581542441, 15749417811480, 15749417853013, 15772581547978, 15749417843948, 15772581547978, 15749417811480, 15772581547978, 15772581547978, 15749417843948, 15772581547978, 15749417787809, 15772581542441, 15791695208487, 15750131361742, 15749417769212, 15749417795121, 15749417853013, 15749417793896, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417811480, 15749417853013, 15749417853013, 15749417810194, 15772581547978, 15749417853013, 15749417791267, 15772581542441, 15772581542441, 15749417806025, 15749417853013, 15772581547978, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15749417810194, 15749417853013, 15749417787809, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15772581542441, 15749417810194, 15749417853013, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15749417811480, 15749417853013, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15749417790016, 15749417853013, 15749417853013, 15749417853013, 15749417811480, 15749417853013, 15772581547978, 15772581547978, 15749417790016, 15749417853013, 15749417810194, 15749417853013, 15772581547978, 15749417790016, 15772581542441, 15749417853013, 15749417791267, 15772581542441, 15749417810194, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15749417843948, 15749417853013, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15791695208487, 15750131361742, 15749417792769, 15749417853013, 15749417853013, 15749417787809, 15772581542441, 15749417791267, 15749417853013, 15749417845100, 15772581547978, 15772581547978, 15749417853013, 15749417845100, 15772581547978, 15749417791267, 15772581542441, 15749417845100, 15772581547978, 15749417791267, 15772581542441, 15772581542441, 15749417793896, 15772581542441, 15772581542441, 15749417660749, 15749417795121, 15749417853013, 15749417853013, 15772581542441, 15749417845100, 15772581547978, 15749417795121, 15749417853013, 15749417768863, 15749417790016, 15749417853013, 15791695208487, 15772581542441, 15772581542441, 15772581542441, 15749417853013, 15749417792769, 15749417853013, 15772581542441, 15749417853013, 15749417791267, 15749417853013, 15749417853013, 15772581542441, 15772581542441, 15749417853013, 15749417795121, 15749417853013, 15749417853013, 15749417819357, 15749417791267, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417819357, 15772581547978, 15749417790016, 15749417853013, 15749417853013, 15749417853013, 15749417819357, 15772581547978, 15772581547978, 15772581547978, 15749417790016, 15772581542441, 15749417816624, 15772581547978, 15749417787809, 15772581542441, 15749417819357, 15749417853013, 15772581547978, 15749417790016, 15749417853013, 15749417770026, 15749417660749, 15749417769459, 15749417787809, 15772581542441, 15749417819357, 15772581547978, 15749417790016, 15749417853013, 15749417819357, 15749417853013, 15749417813994, 15749417853013, 15749417791267, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417791267, 15749417853013, 15772581542441, 15772581542441, 15749417790016, 15772581542441, 15772581542441, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15772581542441, 15749417819357, 15772581547978, 15749417853013, 15749417787809, 15749417853013, 15749417819357, 15772581547978, 15749417787809, 15772581542441, 15791695208487, 15749417853013, 15749417853013, 15749417778075, 15772581544665, 15772581544665, 15749417816624, 15749417853013, 15772581547978, 15749417778075, 15772581544665, 15749417816624, 15772581547978, 15772581547978, 15749417812753, 15749417853013, 15749417787809, 15772581542441, 15772581542441, 15749417853013, 15749417660749, 15791695208487, 15749417790016, 15772581542441, 15772581542441, 15749417853013, 15749417788831, 15772581542441, 15772581542441, 15749417853013, 15749417817953, 15749417853013, 15749417853013, 15749417819357, 15772581547978, 15749417853013, 15772581547978, 15749417791267, 15772581542441, 15749417819357, 15772581547978, 15772581547978, 15749417790016, 15772581542441, 15749417853013, 15749417853013, 15749417788831, 15772581542441, 15749417793896, 15772581542441, 15749417853013, 15749417790016, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417819357, 15749417853013, 15749417853013, 15749417790016, 15749417853013, 15749417819357, 15749417853013, 15749417817953, 15772581547978, 15749417853013, 15749417792769, 15749417853013, 15772581542441, 15749417812753, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15749417853013, 15750131361742, 15749417790016, 15749417853013, 15749417853013, 15749417792769, 15772581542441, 15749417791267, 15772581542441, 15772581542441, 15749417853013, 15749417664399, 15749417853013, 15749417790016, 15749417853013, 15772581544665, 15772581544665, 15749417790016, 15749417853013, 15749417853013, 15749417663076, 15749417853013, 15749417774225, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15772581544665, 15749417770026, 15749417663076, 15791695208487, 15749417853013, 15749417787809, 15772581542441, 15749417782586, 15749417853013, 15749417853013, 15749417774225, 15772581544665, 15749417787809, 15772581542441, 15749417774225, 15749417853013, 15749417781490, 15772581547978, 15749417774225, 15749417853013, 15749417795121, 15749417853013, 15772581542441, 15749417782586, 15749417853013, 15749417853013, 15772581547978, 15749417853013, 15749417795121, 15749417853013, 15749417769779, 15772581542441, 15749417780380, 15749417853013, 15749417790016, 15749417853013, 15749417853013, 15749417780380, 15749417853013, 15749417853013, 15749417781490, 15772581547978, 15772581547978, 15749417792769, 15772581542441, 15749417781490, 15772581547978, 15749417792769, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417792769, 15749417853013, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417782586, 15749417853013, 15772581547978, 15772581547978, 15749417853013, 15749417782586, 15749417853013, 15749417853013, 15749417781490, 15772581547978, 15749417792769, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417664399, 15749417811480, 15749417791267, 15749417853013, 15749417853013, 15749417811480, 15749417853013, 15749417853013, 15749417853013, 15749417811480, 15772581547978, 15749417853013, 15772581547978, 15749417791267, 15772581542441, 15749417790016, 15772581542441, 15749417811480, 15749417853013, 15772581547978, 15749417791267, 15772581542441, 15749417811480, 15772581547978, 15772581547978, 15749417791267, 15772581542441, 15749417804717, 15772581547978, 15772581547978, 15749417791267, 15772581542441, 15772581542441, 15749417664399, 15749417769779, 15749417790016, 15749417853013, 15749417663076, 15749417782586, 15749417853013, 15772581547978, 15749417774225, 15749417853013, 15749417782586, 15749417853013, 15749417774225, 15749417853013, 15749417664399, 15749417790016, 15749417853013, 15749417853013, 15749417774225, 15772581544665, 15749417782586, 15749417853013, 15772581547978, 15749417774225, 15772581544665, 15749417782586, 15749417853013, 15772581547978, 15749417774225, 15749417853013, 15749417782586, 15749417853013, 15749417774225, 15772581544665, 15749417853013, 15772581544665, 15749417781490, 15749417853013, 15772581547978, 15772581547978, 15749417810194, 15772581547978, 15772581547978, 15772581547978, 15749417791267, 15749417853013, 15772581542441, 15749417781490, 15749417853013, 15749417774225, 15749417853013, 15772581544665, 15772581544665, 15749417811480, 15749417853013, 15772581547978, 15749417774225, 15749417853013, 15749417853013, 15749417790016, 15749417853013, 15772581542441, 15749417780380, 15749417853013, 15749417853013, 15749417787809, 15772581542441, 15772581542441, 15772581542441, 15772581542441, 15749417781490, 15772581547978, 15772581547978, 15749417853013, 15749417810194, 15772581547978, 15749417781490, 15772581547978, 15749417792769, 15749417853013, 15749417787809, 15772581542441, 15772581542441, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417792769, 15749417853013, 15749417790016, 15749417853013, 15749417853013, 15772581542441, 15772581542441, 15749417787809, 15772581542441, 15749417853013, 15749417660749, 15791695208487, 15749417853013, 15749417811480, 15772581547978, 15749417787809, 15772581542441, 15749417811480, 15772581547978, 15749417853013, 15749417790016, 15749417853013, 15749417853013, 15772581542441, 15749417811480, 15749417853013, 15772581547978, 15749417853013, 15749417853013, 15772581547978, 15749417787809, 15772581542441, 15754678863089, 15749417792769, 15749417853013, 15749417769862, 15755297839936, 15749417853013, 15749417770614, 15754678863089, 15772581544665, 15749417792769, 15772581542441, 15772581544665, 15772581544665, 15749417218041, 15749417770614, 15755522147428, 15749417853013, 15749417665655, 15749417795121, 15749417853013, 15749417853013, 15749417773014, 15749417853013, 15749417793896, 15749417853013, 15749417853013, 15749417769376, 15772581544665, 15749417848652, 15749417853013, 15772581544665, 15749417853013, 15749417771919, 15749417853013, 15772581547978, 15749417770614, 15749417791267, 15772581542441, 15749417853013, 15749417773014, 15749417853013, 15772581547978, 15772581547978, 15749417791267, 15749417853013, 15749417835304, 15749417791267, 15772581542441, 15772581542441, 15772581542441, 15749417834070, 15749417853013, 15772581547978, 15749417770026, 15749417665655, 15749417769121, 15749417790016, 15772581542441, 15749417853013, 15749417853013, 15749417834070, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15772581542441, 15749417834070, 15749417853013, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15749417834070, 15772581547978, 15749417787809, 15772581542441, 15749417773014, 15749417853013, 15772581547978, 15749417834070, 15749417853013, 15749417773014, 15749417853013, 15749417665655, 15749417769121, 15749417787809, 15749417853013, 15772581542441, 15749417834070, 15749417853013, 15772581547978, 15749417787809, 15772581542441, 15749417834070, 15772581547978, 15749417787809, 15772581542441, 15749417834070, 15749417853013, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15772581542441, 15749417834070, 15772581547978, 15749417790016, 15772581542441, 15749417853013, 15749417773014, 15772581547978, 15749417787809, 15772581542441, 15749417853013, 15749417665655, 15749417853013, 15749417770026, 15749417665655, 15749417770614, 15749417795121, 15749417853013, 15749417665655, 15749417770614, 15749417795121, 15749417853013, 15749417666921, 15749417847402, 15749417769376, 15749417853013, 15749417847402, 15772581544665, 15749417791267, 15772581542441, 15749417847402, 15749417853013, 15772581544665, 15772581544665, 15749417791267, 15772581542441, 15772581542441, 15749417770614, 15772581542441, 15749417847402, 15749417853013, 15772581544665, 15749417788831, 15772581542441, 15749417847402, 15772581544665, 15772581544665, 15749417793896, 15772581542441, 15755297839936, 15749417853013, 15749417853013, 15749417666921, 15749417787809, 15772581542441, 15749417853013, 15749417791267, 15749417853013, 15749417853013, 15749417769296, 15749417827619, 15749417853013, 15749417853013, 15749417853013, 15772581547978, 15749417788831, 15772581542441, 15749417791267, 15772581542441, 15749417787809, 15772581542441, 15749417827619, 15749417853013, 15772581547978, 15749417853013, 15749417787809, 15772581542441, 15749417824935, 15749417853013, 15749417853013, 15772581547978, 15772581547978, 15749417853013, 15749417853013, 15749417790016, 15749417853013, 15749417787809, 15772581542441, 15749417853013, 15772581542441, 15749417668300, 15749417795121, 15749417769599, 15749417853013, 15772581547978, 15749417770026, 15749417668300, 15749417770026, 15749417791267, 15772581542441, 15749417853013, 15749417668300, 15749417769376, 15749417795121, 15772581542441, 15772581542441, 15749417853013, 15772581547978, 15749417788831, 15772581542441, 15749417853013, 15772581547978, 15772581547978, 15749417787809, 15772581542441, 15749417666921, 15749417795121, 15749417853013, 15749417669628, 15749417853013, 15749417666921, 15749417791267, 15749417853013, 15772581542441, 15749417795121, 15772581542441, 15749417669628, 15749417769121, 15749417795121, 15772581542441, 15749417853013, 15749417853013, 15749417842463, 15749417853013, 15749417790016, 15772581542441, 15772581542441, 15772581542441, 15772581542441, 15749417837827, 15749417853013, 15749417853013, 15749417770026, 15749417669628, 15749417769296, 15749417795121, 15749417853013, 15756307327684, 15749417669628, 15749417787809, 15749417853013, 15749417853013, 15749417837827, 15749417853013, 15772581547978, 15749417787809, 15772581542441, 15749417842463, 15749417853013, 15749417787809, 15749417853013, 15749417669628, 15749417769296, 15749417790016, 15749417853013, 15749417842463, 15749417853013, 15749417853013, 15772581547978, 15749417787809, 15772581542441, 15772581542441, 15749417842463, 15749417853013, 15749417853013, 15749417787809, 15749417853013, 15749417669628, 15749417842463, 15749417853013, 15772581547978, 15749417787809, 15749417853013, 15772581542441, 15749417853013, 15749417853013, 15749417842463, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417787809, 15749417853013, 15749417668300, 15749417769296, 15749417787809, 15749417853013, 15772581542441, 15749417853013, 15749417666921, 15749417835304, 15749417853013, 15749417853013, 15772581547978, 15749417668300, 15749417787809, 15772581542441, 15749417666921, 15749417835304, 15749417853013, 15772581547978, 15749417668300, 15749417787809, 15772581542441, 15749417853013, 15749417666921, 15749417832600, 15749417853013, 15772581547978, 15749417668300, 15749417787809, 15772581542441, 15749417666921, 15749417832600, 15749417853013, 15772581547978, 15749417853013, 15749417668300, 15749417787809, 15749417853013, 15772581542441, 15772581542441, 15749417668300, 15749417769296, 15749417790016, 15749417853013, 15749417853013, 15749417853013, 15749417666921, 15755432902752, 15749417853013, 15749417853013, 15749417787809, 15749417853013, 15772581542441, 15755432905696, 15772581547978, 15749417853013, 15754678863089, 15749417769035, 15749417787809, 15749417853013, 15749417659653, 15749417791267, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417659653, 15749417769035, 15749417787809, 15749417853013, 15749417853013, 15772581542441, 15749417674615, 15749417662001, 15749417853013, 15749417659653, 15749417791267, 15749417853013, 15749417853013, 15749417853013, 15749417853013, 15749417793896, 15749417853013, 15772581542441, 15749417853013, 15772581542441, 15749417659653, 15749417770614, 15749417793896, 15749417853013, 15772581542441, 15749417853013, 15772581542441, 15750131361742, 15749417770614, 15749417795121, 15749417853013, 15749417853013, 15749417853013, 15749417770026, 15750131361742, 15749417770614, 15749417795121, 15772581542441, 15749417853013, 15749417770026, 15750131361742, 15749417770614, 15749417795121, 15772581542441, 15749417853013, 15749417770026, 15750131361742, 15749417769296, 15749417853013, 15750131361742, 15749417791267, 15749417853013, 15749417811480, 15749417853013, 15772581547978, 15772581547978, 15749417791267, 15772581542441, 15749417853013, 15749417811480, 15749417853013, 15749417853013, 15749417853013, 15772581547978, 15749417810194, 15749417853013, 15772581547978, 15750131361742, 15749417769296, 15749417795121, 15749417853013, 15749417847402, 15772581544665, 15749417791267, 15772581542441, 15749417847402, 15749417853013, 15772581544665, 15749417791267, 15749417853013, 15772581542441, 15749417847402, 15772581544665, 15749417790016, 15772581542441, 15772581542441, 15750131361742, 15749417795121, 15749417853013, 15772581544665, 15749417853013, 15749417790016, 15772581542441, 15772581542441, 15750131361742, 15749417795121, 15749417853013, 15749417853013, 15749417768952, 15749417811480, 15749417853013, 15749417853013, 15749417810194, 15772581547978, 15750131361742, 15749417810194, 15749417853013, 15749970580652, 15749417853013}
	nmap := map[int64]int64{}
	for _, v := range numbers {
		nmap[v] = v
	}
	for _, v := range nmap {
		fmt.Println(v)
	}
}
