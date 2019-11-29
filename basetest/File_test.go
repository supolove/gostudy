package basetest

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Unknwon/com"
	"io"
	"log"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	// Create 文件不存在创建，文件存在，将内容清空
	// Open 以读方式打开文件
	// OpenFile 以只读，只写，读写方式打开文件

	//f, err := os.Create("C:/test.txt")
	//
	//if err != nil {
	//	fmt.Println("create file error")
	//	return
	//}
	//
	//defer f.Close()

	readOnLine("PlayerInfoRecord", "/Users/supozheng/work/xinhun/gamelog/rlog/1_1_data.2019-10-30-15-20")
	f, _ := os.OpenFile("/Users/supozheng/Downloads/PSCC20192003.zip", os.O_RDONLY, os.ModePerm)
	finfo, _ := f.Stat()
	fmt.Println(finfo.Size())
}

type LogC struct {
	Length string
	Date   string
	Md5    string
}

func readOnLine(name, file string) {
	//Discard方法跳过后续的 n 个字节的数据，返回跳过的字节数。如果0 <= n <= b.Buffered(),该方法将不会从io.Reader中成功读取数据。
	//func (b *Reader) Discard(n int) (discarded int, err error)

	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("openfile error: %V", err)
	}

	defer f.Close()
	reader := bufio.NewReader(f)

	var rIndex int = 0
	c := GetCacheConfig(name)
	if c != nil {
		v := c.(map[string]interface{})["Length"]
		if i, e := com.StrTo(com.ToStr(v)).Int(); e == nil {
			rIndex = i
		}
	}

	i := 0
	reader.Discard(rIndex)
	for {
		if i > 3 {
			fmt.Println("lenth", rIndex)
			SaveCacheConfig(name, LogC{Length: com.ToStr(rIndex)})
			break
		}
		//b, err:= reader.ReadBytes('\n')
		b, _, err := reader.ReadLine()
		rIndex += len(b) + 1
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(string(b))
		}

		i++
	}
	fmt.Printf("index := %d\n", i)
	fmt.Println(rIndex)
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
	f, err := os.OpenFile("C:/test.txt", os.O_RDWR, os.ModeAppend|os.ModePerm)

	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()

	f.WriteString("123456")

}

// 创建目录
func testMakeDir() {
	os.Mkdir("dir", os.ModeDir|os.ModePerm)
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

func TestSectionReader_Seek(t *testing.T) {
	// Verifies that NewSectionReader's Seeker behaves like bytes.NewReader (which is like strings.NewReader)
	br := bytes.NewReader([]byte("foo"))
	sr := io.NewSectionReader(br, 0, int64(len("foo")))

	for _, whence := range []int{io.SeekStart, io.SeekCurrent, io.SeekEnd} {
		for offset := int64(-3); offset <= 4; offset++ {
			brOff, brErr := br.Seek(offset, whence)
			srOff, srErr := sr.Seek(offset, whence)
			if (brErr != nil) != (srErr != nil) || brOff != srOff {
				t.Errorf("For whence %d, offset %d: bytes.Reader.Seek = (%v, %v) != SectionReader.Seek = (%v, %v)",
					whence, offset, brOff, brErr, srErr, srOff)
			}
		}
	}

	// And verify we can just seek past the end and get an EOF
	got, err := sr.Seek(100, io.SeekStart)
	if err != nil || got != 100 {
		t.Errorf("Seek = %v, %v; want 100, nil", got, err)
	}

	n, err := sr.Read(make([]byte, 10))
	if n != 0 || err != io.EOF {
		t.Errorf("Read = %v, %v; want 0, EOF", n, err)
	}
}
