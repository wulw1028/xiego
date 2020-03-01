package main

import "fmt"

func main() {
	var i = 10
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("other")
	}

	switch j := 1; j {
	case 1, 2, 3:
		fmt.Println("in 1,2,3")
	}

	var intType interface{}
	intType = int64(11)
	switch intType.(type) {
	case int64:
		fmt.Println("int64")
	case int:
		fmt.Println("int")
	}
}
