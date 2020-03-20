package splits

import "strings"

func Splits(str string, sep string) (res []string) {
	res = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index > 0 {
		res = append(res, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	res = append(res, str)
	return
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
