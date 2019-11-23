package main

import (
	"fmt"
	interf2 "gostudy/rpc/gorpc/interface"
	"log"
	"net/rpc"
)

func main1() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

// 2.0

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(interf2.HelloServiceName+".Hello", request, reply)
}

//var _ interf.HelloInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main2() {

	/*
		client, err := rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		var reply string
		err = client.Call(HelloServiceName + ".Hello", "world", &reply)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(reply)*/

	helloClient, _ := DialHelloService("tcp", "localhost:1234")
	var result string
	helloClient.Hello("Aaa", &result)
	fmt.Println(result)
}
