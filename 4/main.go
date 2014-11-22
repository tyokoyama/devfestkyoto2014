package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("*********** deferで呼ばれた無名関数だよ。 ***************")
	}()

	log.Println("main()が終わったよ。")
	// log.Fatalln("main()が異常終了！？")
}