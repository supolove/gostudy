package basetest

import (
	"fmt"
	"io"
	"os"
)

func main13() {

}

// copy mp3file
func copyMp3() {
	f_r, err := os.Open("")
	if err != nil {
		fmt.Println("open err")
		return
	}

	defer f_r.Close()

	f_w, err := os.Create("")
	if err != nil {
		fmt.Println("create file error")
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完")
		}

		f_w.Write(buf[:n])
	}

}

// access dir
func test1() {
	fmt.Println("请输入查询的目录")
	var path string
	fmt.Scan(&path)
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)

	if err != nil {
		fmt.Println("openFile error :", err)
		return
	}

	defer f.Close()

	info, err := f.Readdir(-1)

	if err != nil {
		fmt.Println("read file error")
		return
	}

	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name(), "是一个目录")
		} else {
			fmt.Println(fileInfo.Name(), "是一个文件")
		}
	}
}
