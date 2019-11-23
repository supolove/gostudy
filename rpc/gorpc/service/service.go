package service

import (
	interf "gostudy/rpc/gorpc/interface"
	"log"
	"net"
	"net/rpc"
)

/* 1.0 */
type HelloService struct{}

// 其中Hello方法必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，
// 其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main1() {

	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}

/* 2.0 */
const HelloServiceName = "path/to/pkg.HelloService"

func RegisterHelloService(svc interf.HelloInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
