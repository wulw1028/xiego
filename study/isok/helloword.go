package main

import "fmt"

var (
	name  string
	age   int
	test         = "wlw"
	name1 string = "wlw1"
)

const (
	// PI init
	PI string = "3.14"
	// CI init
	CI = 3.14
)

const (
	// t2、t3没有赋值，则和上面有赋值的常量相等
	t1 = 100
	t2
	t3
)

const (
	c1 = iota
	c2
	c3
)

const (
	// iota在const中出现初始化为0。const中每新增一行常量定义iota计数一次
	d1, d2 = iota + 1, iota + 2 // d1 = 0 + 1, d2 = 0 + 2
	d3, d4                      // d3 = 1 + 1, d4 = 1 + 2
)

func main() {
	var isOk bool
	name1 := "wlw2"
	var s3 = "wlw3"
	fmt.Println("hello word", isOk, name1, s3, PI, CI)
	fmt.Println(t1, t2, t3)
	fmt.Println(d1, d2, d3, d4)
}
