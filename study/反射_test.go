package study

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFanshe(t *testing.T) {
	str := "123"
	var fl float64 = 10
	rt := reflect.TypeOf(str)
	fmt.Println(rt.Name())
	fmt.Println(rt.Kind())
	fmt.Println(rt.String())

	rt = reflect.TypeOf(fl)
	fmt.Println(rt.Name())
	fmt.Println(rt.Kind())
	fmt.Println(rt.String())
	if rt.Kind() == reflect.Float64 {
		fmt.Println("isFloat64")
	}
}
