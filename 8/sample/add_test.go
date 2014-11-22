package sample

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain Start")

	// Run()を呼び出すことで、テストが実行される。
	m.Run()

	fmt.Println("TestMain End")
}

func TestAdd(t *testing.T) {
	fmt.Println("TestAdd Start")

	if x := Add(1, 1); x != 2 {
		t.Fatalf("1 + 1 = %d", x)
	}

	fmt.Println("TestAdd End")
}