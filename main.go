package main

import "fmt"
import "github.com/yyl719133368/imooc_demo/pb/request"
import "github.com/golang/protobuf/proto"

func main() {
	fmt.Println("fk")
	loginRequest := request.LoginRequest{
		UserName:   "Gavin.Yang",
		Password:   "92d273941d98a8e1c1bb13ac163f0d4e40c5aa70",
		RetryTimes: 0}

	out, err := proto.Marshal(&loginRequest)

	if err == nil {
		fmt.Println(out)
	}
}
