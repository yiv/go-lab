package main

import (
	"fmt"
	"runtime"
)

func main() {
	BlockSchedule()
}

func BlockSchedule() {
	//有可能会出现这种情况，一个无耻的goroutine阻止其他goroutine运行。当你有一个不让调度器运行的for循环时，这就会发生。
	runtime.GOMAXPROCS(1)
	done := false
	go func() {
		done = true
	}()
	for !done {
		//for循环并不需要是空的。只要它包含了不会触发调度执行的代码，就会发生这种问题。

		//显式的唤起调度器
		//runtime.Gosched()
	}
	fmt.Println("done!")
}
