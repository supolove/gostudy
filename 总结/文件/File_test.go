package 文件

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

//遍历文件夹
func ReadDir(src string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			// 是文件夹
			fmt.Println(file.Name())
		} else {
			// 是文件
			fmt.Println(file.Name())
		}
	}
}

// 深度遍历文件夹
func ReadDeepDir(src string) {

	err := filepath.Walk(src, func(path string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func dir(src string) {
	// 获取路径指向文件夹名称
	filepath.Base(src)

	// 获取文件所在的路径
	filepath.Dir(src)

	// 获取文件扩展名
	filepath.Ext(src)

	// 获取文件名称
	name := strings.Replace(filepath.Base(src), filepath.Ext(src), "", 1)

	// 绝对路径
	absPath, _ := filepath.Abs("")

	// 相对路径
	relPath, _ := filepath.Rel("", "")

	// 路径拼接
	filepath.Join("basepath", "appendpath")

	fmt.Println(absPath)
	fmt.Println(relPath)
	fmt.Println(name)
}

// 创建文件夹

func operation(src string) {
	// 创建文件夹
	_ = os.Mkdir("xxx", os.ModePerm)

	// 删除单个文件,不能删除文件夹
	_ = os.Remove("xxx")

	// 删除文件夹或者文件
	_ = os.RemoveAll("xxx")

	// 重命名文件
	_ = os.Rename("xxx", "qqqq")
}

//ioutil
func fileop() {
	// 读取文件，整个文件读取到内存，不能读取大文件
	ioutil.ReadFile("filename")
	// 文件写入，每次都会新创建文件写入所有的内容，所以会覆盖之前的文件
	ioutil.WriteFile("filename", []byte{}, os.ModePerm)
}

// 大文件的读写
func copy() {
	fr, err := os.Open("")
	if err != nil {
		fmt.Println("open err")
		return
	}
	defer fr.Close()

	fw, err := os.Create("")
	if err != nil {
		fmt.Println("create file error")
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := fr.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完")
		}
		_, _ = fw.Write(buf[:n])
	}
}

func Test_file(t *testing.T) {
	dir := "/Users/supozheng/Downloads/bbb/usercenter-release-2.7.3"
	//ReadDir(dir)
	//ReadDeepDir(dir)
	//BaseDir(dir)
	//test(dir)
	fmt.Println(filepath.Rel(dir, path.Join(dir, "classes.jar")))
	fmt.Println(strings.Replace(filepath.Base(path.Join(dir, "classes.jar")), filepath.Ext(path.Join(dir, "classes.jar")), "", 1))
}
