package main

import (
	"fmt"
	"github.com/Unknwon/com"
	"math"
	"strings"
)

func main() {

	versionName := "9.9.9"
	versionCode := 0

	fmt.Println(math.Pow10(3))

	vns := strings.Split(versionName, ".")
	for i, vn := range vns {
		v, _ := com.StrTo(vn).Int()
		x := int(math.Pow10((len(vns) - i - 1) * 2))
		versionCode += v * x
		//com.ex
	}
	fmt.Println(versionCode)

}
