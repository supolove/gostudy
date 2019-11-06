package basetest

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main3() {
	// Create 文件不存在创建，文件存在，将内容清空
	// Open 以读方式打开文件
	// OpenFile 以只读，只写，读写方式打开文件

	f, err := os.Create("C:/test.txt")

	if err != nil {
		fmt.Println("create file error")
		return
	}

	defer f.Close()
}

func testOpen() {
	f, err := os.Open("C:/test.txt")

	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

}

func testOpenFile() {
	f, err := os.OpenFile("C:/test.txt", os.O_RDWR, 6)

	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

	f.WriteString("123456")

}

func testWrite() {
	f, err := os.OpenFile("C:/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Print("open file error")
		return
	}
	defer f.Close()
	f.WriteString("11111111111111")
	// seek() 修改文件读写指针位置
}

func testRead() {
	f, err := os.OpenFile("C:/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Print("open file error")
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		} else if err != nil {
			fmt.Print("ReadBytes err :", err)
		}
		fmt.Println(string(buf))
	}

}
