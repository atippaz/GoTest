package main

import (
	"context"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}
// implement regex string here
func (s *server) BeefSummary(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.Text)
	return &pb.Response{Beef: map[string]int32{
		"t-bone": 4,
	"fatback": 1,
	"pastrami": 1,
	"pork": 1,
	"meatloaf": 1,
	"jowl": 1,
	"enim": 1,
	"bresaola": 1}}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
