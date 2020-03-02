package main

import "fmt"

func sums(x, y int) (res int) {
	res = x + y
	return
}

func sum1(x int, y int) int {
	res := x + y
	return res
}

// 可变参数，必须放在函数参数的最后面
func sum2(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func f1() (res int) {
	return 100
}

// 定义一个接受值为函数并且返回值也是函数的函数
func funTofun(fn1 func() int) (res func(int) int) {
	res = func(x int) int {
		j := fn1()
		y := x + j
		return y
	}
	return
}

// 闭包：一个函数，一般作为返回值，和他相关的引用环境组合而成的实体，比如返回的这个函数还引用了外部函数的变量等
// 闭包的应用，比如引用了其他人的代码包中的函数readFunc，接受的参数为func()，但是我们想把taskFunc作为参数，这个时候需要用进行转换
func readFunc(f func()) {
	fmt.Println("this is read func")
	f()
}

func taskFunc(x, y int) {
	fmt.Println("this is task func", x+y)
}

// Task 用于实现readFunc(taskFunc)，因为readFunc接受的参数和taskFunc不一致，使用Task进行转换
func Task(tF func(int, int), x, y int) func() {
	return func() {
		tF(x, y)
	}
}

func main() {
	a := sums(1, 2)
	b := sum1(3, 5)
	sl1 := []int{1, 2, 3, 4, 5}
	c := sum2(sl1...)
	fmt.Println(a, b, c)

	fmt.Println(funTofun(f1)(200))
	// funcTofun(f1)，把f1函数作为参数传递过去，得到返回值为func(int)int的函数，所以需要funcTofun(f1)(200)给他传递一个参数

	readFunc(Task(taskFunc, 1, 2))
}
