package main

import (
	"fmt"
	"testing"
)

func TestDeferOrder1(t *testing.T) {
	// 打印顺序：B A
	defer fmt.Println("A")
	fmt.Println("B")
}

func TestDeferOrder2(t *testing.T) {
	// 打印顺序：C B A  （FILO）先入后出
	defer fmt.Println("A")
	defer fmt.Println("B")
	fmt.Println("C")
}

func TestDeferOrder3(t *testing.T) {
	// 打印顺序：D C B A
	defer fmt.Println("A")
	defer func() {
		defer fmt.Println("B")
		defer fmt.Println("C")
	}()
	defer fmt.Println("D")
}

func TestDeferOrder4(t *testing.T){
	// 打印顺序：A B
	testDeferOrder4()
}
func testDeferOrder4()int{
	f := func() int{
		fmt.Println("A")
		return 0
	}
	defer fmt.Println("B")
	return f()
}