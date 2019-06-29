package main

import (
	"fmt"
	"strings"
	"sync"
)

func LineNum(file []string, substr string) int {
	var (
		outStrs   []string
		lineCount int
	)
	outChan := make(chan string)
	doneChan := make(chan struct{})
	wg := sync.WaitGroup{}
	for _, v := range file {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			if strings.Contains(str, substr) {
				outChan <- str
			}
		}(v)
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	for {
		select {
		case v := <-outChan:
			outStrs = append(outStrs, v)
		case <-doneChan:
			lineCount = len(outStrs)
			return lineCount
		}
	}
}

func main() {
	file := []string{"A Builderisused ", "toefficientlybuild ", "astringusingWritemethods"}
	fmt.Println("有几行数据分别是：", file)
	fmt.Println("包含 en 字符的行数为 ：", LineNum(file, "en"))
}
