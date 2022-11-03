package go_lib_1

import (
	go_lib_2 "go-lib-2"
	"strings"
)

func GoLib1(name string) string {
	result := []string{"GoLib1", name, go_lib_2.GoLib2(name)}
	return strings.Join(result, " ")
}
