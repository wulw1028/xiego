package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {
	urls := "http://127.0.0.1:8080/test"
	u, err := url.ParseRequestURI(urls)
	if err != nil {
		fmt.Println(err)
	}
	query := url.Values{}
	query.Set("name","wlw")
	query.Set("age","1000")
	u.RawQuery = query.Encode()

	tr := &http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(err)
	}
	rep, err := client.Do(req)
	defer rep.Body.Close()

	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}