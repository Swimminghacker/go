package main

import (
	"fmt"
	"testing"
	"time"
)

//  panic在当前函数执行之后，程序直接退出，其他goroutine感知不到panic的存在
func TestPanic1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		panic(1)
	}()

	time.Sleep(1)
	fmt.Println("test")
}
