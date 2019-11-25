package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {

	for range time.Tick(time.Duration(time.Second * 30)) {
		//fmt.Println("tick",time.Now())
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)

		cmd := exec.Command("git", "pull")
		cmd.Dir = "."
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr

		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		if len(bufErr.Bytes()) > 0 {
			panic(bufErr.String())
		}

		fmt.Println(bufOut.String())
	}
}
