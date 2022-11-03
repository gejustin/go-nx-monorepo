package main

import (
	"fmt"

	go_lib_1 "go-lib-1"

	"rsc.io/quote"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(go_lib_1.GoLib1("test"))
	fmt.Println(quote.Go())
}
