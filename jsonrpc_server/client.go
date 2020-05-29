package main 
import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
	"net"
	"fmt"
)

func main() {
	/*
	client, err := rpc.Dial("tcp", "localhost:1234")	// 拨号rpc服务
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply String
	// 注意请求和返回参数类型都使用了hello.pd.go中的String类型(使用了Protobuf编码)
	err = client.Call("HelloService.Hello", String{Value:"truman"}, &reply)	// 调用具体方法，后两个为参数，将response传给reply
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	*/

	// 外层换用json编码而不是默认的Gob编码，以实现跨语言
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply String
	err = client.Call("HelloService.Hello", String{Value:"truman"}, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}