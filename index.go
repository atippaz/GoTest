package main

import (
	client "test/grpc/client"
	server "test/grpc/server"
	two "test/two"
)

func main(){
	two.Test()
	server.StartServer()
	client.StartSendRequest()
}