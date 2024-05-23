package main

import (
	"context"

	pb "grpc/proto"

	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.BeefSummary(ctx, &pb.Request{Text: "sdsd"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	resultString:="\n{\n"
	for i,d :=  range r.Beef{
		resultString+= "	"+i +" : "+ strconv.Itoa(int(d)) + "\n"
	}
	resultString+="}"
	log.Printf(resultString)
}
