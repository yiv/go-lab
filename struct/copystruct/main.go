//try to use "=" to copy value of different struct type
//fail
package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type UserId string

type Card byte
type Cards []Card
type Cards2 []byte
type Seat int32

type Big struct {
	Uid    string
	Nick   string
	Cards  Cards
	Cards2 Cards2
	Seat   Seat
}
type Small struct {
	Uid    string
	Cards  []byte
	Cards2 []byte
	Seat   int32
}

func main() {
	CopyArrayOfStruct()
}

func RenameType() {
	type User1 struct {
		Uid       UserId
		Friends   map[UserId]int32
		Relatives map[string]int32
	}
	type User2 struct {
		Uid       string
		Friends   map[string]int32
		Relatives map[string]int32
	}
	u1 := &User1{Uid: "edwin", Friends: map[UserId]int32{"nick": 30}, Relatives: map[string]int32{"padme": 30}}
	u2 := &User2{}
	copier.Copy(u2, u1)
	fmt.Printf("%#v", u2)
}

func Nest1() {
	cards := Cards{Card(0x15), Card(0x11)}
	var bytes []byte
	copier.Copy(&bytes, &cards)
	fmt.Println(bytes)
}

func Nest2() {
	big := Big{Uid: "22222", Nick: "edwin", Cards: Cards{Card(0x15), Card(0x11)}, Cards2: Cards2{0x1}, Seat: Seat(33)}
	small := Small{}
	copier.Copy(&small, &big)
	fmt.Println(small)
}

func Nest3() {
	type LordRecord struct {
		Ante   int32
		Times  int32
		Settle int32
		Ctime  int32
	}
	type GetLordRecordsRes struct {
		Records []*LordRecord
	}

	type LordRecord2 struct {
		Ante   int32
		Times  int32
		Settle int32
		Ctime  int32
	}
	type GetLordRecordsRes2 struct {
		Records []LordRecord2
	}
	var records = []*LordRecord{&LordRecord{Ante: 5}}
	res := &GetLordRecordsRes{Records: records}
	var res2 GetLordRecordsRes2
	copier.Copy(&res2, res)
	fmt.Println("%v", res2)
}

func CopyArrayOfStruct() {
	type IntArray []int
	type Account struct {
		Name string
		Age  int
		Attr []int
	}
	type Account2 struct {
		Name string
		Age  int
		Attr IntArray
	}

	origin := []Account{{Name: "haha", Age: 5, Attr: IntArray{1, 2, 3}}, {Name: "2222", Age: 5, Attr: IntArray{4, 5, 6}}}
	var dest []Account2
	copier.Copy(&dest, &origin)
	fmt.Println(dest)
}
