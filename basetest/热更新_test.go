package basetest

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"testing"
)

// fork一个进程
func fork() {
	netListener := net.Listener()
	file := netListener.File() // this returns a Dup()
	path := "/path/to/executable"
	args := []string{
		"-graceful"}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}
}

// Child initialization
func childinitialization() {
	server := &http.Server{Addr: "0.0.0.0:8888"}

	var gracefulChild bool
	var l net.Listever
	var err error

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")

	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		l, err = net.Listen("tcp", server.Addr)
	}
}

func Test_race(t *testing.T) {

}
