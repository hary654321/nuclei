package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.baidu.com/")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
}
