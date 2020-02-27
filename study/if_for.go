package main

import "fmt"

func main() {
	var age = 20
	if age > 10 {
		fmt.Println(">10")
	} else if age > 20 {
		fmt.Println(">20")
	} else {
		fmt.Println("other")
	}

	// age1在if作用域中生效
	if age1 := 19; age1 > 18 {
		fmt.Println(">18")
	}

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	var i = 1
	for ; i < 2; i++ {
		fmt.Println(i)
	}

	for i < 4 {
		fmt.Println(i)
		i = i + 2
	}

	str := "hello_wlw"
	for k, v := range str {
		fmt.Println(k, string(v))
	}
}
