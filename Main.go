package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

var a string
var done chan int

func setup() {
	a = "hello, world"
	time.Sleep(time.Second * 5)
	done <- 1
}

var cccc chan int

func testBinxing4(i int) {
	go func() {
		d := <-cccc
		if i == d {
			fmt.Println("is true")
		} else {
			fmt.Println("is false")
		}

	}()
}

func main() {

	/*
		fmt.Println("start")

		f := os.Stdin

		data, err := ioutil.ReadAll(f)

		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	*/

	var buf bytes.Buffer
	si, _ := buf.ReadFrom(os.Stdin)
	fmt.Println(si)

	/*
		reader := bufio.NewReader(os.Stdin)

		result, err := reader.ReadString('\n')
		if err != nil {

			fmt.Println("read error:", err)
		}


		fmt.Println("result:", result)

	*/
}
