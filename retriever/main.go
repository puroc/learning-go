package main

import (
	"fmt"
	"pud/learning-go/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}
func main() {
	var r Retriever = mock.RetrieverImpl{Contents: "12345"}
	fmt.Println(download(r))
}
