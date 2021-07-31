package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	go func() {
		for range time.Tick(time.Second * 2) {
			bufOut := new(bytes.Buffer)
			bufErr := new(bytes.Buffer)

			cmd := exec.Command("pkgtool2")
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
	}()

	time.Sleep(time.Second * 5)

}
