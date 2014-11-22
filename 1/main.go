package main

import (
	"fmt"
//	"runtime"
)

func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		go func (i int, ch chan<- int) {
			fmt.Printf("goroutine[%d] called.\n", i)
			ch <- i
		}(i, ch)
	}

	fmt.Printf("main wait start.\n")
	count := 0
	for {
		if count >= 10 {
			break
		}
		fmt.Printf("Channel received[%d]\n", <-ch)
		count++
	}
}