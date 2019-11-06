package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/*
1.打开文件
2.读取文件
3.打印
*/
func main() {
	file, err := os.Open("resources/song.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	fmt.Println("read the rile = " + file.Name())

	reader := bufio.NewReader(file)
	for {
		rb, rerr := reader.ReadBytes('\n')

		fmt.Println(string(rb))

		if rerr != nil && rerr == io.EOF {
			fmt.Println(rerr.Error())
			break
		}

		time.Sleep(time.Second * 2)
	}

}
