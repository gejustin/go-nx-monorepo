package go_lib_2

import (
	"testing"
)

func TestGoLib2(t *testing.T) {
	result := GoLib2("works")
	if result != "GoLib2 works" {
		t.Error("Expected GoLib2 to append 'works'")
	}
}
