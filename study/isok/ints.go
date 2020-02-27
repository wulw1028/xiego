package main

import "fmt"

func main() {
	var ints = 110
	fmt.Printf("10 -> %d\n", ints)
	fmt.Printf("2 -> %b\n", ints)
	fmt.Printf("8 -> %o\n", ints)
	fmt.Printf("16 -> %x\n", ints)
	fmt.Printf("16 -> %X\n", ints)

	var baInt = 0777
	fmt.Printf("10 -> %d\n", baInt)

	var shiliuInt = 0x12ef
	fmt.Printf("10 -> %d\n", shiliuInt)

	fmt.Printf("type is %T\n", baInt)

	int88 := int8(12)
	fmt.Printf("type is %T\n", int88)

	var int166 int16 = 123
	fmt.Printf("type is %T\n", int166)
}
