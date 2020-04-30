package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var arr []int

	ch := sleepSort([]int{2, 3, 5, 3, 1, 4, 10})

	for n := range ch {
		arr = append(arr, n)
	}

	fmt.Println(arr)
}

func sleepSort(nums []int) <-chan int {
	ch := make(chan int, 5)

	go func() {
		defer close(ch)

		var wg sync.WaitGroup

		for _, num := range nums {
			wg.Add(1)

			n := num
			go func() {
				defer wg.Done()
				time.Sleep(time.Duration(n) * time.Millisecond)

				ch <- n
			}()
		}

		wg.Wait()
	}()

	return ch
}
