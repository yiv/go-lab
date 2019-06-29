package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Coord struct {
	x int32
	y int32
}

type CoordDistance struct {
	index int
	dis   float64
}
type Distances []CoordDistance

func (d Distances) Len() int           { return len(d) }
func (d Distances) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Distances) Less(i, j int) bool { return d[i].dis < d[j].dis }

func GetNearestCoord(vec []Coord, e Coord, num int) []Coord {
	var temp Distances
	var nearest []Coord
	for k, v := range vec {
		dis := v.Distance(e)
		temp = append(temp, CoordDistance{index: k, dis: dis})
	}
	sort.Sort(temp)
	for k, v := range temp {
		if k < num {
			nearest = append(nearest, vec[v.index])
		}
	}
	return nearest
}

func (c Coord) Distance(other Coord) (dis float64) {
	x := math.Abs(float64(other.x) - float64(c.x))
	y := math.Abs(float64(other.y) - float64(c.y))
	dis = math.Sqrt(x*x + y*y)
	return
}

func main() {
	rand.Seed(time.Now().Unix())
	var vecs []Coord
	for i := 0; i < 50; i++ {
		x := rand.Int31n(500)
		y := rand.Int31n(500)
		vecs = append(vecs, Coord{x: x, y: y})
	}
	mypos := Coord{x: rand.Int31n(500), y: rand.Int31n(500)}
	near := GetNearestCoord(vecs, mypos, 5)
	fmt.Println("随机生成 50 个坐标，它们的 x 和 y 值范围是 0 - 499")
	fmt.Println("我的坐标是：", mypos)
	fmt.Println("获取离我最近的 5 个坐标及离我的距离，：", mypos)
	for _, v := range near {
		fmt.Println(v, " = ", mypos.Distance(v))
	}
}
