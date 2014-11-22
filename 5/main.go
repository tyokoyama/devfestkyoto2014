package main

import (
	"errors"
	"fmt"
)

var (
	error1 = errors.New("エラー1")
	error2 = fmt.Errorf("エラー2")
)

func main() {
	err1 := errors.New("エラー1")
	err2 := fmt.Errorf("エラー2")
	err3 := error1
	err4 := error2

	fmt.Printf("err1 == error1の結果[%v]\n", err1 == error1)
	fmt.Printf("err2 == error2の結果[%v]\n", err2 == error2)
	fmt.Printf("err3 == error1の結果[%v]\n", err3 == error1)
	fmt.Printf("err4 == error2の結果[%v]\n", err4 == error2)
}