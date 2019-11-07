package basetest

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestXunliang(t *testing.T) {
	//accessFile("F:\\chromedownload\\LPE-DLX")
	getCharFromString()
}

// 访问文件夹
func accessFile(fileDir string) {
	files, err := ioutil.ReadDir(fileDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 遍历字符串字符
func getCharFromString() {
	str := "abcdefg ffl"

	fmt.Println(strings.Title(str))

	for _, a := range []rune(str) {
		fmt.Printf("%c", a)
		fmt.Println("")
	}

	str2 := "烫烫烫烫"
	array := []rune(str2)
	n := len(array)
	for i := 0; i < n; i++ {
		ch := array[i]     // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch) //unicode 编码转十进制输出
	}

	//for _,s := range str {
	//	fmt.Println(s)
	//}
}
