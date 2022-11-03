package go_lib_2

import "strings"

func GoLib2(name string) string {
	result := []string{"GoLib2", name}
	return strings.Join(result, " ")
}
