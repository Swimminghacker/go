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

func TestDeferOrder4(t *testing.T) {
	// 打印顺序：A B
	testDeferOrder4()
}
func testDeferOrder4() int {
	f := func() int {
		fmt.Println("A")
		return 0
	}
	defer fmt.Println("B")
	return f()
}

func TestDeferVar1(t *testing.T) {
	// 普通变量 会直接压入栈
	for index := 0; index < 5; index++ {
		defer func(i int) {
			fmt.Println(i)
		}(index)
	}
}

func TestDeferVar2(t *testing.T) {
	// 指针变量  会记录指针的地址，注意此处的index其实是同一个地址
	for index := 0; index < 5; index++ {
		defer func(i *int) {
			fmt.Println(*i) //  注意，这里打印的是55555，而不是44444
		}(&index)
	}
}

func TestDeferVar3(t *testing.T) {
	// 函数语句 会先计算函数的返回值再压入栈
	for index := 0; index < 5; index++ {
		defer func(i int) {
			fmt.Println(i)
		}(testDeferVar3(index))
	}
}
func testDeferVar3(i int) int {
	return i * 10
}

//
//---------------------------------------------------------
//知识点：return的作用就是给返回值赋值。如果定义了返回值名称，就给这个返回值赋值，
//如果没有，就给一个临时变量赋值，然后返回。
// ------------------------------------------------------------
func TestDeferReturnVar1(t *testing.T) {
	fmt.Println(testDeferReturnVar1())
	fmt.Println(testDeferReturnVar2())
}

func testDeferReturnVar1() (res int) {
	defer func() {
		res = res + 1 // res = res + 1= 2,最后返回的是res=2
	}()
	return 1 // 此处给res赋值res=1，然后再执行defer
}
func testDeferReturnVar2() int {
	res := 1
	defer func() {
		res = res + 1 // 此处res = res + 1 = 2  但是并没有影响返回值tmp，所以返回的结果还是1
	}()
	return res //这里会产生一个临时变量，比如tmp，然后tmp=res=1，然后执行defer
}

// 闭包
func TestDeferClosure(t *testing.T) {
	for index := 0; index < 5; index++ {
		defer func() {
			fmt.Println(index)
		}()
	}
}

// 当有多个panic的时候，最后一个才生效，而且捕捉到的是最后一个panic
func TestDeferPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("main panic")
}
