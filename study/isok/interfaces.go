package main

import "fmt"

// interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量
type says interface {
	say()
}

type dog struct{}

type cat struct{}

// 一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
func (d dog) say() {
	fmt.Println("dog ~ dog ~")
}

func (c cat) say() {
	fmt.Println("cat ~ cat ~")
}

func da(s says) {
	s.say()
}

func main() {
	var c cat
	var d dog
	var s says

	da(c)

	says.say(d)

	s = c
	s.say()
	s = d
	s.say()

}
