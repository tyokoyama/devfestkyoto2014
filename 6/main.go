package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}

	// 今まではこんな感じだった。
	// いちいち、"_"を指定して受け取らないことを明示する。
	i := 1
	for _ = range array {
		fmt.Printf("%d回目\n", i)
		i++
	}

	fmt.Println("***********************")

	// 1.4からは
	i = 1
	for range array {
		fmt.Printf("%d回目\n", i)
		i++
	}
}