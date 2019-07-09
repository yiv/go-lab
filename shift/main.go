package main

import "fmt"

func main() {
	MoveRight()
}

func MoveRight() {
	big := 256211515504331795
	div := 10000000000000000

	fmt.Println(big % div)

}
