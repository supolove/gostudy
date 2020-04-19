package interf

const HelloServiceName = "path/to/pkg.HelloService"

type HelloInterface interface {
	Hello(request string, reply *string) error
}
