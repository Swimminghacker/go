package main

import (
	"fmt"
	"testing"
)

type testInterface interface {
	test()
}
type testStruct struct {
}

func (t *testStruct) test() {

}
func TestNilInterface(t *testing.T) {
	var test testInterface


	test = getTest1()
	//test = nil 是true，因为返回的接口类型和接口值都是nil
	if test == nil {
		fmt.Println("test1 is nil")
	}

	// test不是nil，因为返回的接口类型为*testStruct类型，接口类型和接口值都为nil的情况下，接口才为nil。
	test = getTest2()
	if test == nil {
		fmt.Println("test2 is nil")
	}

}
func getTest1() testInterface {
	return nil
}
func getTest2() *testStruct {
	return nil
}
