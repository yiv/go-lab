package main

import "fmt"

type (
	BetType int32
	UserID string
)

const (
	BetKing     BetType = iota //压龙
	BetQueen                   //压凤
	BetTie                     //压和
	BetMinister                //压三公
	BetTypeMax                 //所有压注类型数量
)

type BetPool [5]int64

func main() {
	var playerSettle = make(map[UserID]BetPool)
	playerSettle["uid"] = BetPool{}
	betPool := playerSettle["uid"]
	betPool[BetKing] = 500
	playerSettle["uid"] = betPool
	fmt.Println(playerSettle)
}
