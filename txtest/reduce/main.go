package main

import (
	"fmt"
	"math/rand"
	"time"
)

func StartTimes(arr []int32, cnt int32) []int32 {
	var start, end int
	max := len(arr)
	countMap := make(map[int32]int32)
	for {
		end = start + 100000
		if end > max {
			end = max
		}
		tmp := arr[start:end]
		for _, v := range tmp {
			countMap[v]++
		}
		if end == max {
			break
		}
	}

	var res []int32
	for k, v := range countMap {
		if v > cnt {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	arr := []int32{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100000; i++ {
		r := rand.Int31n(100)
		arr = append(arr, int32(r))
	}

	a := StartTimes(arr, 1000)
	fmt.Println("随机生成 100000 个 0 - 99 间的数字")
	fmt.Println("重复次数超过 1000 次的数字是： ", a)
}
