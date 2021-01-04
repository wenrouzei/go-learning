package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-gateway/proto"
)

const (
	port = ":8088"
)

type server struct{
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: req.GetValue()}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}