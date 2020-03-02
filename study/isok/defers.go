package main

import "fmt"

/*
Go语言中函数的return不是原子操作，在底层是分两步执行
    - 返回值赋值
    - 真正的RET返回
	- 函数中如果存在defer，那么defer执行的时机：返回值赋值 -> defer —> 真正的RET返回
*/

func f1() int {
	x := 6
	defer func() {
		x++
	}()
	return x // 因为f1返回值为int类型，没有声明变量名。返回值赋值：x=6，开辟了一个内存空间，存放了6，所以返回值和x不是一个地址空间，RET结果为6
}

func f2() (x int) {
	x = 6
	defer func() {
		x++
	}()
	return x // 因为f2返回值定义了变量x，所以。返回值赋值：x=6，defer 中x和返回值x为一个值，x=7，RET结果为7
}

func f3() (x int) {
	defer func() {
		x++
	}()
	return 6 // 因为f3定义了变量x。return 6。返回值赋值：x=6，defer中 x++=7，RET结果为7
}

func f4() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值赋值：y=x=5。因为int是值类型，这里是值copy，所以defer修改x不会对y有改动。RET结果为5
}

func f5() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值赋值：x=5，defer这里是以函数传参的方式传递给匿名函数的，是值copy，所有defer修改的x和返回值的x不是一个。RET结果5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}
// defer calc("AA", x, calc("A", x, y))虽然会压栈延时处理，但是函数的初始化执行和变量的引用还是会压栈前引用
// defer calc("AA", x, calc("A", x, y))
// 1. 执行calc("A", x, y) -> calc("A", 1, 2)，x和y会取当前已经初始化的变量定义
// 2. 压栈 defer calc("AA", 1, 2)

func main() {
	fmt.Println("start")
	defer fmt.Println(2) // 压到栈中
	defer fmt.Println(1) // 压到栈中，当函数快结束的时候在弹出执行
	fmt.Println("end")
	fmt.Println("f1", f1())
	fmt.Println("f2", f2())
	fmt.Println("f3", f3())
	fmt.Println("f4", f4())
	fmt.Println("f5", f5())
}
