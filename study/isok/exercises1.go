package main

import (
	"fmt"
	"unicode"
)

func main() {
	// 1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	var (
		ints   = 12
		floats = 1.2
		bools  = true
		strs   = "test"
	)
	fmt.Printf("%T %T %T %T\n", ints, floats, bools, strs)

	// 2. 编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	var count int
	s := "hello沙河小王子"
	s1 := []rune(s)
	for _, v := range s1 {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)

	// 3. 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 10-i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 9; i >= 1; i-- {
		for j := 1; j <= 10-i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
