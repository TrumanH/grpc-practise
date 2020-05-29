package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
	// ""
)


type HelloService struct {}

func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil 
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	fmt.Println("Register service")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen TCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// go rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))	// 换成json编码格式以实现跨语言
	}
}


