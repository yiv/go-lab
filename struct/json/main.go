package main

import (
	"encoding/json"
	"fmt"
)

type Player struct {
	UserID int64
	Nick   string
	Coin   int64
}
type Table struct {
	Seat    int
	Players map[int64]*Player
}

//主要测试结构里面的指针成员，在进行json encode的时候会不会有成员的详细信息
func main() {
	players := make(map[int64]*Player)
	players[55] = &Player{UserID: 55, Nick: "edwin", Coin: 5000}
	table := Table{Seat: 50, Players: players}
	js, err := json.Marshal(table)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
