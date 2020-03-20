package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	str := "hello wlw"
	w.WriteHeader(200)
	q := r.URL.Query()
	name := q.Get("name")
	fmt.Println(name)
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
