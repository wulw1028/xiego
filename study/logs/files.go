package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fh, err := os.Open("./wlw.log")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	var fileByte [128]byte
	for {
		n, err := fh.Read(fileByte[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(fileByte[:n]))
		if n < 128 {
			return
		}
	}
}

func readFromFileByBufio() {
	fh, err := os.Open("./wlw.log")
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()

	reader := bufio.NewReader(fh)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(str)
	}
}

func readFromFileByIoutil() {
	if fileByte, err := ioutil.ReadFile("./wlw.log"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(fileByte))
	}

}

func writeFromFile() {
	fh, err := os.OpenFile("./wlw.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()
	fh.Write([]byte("wulw"))
	fh.WriteString("wulw")
}

func writeFromFileByBufio() {
	fh, err := os.OpenFile("./wlw.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()
	writer := bufio.NewWriter(fh)
	writer.WriteString("writeFromFileByBufio")
	writer.Flush()
}

func writeFromFileByIoutil() {
	err := ioutil.WriteFile("./wlw.log", []byte("writeFromFileByIoutil"), 0664)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//readFromFile()
	//readFromFileByBufio()
	//readFromFileByIoutil()
	//writeFromFile()
	//writeFromFileByBufio()
	writeFromFileByIoutil() // 覆盖写入，会自动创建文件
}
