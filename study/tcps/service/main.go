package main

import (
	"fmt"
	"net"
)

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil{
		fmt.Println(err)
		return
	}

	var tmp [128]byte
	n , err := conn.Read(tmp[:])
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(tmp[:n])

}