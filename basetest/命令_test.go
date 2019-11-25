package basetest

import (
	"bytes"
	"fmt"
	"github.com/Unknwon/com"
	"os/exec"
	"testing"
	"time"
)

func TestCMD(t *testing.T) {

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

func TestComCMD(t *testing.T) {
	result, errResult, err := com.ExecCmd("ls", ".")
	if err != nil {
		panic(err)
	}

	if len(errResult) > 0 {
		fmt.Println(errResult)
	}

	fmt.Println(result)
}
