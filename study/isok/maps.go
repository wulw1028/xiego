package main

import (
	"fmt"
	"strings"
)

func main() {
	var mp1 map[string]int
	mp1 = make(map[string]int, 10)
	mp1["test"] = 1
	fmt.Println(mp1, len(mp1))

	if v, ok := mp1["test"]; ok {
		fmt.Println(v)
	}

	// delete map key
	mp1["wlw"] = 2
	mp1["wulw"] = 3
	fmt.Println(mp1)
	delete(mp1, "wlw")
	fmt.Println(mp1)

	// []map[int]string
	// slice、map、chan 赋值的时候一定要初始化，make申请内存空间
	var sli1 = make([]map[int]string, 10, 10)
	sli1[0] = make(map[int]string, 2)
	sli1[0][100] = "wlw"
	sli1[0][200] = "wulw"
	fmt.Println(sli1)

	// 统计一个字符串中每个单词出现的次数，比如 "how do you do"中how=1，do=2，you=1
	var (
		str1 = "how do you do"
		mp2  = make(map[string]int, 4)
	)
	for _, v := range strings.Split(str1, " ") {
		if _, ok := mp2[v]; ok {
			mp2[v]++
		} else {
			mp2[v] = 1
		}
	}
	fmt.Println(mp2)
}
