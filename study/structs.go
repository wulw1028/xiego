package main

import "fmt"

// 自定义类型，有时候需要对内置类型定义一些方法
type myInt int

// 类型别名，rune是int32的类型别名
type aliasInt = int

// 结构体，值类型
type person struct {
	name  string
	age   int
	hobby []string
}

// 构造函数：约定俗成用new开头
func newPerson(name string, age int, hobby []string) *person {
	return &person{
		name:  name,
		age:   age,
		hobby: hobby,
	}
}

// 方法：是作用于特定类型的函数，特定类型+函数
// 接受者：调用方法的具体类型，这里是person，约定俗成命名为类型的首字母小写
func (p person) zouLu() {
	fmt.Printf("%v 走路\n", p.name)
}

func main() {
	var myint myInt = 100
	var aliasint aliasInt = 200

	fmt.Printf("%v %T\n", myint, myint)       // 100 main.myInt
	fmt.Printf("%v %T\n", aliasint, aliasint) // 200 int

	// 结构体
	var p1 person
	p1.name = "wlw"
	fmt.Printf("%T %v\n", p1, p1)

	// 匿名结构体，一般用于临时使用
	var nSt struct {
		name string
		age  int
	}
	nSt.name, nSt.age = "wulw", 100
	fmt.Printf("匿名结构体 %T %v\n", nSt, nSt)

	// 构造函数
	var stSlice = []string{"aaa", "bbb"}
	p2 := newPerson("wlw", 10, stSlice)
	fmt.Printf("构造函数 %v\n", p2)

	p2.zouLu()
}
