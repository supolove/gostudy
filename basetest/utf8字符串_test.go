package basetest

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUtf8(t *testing.T) {
	u := "我们"
	fmt.Println(len(u))
	fmt.Printf("%T\n", u)

	fmt.Println(utf8.RuneCountInString(u))
}

func Test(ta *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
		{name: "123"},
		{name: "456"},
	}
	for _, test := range tests {
		ta.Run(test.name, func(t *testing.T) {
			//fmt.Println()
		})
	}
}
