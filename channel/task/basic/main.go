package main

import (
	"fmt"
	"time"
)

// channel/basicのInt版
// Goroutine内で好きなIntをchannelを経由でMain Goroutineに渡し標準出力にPrintしてみる。
func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- 100
	}()

	num := <-ch
	fmt.Println(num)
}
