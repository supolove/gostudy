package basetest

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var in string
var out string

func init() {
	flag.StringVar(&in, "in", "", "input source dir!")
	flag.StringVar(&out, "out", "", "input target dir!")
}

func TestWxmain(t *testing.T) {
	// jpg : ffd8ff
	// png : 89504e47
	// gif : 47494638

	//flag.Parse()
	//if len(in) == 0 {
	//	fmt.Print("请输入-in 文件夹路径")
	//	return
	//}
	//
	//if len(in) == 0 {
	//	fmt.Print("请输入-out 文件夹路径")
	//	return
	//}

	in := "/Users/supozheng/Downloads/supo"
	out := "/Users/supozheng/Downloads/supo"

	files, error := GetAllFiles(in)

	if error != nil {
		println(error)
		return
	}
	var xorValue int
	const jpgHex int = 0xffd8
	const pngHex int = 0x8950
	const gifHex int = 0x4749

	tempFile := files[0]
	datTempFile, _ := os.OpenFile(tempFile, os.O_RDONLY, os.ModeDir)
	defer datTempFile.Close()
	reader := bufio.NewReader(datTempFile)
	b, _ := reader.Peek(2)
	rb, _ := bytesToIntS(b)
	xorValue = rb ^ pngHex
	if xorValue>>8 != xorValue&0x00FF {
		xorValue = rb ^ jpgHex
		if xorValue>>8 != xorValue&0x00FF {
			return
		}
	}

	xorValue = xorValue >> 8

	fmt.Printf("------ %x \n", xorValue)

	var index = 0
	for _, file := range files {
		datFile, err := os.OpenFile(file, os.O_RDONLY, os.ModeDir)
		if err != nil {
			continue
		}

		defer datFile.Close()
		reader2 := bufio.NewReader(datFile)

		s := fmt.Sprintf("%s/%d.png", out, index)
		f, err := os.Create(s)
		defer f.Close()
		writer := bufio.NewWriter(f)

		writeBuffer := []byte{}
		for {

			value, err2 := reader2.ReadByte()

			if err2 != nil && err2 == io.EOF {
				fmt.Println(err2.Error())
				break
			}

			valueResult := byte(value ^ byte(xorValue))
			writeBuffer = append(writeBuffer, valueResult)

		}
		writer.Write(writeBuffer)
		index++

	}

}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".dat")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".dat")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}
