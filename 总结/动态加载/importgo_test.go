package importgo

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestImportgo(t *testing.T) {
	//err := exec.Command("go","run", "/Users/supozheng/go/src/gostudy/动态加载/testgo/gofile.go").Run()

	cmd := exec.Command("go", "run", "/Users/supozheng/go/src/gostudy/动态加载/testgo/gofile.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("err : %v", err)
	}
}
