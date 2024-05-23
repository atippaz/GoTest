package main

import (
	"fmt"
	client "test/grpc/client"
	server "test/grpc/server"
	two "test/two"
)

func main(){
	fmt.Println("1. \n")
	two.Test()
	fmt.Println("\n")
	fmt.Println("2. \n")
	server.StartServer()
	client.StartSendRequest()
	fmt.Println("\n")

}