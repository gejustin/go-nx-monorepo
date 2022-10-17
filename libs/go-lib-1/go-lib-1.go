package go_lib_1

import (
	go_lib_2 "go-lib-2"
)

func GoLib1(name string) string {
	result := "GoLib1 " + name + go_lib_2.GoLib2(name)
	return result
}
