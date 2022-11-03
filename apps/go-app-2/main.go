package main

import (
	"fmt"
	go_lib_2 "go-lib-2"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(go_lib_2.GoLib2("go-app-2"))
	fmt.Println(Hello("go-app-2"))
}
