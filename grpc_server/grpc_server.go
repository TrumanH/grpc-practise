package main

import (
	"fmt"
	"log"
	"net"
)

type HelloServiceImpl struct {}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(gprcServer, new(HelloServiceImpl)) 	// defined in hello.pb.go
	fmt.Println("Register service")
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("net Listen:", err)
	}
	grpcServer.Serve(lis)
}